package instagram

// pythonService uses all info to fetch details exactly once, then returns the same values ad infinitum
type pythonService struct {
	followers    map[UserID]User
	post         Media
	postLikers   map[UserID]User
	postComments []Comment
}

func newPythonService(authUsername, authPassword, targetUsername, targetPost string) (*pythonService, error) {
	ps := &pythonService{}
	err := ps.run(authUsername, authPassword, targetUsername, targetPost)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (ps pythonService) run(authUsername, authPassword, targetUsername, targetPost string) error {
	return nil
}
