package entity

// Course defines the structure for the API
type Course struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Description      string `json:"description"`
	TechStack        string `json:"techstack"`
	ModuleQuantity   string `json:"moduleQuantity"`
	WorkshopQuantity string `json:"workshopQuantity"`
	CreatedOn        string `json:"createdon"`
	Version          int    `json:"version"`
}
