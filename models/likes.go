package models

// TableName overrides the table name used by LikesDB to `likes`
func (LikesDB) TableName() string {
	return "likes"
}

// LikesDB a ContentLink record in the db
type LikesDB struct {
	UUID       string `gorm:"column:UUID; primaryKey"`
	UserUUID   string `gorm:"column:user_UUID"`
	ConentUUID string `gorm:"column:Content_ID"`
}
