package daos

import "github.com/tombiers/estuary-backend/models"

// GetAllWorkshops returns all workshops
func GetAllWorkshops() []models.Workshop {
	workshops := []models.Workshop{}
	db.Find(&workshops)
	return workshops
}

// GetWorkshopByID returns workshop with the given id
func GetWorkshopByID(id int) models.Workshop {
	var workshop models.Workshop
	db.First(&workshop, id)
	return workshop
}

// CreateWorkshop creates a new workshop
func CreateWorkshop(workshop models.Workshop) {
	db.Create(&workshop)
}

// UpdateWorkshop updates the workshop with the given id to the given data and return the updated workshop
func UpdateWorkshop(id int, workshop models.Workshop) models.Workshop {
	var dbWorkshop models.Workshop
	db.Where("ID = ?", id).First(&dbWorkshop)
	dbWorkshop = workshop
	db.Save(&dbWorkshop)
	return dbWorkshop
}

// DeleteWorkshop delete the workshop with the given id
func DeleteWorkshop(id int) {
	var workshop models.Workshop
	db.First(&workshop, id)
	db.Delete(&workshop)
}
