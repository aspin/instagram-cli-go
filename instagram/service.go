package instagram

type Service interface {
	Query(username string, postURL string) error
	Followers(username string) (map[UserID]User, error)
	Post(postURL string) (Media, error)
	PostLikers(postURL string) (map[UserID]User, error)
	PostComments(postURL string) ([]Comment, error)
}

type cachedService struct {
	cached map[string]pythonService
}

func NewService(authUsername string, authPassword string) Service {
	return &cachedService{}
}

func (c *cachedService) Query()

func (c *cachedService) Followers(username string) (map[UserID]User, error) {
}

func (c *cachedService) Post(postURL string) (Media, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cachedService) PostLikers(postURL string) (map[UserID]User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cachedService) PostComments(postURL string) ([]Comment, error) {
	//TODO implement me
	panic("implement me")
}
