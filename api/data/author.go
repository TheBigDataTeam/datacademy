package data

// Author defines the structure of the author api
type Author {
	ID int `json:"id"`
	CourseID []int ´json:"features"´
	Email string `json:"email"`
	Fullname string `json:"fullname"`
	Speciality string `json:"speciality"`
	Features []string `json:"features"`
	CreatedOn   string `json:"-"`
}

