package instagram

import (
	"fmt"
)

type Service interface {
	Followers(username string) (UserSet, error)
	Post(postURL string) (Media, error)
	PostLikers(postURL string) (UserSet, error)
	PostComments(postURL string) ([]Comment, error)
}

type cachedService struct {
	cached map[string]pythonService
}

func NewService(authUsername string, authPassword string) Service {
	return &cachedService{}
}

func (c *cachedService) Followers(username string) (UserSet, error) {
	fixed := map[UserID]User{
		"1234": {ID: "1234", Username: "kevinchen"},
	}
	return fixed, nil
}

func (c *cachedService) Post(postURL string) (Media, error) {
	fixed := Media{
		PostName: "about bears!",
	}
	return fixed, nil
}

func (c *cachedService) PostLikers(postURL string) (UserSet, error) {
	fixed := map[UserID]User{
		"1234": {ID: "1234", Username: "kevinchen"},
	}
	return fixed, fmt.Errorf("http 500 failed: %s", postURL)
}

func (c *cachedService) PostComments(postURL string) ([]Comment, error) {
	fixed := []Comment{
		{
			Text: "answered the question",
		},
		{
			Text: "congrats on 5k!",
		},
	}
	return fixed, nil
}
