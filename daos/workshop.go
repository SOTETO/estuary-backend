package daos

import "github.com/tombiers/estuary-backend/models"

// GetAllWorkshops returns all workshops
func GetAllWorkshops() []models.WorkshopDB {
	workshops := []models.WorkshopDB{}
	db.Find(&workshops)
	return workshops
}

// GetWorkshopByID returns workshop with the given id
func GetWorkshopByID(id int) models.WorkshopDB {
	var workshop models.WorkshopDB
	db.First(&workshop, id)
	return workshop
}

// CreateWorkshop creates a new workshop
func CreateWorkshop(workshop models.WorkshopDB) {
	db.Create(&workshop)
}

// UpdateWorkshop updates the workshop with the given id to the given data and return the updated workshop
func UpdateWorkshop(id int, workshop models.WorkshopDB) models.WorkshopDB {
	var dbWorkshop models.WorkshopDB
	db.Where("ID = ?", id).First(&dbWorkshop)
	dbWorkshop = workshop
	db.Save(&dbWorkshop)
	return dbWorkshop
}

// DeleteWorkshop delete the workshop with the given id
func DeleteWorkshop(id int) {
	var workshop models.WorkshopDB
	db.First(&workshop, id)
	db.Delete(&workshop)
}
