package users

// User represents a user
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	IsAdmin   bool   `json:"-"`
	CreatedOn string `json:"createdon"`
	Version   int    `json:"version"`
}
