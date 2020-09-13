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
type WorkshopDB struct {
	ID            int
	Kind          int
	Status        int
	Date          int
	Teaser        string
	LocationName  string
	LocationOnMap string
}

// Workshop a workshop with all (linked) data and json mapping
type Workshop struct {
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
	ContentIDs    []string `json:"content"`
}

// toDTO() creates a WorkshopDB from a Workshop only maps common fields
func (workshop Workshop) ToDB() WorkshopDB {
	var db WorkshopDB
	db.ID = workshop.ID
	db.Kind = workshop.Kind
	db.Status = workshop.Status
	db.Date = workshop.Date
	db.Teaser = workshop.Teaser
	db.LocationName = workshop.LocationName
	db.LocationOnMap = workshop.LocationName
	return db
}

// fromDB returns a Workshop from a WorkshopDB, only maps common fields
func (db WorkshopDB) FromDB() Workshop {
	var workshop Workshop
	workshop.ID = db.ID
	workshop.Kind = db.Kind
	workshop.Status = db.Status
	workshop.Date = db.Date
	workshop.Teaser = db.Teaser
	workshop.LocationName = db.LocationName
	workshop.LocationName = db.LocationOnMap
	return workshop
}
