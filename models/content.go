package models

// TableName overrides the table name used by ConentDB to `Content`
func (ContentDB) TableName() string {
	return "content"
}

//ContentDB model for a record in the db
type ContentDB struct {
	UUID         string `gorm:"column:UUID; primaryKey"`
	Title        string `gorm:"column:title"`
	WorkshopUUID string `gorm:"column:workshop_UUID"`
}

// Content a general workshop Content object with json mapping
type Content struct {
	UUID           string        `json:"UUID"`
	Title          string        `json:"title"`
	WorkshopUUID   string        `json:"workshop_UUID"`
	LinkedContents []ContentLink `json:"linkedContent"`
	Likes          int           `json:"likes"`
}

// ToDB creates a ContentDB from a Content object
func (content Content) ToDB() ContentDB {
	var db ContentDB
	db.UUID = content.UUID
	db.Title = content.Title
	db.WorkshopUUID = content.WorkshopUUID
	return db
}

// FromDB returns a Content from a ContentDB
func (db ContentDB) FromDB() Content {
	var content Content
	content.UUID = db.UUID
	content.Title = db.Title
	content.WorkshopUUID = db.WorkshopUUID
	return content
}

// ToContentLinkDB returns ContentLinkDB objects for this workshop
func (content Content) ToContentLinkDB() []ContentLinkDB {
	var links []ContentLinkDB
	for _, value := range content.LinkedContents {
		var db ContentLinkDB
		db.LinkTag = value.LinkTag
		db.Content1 = content.UUID
		db.Content2 = value.ContentUUID
		links = append(links, db)
	}
	return links
}
