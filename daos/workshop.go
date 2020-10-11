package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetAllWorkshops returns all workshops
func GetAllWorkshops() ([]models.Workshop, error) {
	workshopsDB := []models.WorkshopDB{}
	workshops := []models.Workshop{}
	err := db.Find(&workshopsDB).Error
	if err == nil {
		for _, workshopDB := range workshopsDB {
			workshops = append(workshops, workshopDB.FromDB())
		}
	}
	return workshops, err
}

// GetWorkshopByUUID returns workshop with the given id
func GetWorkshopByUUID(uuid string) (models.Workshop, error) {
	var dbWorkshop models.WorkshopDB
	err := db.Where("UUID = ?", uuid).First(&dbWorkshop).Error
	return dbWorkshop.FromDB(), err
}

// CreateWorkshop creates a new workshop
func CreateWorkshop(workshop models.Workshop) (models.Workshop, error) {
	var dbWorkshop = workshop.ToDB()
	err = db.Create(dbWorkshop).Error
	return dbWorkshop.FromDB(), err
}

// UpdateWorkshop updates the workshop with the given id to the given data and return the updated workshop
func UpdateWorkshop(uuid string, update models.Workshop) (models.Workshop, error) {
	var dbUpdate = update.ToDB()
	var dbWorkshop models.WorkshopDB
	err := db.Where("UUID = ?", uuid).First(&dbWorkshop).Updates(dbUpdate).Error
	return dbWorkshop.FromDB(), err
}

// DeleteWorkshop delete the workshop with the given id
func DeleteWorkshop(uuid string) error {
	var dbWorkshop models.WorkshopDB
	err := db.Where("UUID = ?", uuid).First(&dbWorkshop).Delete(&dbWorkshop).Error
	return err
}
