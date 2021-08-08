package authors

// Author defines the structure of the author api
type Author struct {
	ID               string `json:"id"`
	CourseID         string `json:"courseid"`
	Email            string `json:"email"`
	Fullname         string `json:"fullname"`
	Twitter          string `json:"twitter"`
	Facebook         string `json:"facebook"`
	Instagram        string `json:"instagram"`
	Location         string `json:"location"`
	Bio              string `json:"bio"`
	ShortDescription string `json:"shortdescription"`
	Speciality       string `json:"speciality"`
	Features         string `json:"features"`
	CreatedOn        string `json:"createdon"`
	Version          int    `json:"version"`
}
