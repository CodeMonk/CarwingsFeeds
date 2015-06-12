package radar

import (
	"testing"
)

func TestNew(t *testing.T) {
	wr := New("MTX")
	t.Logf("%#v", wr)
}

func TestGetImageBlob(t *testing.T) {
	wr := New("MTX")
	t.Logf("%#v", wr)

	blob, err := wr.GetImageBlob()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Blob is %d bytes", len(blob))
}
