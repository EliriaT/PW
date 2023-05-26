package storage

import (
	"crypto/sha512"
	"errors"
	"fmt"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"io"
)

type Storage interface {
	Save(l *Link) error
	PickRandom(userName string) (*Link, error)
	Remove(l *Link) error
	IsLinkPresent(l *Link) (bool, error)
	GetAllLinks(userID string) (page []*Link, err error)
}

// one link
type Link struct {
	URL    string
	UserID string
	//Created time.Time
}

func (l Link) Hash() (string, error) {
	hash := sha512.New()

	if _, err := io.WriteString(hash, l.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(hash, l.UserID); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func PagesToString(links []*Link) string {
	response := ""
	for _, l := range links {
		response = fmt.Sprintf("%s \n\n\n %s", response, l.URL)
	}
	return response
}

var ErrNoSavedLinks = errors.New("no saved links")
