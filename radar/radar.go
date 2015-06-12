package radar

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gographics/imagick/imagick"
)

var (
	images = []*imageInfo{
		&imageInfo{"TopoShort", "http://radar.weather.gov/ridge/Overlays/Topo/Short/%s_Topo_Short.jpg"},
		&imageInfo{"RadarImg", "http://radar.weather.gov/ridge/RadarImg/N0R/%s_N0R_0.gif"},
		&imageInfo{"CountyShort", "http://radar.weather.gov/ridge/Overlays/County/Short/%s_County_Short.gif"},
		&imageInfo{"HighwaysShort", "http://radar.weather.gov/ridge/Overlays/Highways/Short/%s_Highways_Short.gif"},
		&imageInfo{"CityShort", "http://radar.weather.gov/ridge/Overlays/Cities/Short/%s_City_Short.gif"},
	}
)

func init() {
	imagick.Initialize()
}

type imageInfo struct {
	Name string
	Fmt  string
}

type WeatherRadar struct {
	Location string
	Images   []*Image
	Radar    []byte
}

type Image struct {
	URL string
	Raw []byte
}

func New(location string) *WeatherRadar {
	wr := &WeatherRadar{Location: location}
	return wr
}

func (wr *WeatherRadar) GetImageBlob() ([]byte, error) {
	err := wr.loadImages()
	if err != nil {
		return []byte(""), err
	}
	// Now return the data
	return wr.Radar, nil
}

func (wr *WeatherRadar) loadImages() error {
	wr.Images = make([]*Image, len(images))
	ch := make(chan int)
	for index, image := range images {
		go wr.downloadImage(index, image, ch)
	}
	for _ = range images {
		<-ch
	}

	// Finally, generate our composites
	return wr.generateComposite()
}

func (wr *WeatherRadar) downloadImage(index int, ii *imageInfo, ch chan int) {
	wr.Images[index] = &Image{
		URL: fmt.Sprintf(ii.Fmt, wr.Location),
	}

	var err error
	wr.Images[index].Raw, err = urlFetch(wr.Images[index].URL)
	if err != nil {
		log.Printf("Unable to download %s: %s", wr.Images[index].URL, err)
	}
	ch <- 1
}

func urlFetch(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}
	defer func() { _ = r.Body.Close() }()
	data, err := ioutil.ReadAll(r.Body)
	return data, err
}

func (wr *WeatherRadar) generateComposite() error {
	// Start with Radar image, then add the rest
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.ReadImageBlob(wr.Images[0].Raw)
	if err != nil {
		return fmt.Errorf("ReadImageBlob: %s", err)
	}

	for i := 1; i < len(wr.Images); i++ {
		image := wr.Images[i]
		newMw := imagick.NewMagickWand()
		err = newMw.ReadImageBlob(image.Raw)
		if err != nil {
			return fmt.Errorf("ReadImageBlob: %s", err)
		}
		err = mw.CompositeImage(newMw, imagick.COMPOSITE_OP_ATOP, 0, 0)
		if err != nil {
			return fmt.Errorf("CompositeImage: %s", err)
		}
		newMw.Destroy()
	}

	wr.Radar = mw.GetImageBlob()

	return nil
}

//   # Create the final radar-image using imagemagick.
//   composite -compose atop ${cache}/${1}_N0R_0.gif ${cache}/${1}_Topo_Short.jpg ${cache}/radar-image.jpg
//   composite -compose atop ${cache}/${1}_County_Short.gif ${cache}/radar-image.jpg ${cache}/radar-image.jpg
//   composite -compose atop ${cache}/${1}_Highways_Short.gif ${cache}/radar-image.jpg ${cache}/radar-image.jpg
//   composite -compose atop ${cache}/${1}_City_Short.gif ${cache}/radar-image.jpg ${cache}/radar-image.jpg
//   composite -compose atop ${cache}/${1}_Warnings_0.gif ${cache}/radar-image.jpg ${cache}/radar-image.jpg
//   composite -compose atop ${cache}/${1}_N0R_Legend_0.gif ${cache}/radar-image.jpg ${cache}/radar-image.jpg
