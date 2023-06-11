package auth

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type Project struct {
	ID   string
	Name string
}

type Client struct {
	ID        string
	ProjectID string
	Secret    string
}

type Channel struct {
	ID       string
	ClientID string
}
