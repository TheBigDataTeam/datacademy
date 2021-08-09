package users

// User represents a user
type User struct {
	ID        string
	Email     string
	Fullname  string
	IsAdmin   bool
	CreatedOn string
	Version   int
}
