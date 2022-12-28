package instagram

type UserID string

type Media struct {
	PostName string
}

type User struct {
	ID       UserID
	Username string
}

type UserSet map[UserID]User

type Comment struct {
	Text string
}
