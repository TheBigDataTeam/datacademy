package entity

type Module struct {
	Id        string `json:"id"`
	CourseId  string `json:"course_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Badge     string `json:"badge"`
	CreatedOn string `json:"created_on"`
	Version   int    `json:"version"`
}
