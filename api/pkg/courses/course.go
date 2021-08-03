package courses

// Course defines the structure for an API
type Course struct {
	ID          int      `json:"id"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Theme       string   `json:"theme"`
	Author      string   `json:"author" validate:"required,author"` // TODO author is to be Autor struct
	AuthorID    int      `json:"authorid"`
	TechStack   []string `json:"techstack"`
	Syllabus    []string `json:"syllabus"`
	Duration    string   `json:"-"`
	Beneficiars []string `json:"beneficiars"`
	Difficulty  string   `json:"-"`
	CreatedOn   string   `json:"-"`
	UpdatedOn   string   `json:"-"`
}
