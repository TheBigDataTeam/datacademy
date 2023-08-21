package entity

type User struct {
	ID        string `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	IsAdmin   bool   `json:"-"`
	CreatedOn string `json:"created_on"`
	Version   int    `json:"version"`
}
