package models

// TableName overrides the table name used by workshopTagsDB to `workshoptags`
func (WorkshopTagDB) TableName() string {
	return "workshoptags"
}

// WorkshopTagDB a workshop tag record in the db
type WorkshopTagDB struct {
	ID           string `gorm:"column:ID; primaryKey"`
	WorkshopUUID string `gorm:"column:workshop_UUID"`
	TagUUID      string `gorm:"column:tags_UUID"`
}
