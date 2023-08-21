package entity

type Course struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Description      string `json:"description"`
	TechStack        string `json:"techstack"`
	ModuleQuantity   string `json:"module_quantity"`
	WorkshopQuantity string `json:"workshop_quantity"`
	CreatedOn        string `json:"created_on"`
	Version          int    `json:"version"`
}
