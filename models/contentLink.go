package models

// TableName overrides the table name used by ContentLinkDB to `contenttlink`
func (ContentLinkDB) TableName() string {
	return "contenttlink"
}

// ContentLinkDB a ContentLink record in the db
type ContentLinkDB struct {
	ID       int    `gorm:"column:ID; primaryKey"`
	LinkTag  string `gorm:"column:linkTag_UUID"`
	Content1 string `gorm:"column:Content1_UUID"`
	Content2 string `gorm:"column:Content2_UUID"`
}

// ContentLink a ContentLink
type ContentLink struct {
	LinkTag     string `json:"linkTag"`
	ContentUUID string `json:"contentUUID"`
}
