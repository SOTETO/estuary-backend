package models

// TableName overrides the table name used by AuthorDB to `authors`
func (AuthorDB) TableName() string {
	return "authors"
}

//AuthorDB model for a record in the db
type AuthorDB struct {
	UUID         string `gorm:"column:UUID; primaryKey"`
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
