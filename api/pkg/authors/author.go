package authors

// Author defines the structure of the author api
type Author struct {
	ID               string   `json:"id"`
	CourseID         []int    `json:"courseid"`
	Email            string   `json:"email"`
	Twitter          string   `json:"twitter"`
	Facebook         string   `json:"facebook"`
	Instagram        string   `json:"instagram"`
	Location         string   `json:"location"`
	Fullname         string   `json:"fullname"`
	Bio              string   `json:"bio"`
	ShortDescription string   `json:"shortdescription"`
	Speciality       string   `json:"speciality"`
	Features         []string `json:"features"`
	CreatedOn        string   `json:"-"`
}
