package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetContentByUUID returns content with the given uuid
func GetContentByUUID(uuid string) (models.Content, error) {
	var dbContent models.ContentDB
	err := db.Where("UUID = ?", uuid).First(&dbContent).Error
	return dbContent.FromDB(), err
}

// CreateContent creates a new content
func CreateContent(content models.Content) (models.Content, error) {
	var contentDB = content.ToDB()
	err := db.Create(&contentDB).Error
	return contentDB.FromDB(), err
}

// UpdateContent updates the content with the given uuid to the given data and return the updated content
func UpdateContent(uuid string, content models.Content) (models.Content, error) {
	var dbContent models.ContentDB
	err := db.Where("UUID = ?", uuid).First(&dbContent).Updates(content.ToDB()).Error
	return dbContent.FromDB(), err
}

// DeleteContent delete the content with the given uuid
func DeleteContent(uuid string) error {
	var dbContent models.ContentDB
	err := db.Where("UUID = ?", uuid).First(&dbContent).Delete(&dbContent).Error
	return err
}

// GetContentsFromWorkshop returns all contents belonging to the workshop with the given UUID
func GetContentsFromWorkshop(workshopUUID string) []models.Content {
	dbContent := []models.ContentDB{}
	db.Where("workshop_UUID = ?", workshopUUID).Find(&dbContent)

	content := []models.Content{}
	for _, value := range dbContent {
		content = append(content, value.FromDB())
	}
	return content
}
