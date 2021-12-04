package users

// User represents a user
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	IsAdmin   bool   `json:"-"`
	CreatedOn string `json:"createdon"`
	Version   int    `json:"version"`
}
