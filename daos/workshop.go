package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetAllWorkshops returns all workshops
func GetAllWorkshops() []models.Workshop {
	workshopsDB := []models.WorkshopDB{}
	db.Find(&workshopsDB)

	var workshops []models.Workshop
	for _, workshopDB := range workshopsDB {
		workshops = append(workshops, workshopDB.FromDB())
	}
	return workshops
}

// GetWorkshopByUUID returns workshop with the given id
func GetWorkshopByUUID(uuid string) models.Workshop {
	var dbWorkshop models.WorkshopDB
	db.Where("UUID = ?", uuid).First(&dbWorkshop)
	return dbWorkshop.FromDB()
}

// CreateWorkshop creates a new workshop
func CreateWorkshop(workshop models.Workshop) models.Workshop {
	var dbWorkshop = workshop.ToDB()
	db.Create(dbWorkshop)
	return dbWorkshop.FromDB()
}

// UpdateWorkshop updates the workshop with the given id to the given data and return the updated workshop
func UpdateWorkshop(uuid string, update models.Workshop) models.Workshop {
	var dbUpdate = update.ToDB()
	var dbWorkshop models.WorkshopDB
	db.Where("UUID = ?", uuid).First(&dbWorkshop)
	if dbWorkshop.UUID != "" {
		db.Where("UUID = ?", uuid).First(&dbWorkshop).Updates(dbUpdate)
	}
	return dbWorkshop.FromDB()
}

// DeleteWorkshop delete the workshop with the given id
func DeleteWorkshop(uuid string) {
	var dbWorkshop models.WorkshopDB
	db.Where("UUID = ?", uuid).First(&dbWorkshop).Delete(&dbWorkshop)
}
