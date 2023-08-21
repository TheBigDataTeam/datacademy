package course

type GetRequest struct {
	Id string `json:"id"`
}

type CreateRequest struct {
	Title            string `json:"title" validate:"required"`
	Author           string `json:"author" validate:"required"`
	Description      string `json:"description" validate:"required"`
	TechStack        string `json:"techstack" validate:"required"`
	ModuleQuantity   string `json:"module_quantity" validate:"required"`
	WorkshopQuantity string `json:"workshop_quantity" validate:"required"`
}

type UpdateRequest struct {
	Title            string `json:"title" validate:"required"`
	Author           string `json:"author" validate:"required"`
	Description      string `json:"description" validate:"required"`
	TechStack        string `json:"techstack" validate:"required"`
	ModuleQuantity   string `json:"module_quantity" validate:"required"`
	WorkshopQuantity string `json:"workshop_quantity" validate:"required"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}
