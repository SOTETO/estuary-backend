package models

// Tabler interface to set the associated db table name for a struct
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by WorkshopDB to `workshop`
func (WorkshopDB) TableName() string {
	return "workshop"
}

//WorkshopDB model for a record i´n the db
type WorkshopDB struct {
	UUID          string `gorm:"column:UUID; primaryKey"`
	Date          int    `gorm:"column:date"`
	Teaser        string `gorm:"column:teaser"`
	LocationName  string `gorm:"column:locationName"`
	LocationOnMap string `gorm:"column:locationOnMap"`
}

// Workshop a workshop with all (linked) data and json mapping
type Workshop struct {
	UUID          string   `json:"UUID"`
	Date          int      `json:"date"`
	Teaser        string   `json:"teaser"`
	LocationName  string   `json:"LocationName"`
	LocationOnMap string   `json:"LocationOnMap"`
	Likes         int      `json:"upvotes"`
	Tags          []string `json:"tags"`
	Authors       []string `json:"authors"`
	ContentUUIDs  []string `json:"content"`
}

// ToDB creates a WorkshopDB from a Workshop only maps common fields
func (workshop Workshop) ToDB() WorkshopDB {
	var db WorkshopDB
	db.UUID = workshop.UUID
	db.Date = workshop.Date
	db.Teaser = workshop.Teaser
	db.LocationName = workshop.LocationName
	db.LocationOnMap = workshop.LocationOnMap
	return db
}

// FromDB returns a Workshop from a WorkshopDB, only maps common fields
func (db WorkshopDB) FromDB() Workshop {
	var workshop Workshop
	workshop.UUID = db.UUID
	workshop.Date = db.Date
	workshop.Teaser = db.Teaser
	workshop.LocationName = db.LocationName
	workshop.LocationOnMap = db.LocationOnMap
	return workshop
}
