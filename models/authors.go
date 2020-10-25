package models

// TableName overrides the table name used by AuthorDB to `authors`
func (AuthorDB) TableName() string {
	return "authors"
}

//AuthorDB model for a record in the db
type AuthorDB struct {
	ID           int    `gorm:"column:ID; primaryKey"`
	WorkshopUUID string `gorm:"column:workshop_UUID"`
	UserUUID     string `gorm:"column:user_UUID"`
	Visible      bool   `gorm:"column:visible"`
}

// Author a model for a workshop author
type Author struct {
	UserUUID string `json:"userUUID"`
	Visible  bool   `json:"visible"`
}

// FromDB returns an Author from an AuthorDB
func (db AuthorDB) FromDB() Author {
	var author Author
	author.UserUUID = db.UserUUID
	author.Visible = db.Visible
	return author
}

// ToDB tales the provided workshop UUID and creates a new AuthorDB object
func (author Author) ToDB(workshopUUID string) AuthorDB {
	var db AuthorDB
	db.WorkshopUUID = workshopUUID
	db.UserUUID = author.UserUUID
	db.Visible = author.Visible
	return db
}
