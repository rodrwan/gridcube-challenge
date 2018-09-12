package image

import (
	"testing"
)

func TestGet(t *testing.T) {
	img, err := Get()
	if err != nil {
		t.Fatal(err)
	}

	if len(img) == 0 {
		t.Error("Image should be greater than 0")
	}
}
