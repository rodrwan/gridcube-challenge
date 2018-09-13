package image

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "loremflickr.com",
	Path:   "612/612", // this value is fixed cause this is the maximum size of an instagram image
}

// Service hold image size
type Service struct {
	Size int32
}

// NewService expose a new image.Service
func NewService(size int32) *Service {
	return &Service{
		Size: size,
	}
}

// Get get imge from loremflickr
func (is *Service) Get() ([]byte, error) {
	url := fmt.Sprintf("%s/%d/%d", baseURL.String(), is.Size, is.Size)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cc := &http.Client{}

	resp, err := cc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
