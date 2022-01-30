package module

type Module struct {
	Id       string `json:"id"`
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Badge    string `json:"badge"`
}
