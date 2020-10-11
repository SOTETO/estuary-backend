package models

// Tabler interface to set the associated db table name for a struct
type Tabler interface {
	TableName() string
}
