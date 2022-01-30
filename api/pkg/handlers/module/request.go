package module

type GetRequest struct {
	Id string `json:"id"`
}
type ListRequest struct {
	CourseId string `json:"course_id"`
}
type CreateRequest struct {
	CourseId string `json:"course_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	Badge    string `json:"badge" validate:"required"`
}
type UpdateRequest struct {
	Id       string `json:"id" validate:"required"`
	CourseId string `json:"course_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	Badge    string `json:"badge" validate:"required"`
}
type DeleteRequest struct {
	Id string `json:"id"`
}
