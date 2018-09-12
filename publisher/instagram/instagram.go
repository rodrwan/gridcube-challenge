package instagram

import (
	"bytes"

	"github.com/ahmdrz/goinsta"
)

// Session hold goinsta session
type Session struct {
	insta *goinsta.Instagram
}

// NewSession creates a new instagram session
func NewSession(username, password string) (*Session, error) {
	insta := goinsta.New(username, password)
	if err := insta.Login(); err != nil {
		return nil, err
	}
	return &Session{
		insta: insta,
	}, nil
}

// Close free the current ig session
func (s *Session) Close() error {
	return s.insta.Logout()
}

// UploadPhoto uploads the given image with the given caption
func (s *Session) UploadPhoto(img []byte, caption string) error {
	readImage := bytes.NewReader(img)
	_, err := s.insta.UploadPhoto(readImage, caption, 87, 0)
	if err != nil {
		return err
	}

	return nil
}
