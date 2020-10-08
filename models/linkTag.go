package models

// TableName overrides the table name used by LinkTagDB to `linktag`
func (LinkTagDB) TableName() string {
	return "linktag"
}

// LinkTagDB a ContentLink record in the db
type LinkTagDB struct {
	UUID string `gorm:"column:UUID; primaryKey"`
	Name string `gorm:"column:name"`
}
