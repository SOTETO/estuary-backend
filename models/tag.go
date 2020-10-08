package models

// TableName overrides the table name used by TagDB to `tag`
func (TagDB) TableName() string {
	return "tag"
}

// TagDB a tag record in the db
type TagDB struct {
	UUID string `gorm:"column:UUID; primaryKey"`
	name string `gorm:"column:name"`
}
