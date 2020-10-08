package models

// TableName overrides the table name used by workshopTagsDB to `workshoptags`
func (workshopTagsDB) TableName() string {
	return "workshoptags"
}

// TagDB a tag record in the db
type workshopTagsDB struct {
	UUID         string `gorm:"column:UUID; primaryKey"`
	WorkshopUUID string `gorm:"column:workshopUUID"`
	TagUUID      string `gorm:"column:tags_UUID"`
}
