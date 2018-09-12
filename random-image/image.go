package image

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "loremflickr.com",
	Path:   "612/612", // this value is fixed cause this is the maximum size of an instagram image
}

// Get get imge from loremflickr
func Get() ([]byte, error) {
	req, err := http.NewRequest("GET", baseURL.String(), nil)
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
