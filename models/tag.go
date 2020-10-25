package models

// TableName overrides the table name used by TagDB to `tag`
func (TagDB) TableName() string {
	return "tag"
}

// TagDB a tag record in the db
type TagDB struct {
	UUID string `gorm:"column:UUID; primaryKey"`
	Name string `gorm:"column:name"`
}

// Tag a tag for workshops
type Tag struct {
	UUID string `json:"UUID"`
	Name string `json:"name"`
}

// FromDB takes a TabDB and returns a Tag
func (tagDB TagDB) FromDB() Tag {
	tag := Tag{
		UUID: tagDB.UUID,
		Name: tagDB.Name,
	}
	return tag
}

// ToDB takes a Tag and returns a TabDB
func (tag Tag) ToDB() TagDB {
	tagDB := TagDB{
		UUID: tag.UUID,
		Name: tag.Name,
	}
	return tagDB
}
