package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetAllWorkshops returns all workshops
func GetAllWorkshops() []models.WorkshopDB {
	workshops := []models.WorkshopDB{}
	db.Find(&workshops)
	return workshops
}

// GetWorkshopByUUID returns workshop with the given id
func GetWorkshopByUUID(uuid string) models.WorkshopDB {
	var dbWorkshop models.WorkshopDB
	db.Where("UUID = ?", uuid).First(&dbWorkshop)
	return dbWorkshop
}

// CreateWorkshop creates a new workshop
func CreateWorkshop(workshop models.WorkshopDB) {
	db.Create(&workshop)
}

// UpdateWorkshop updates the workshop with the given id to the given data and return the updated workshop
func UpdateWorkshop(uuid string, workshop models.WorkshopDB) models.WorkshopDB {
	var dbWorkshop models.WorkshopDB
	db.Where("UUID = ?", uuid).First(&dbWorkshop).Updates(workshop)
	return dbWorkshop
}

// DeleteWorkshop delete the workshop with the given id
func DeleteWorkshop(uuid string) {
	var dbWorkshop models.WorkshopDB
	db.Where("UUID = ?", uuid).First(&dbWorkshop).Delete(&dbWorkshop)
}
