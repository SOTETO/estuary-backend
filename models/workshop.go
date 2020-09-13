package models

const (
	TypeUNKOWN = 0
	TypePS     = 1
	TypeIDEA   = 2
)

const (
	StatusWIP    = 0
	StatusPUBLIC = 1
	StatusUNKOWN = 2
)

//Workshop as in the db
type Workshop struct {
	ID            int
	Kind          int
	Status        int
	Date          int
	Teaser        string
	LocationName  string
	LocationOnMap string
}

// WorkshopDTO a workshop DTO
type WorkshopDTO struct {
	ID            int      `json:"id"`
	Kind          int      `json:"type"`
	Status        int      `json:"status"`
	Date          int      `json:"date"`
	Teaser        string   `json:"teaser"`
	LocationName  string   `json:"LocationName"`
	LocationOnMap string   `json:"LocationOnMap"`
	Likes         int      `json:"upvotes"`
	Tags          []string `json:"tags"`
	Authors       []string `json:"authors"`
	Content       []string `json:"content"`
}

// toDTO() creates a WorkshopDTO from a Workshop
func (workshop Workshop) toDTO() WorkshopDTO {
	var dto WorkshopDTO
	dto.ID = workshop.ID
	dto.Kind = workshop.Kind
	dto.Status = workshop.Status
	dto.Date = workshop.Date
	dto.Teaser = workshop.Teaser
	dto.LocationName = workshop.LocationName
	dto.LocationOnMap = workshop.LocationName
	return dto
}
