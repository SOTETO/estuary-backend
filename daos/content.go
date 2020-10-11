package daos

import (
	"fmt"

	"github.com/tombiers/estuary-backend/models"
)

// GetContentByUUID returns content with the given uuid
func GetContentByUUID(uuid string) models.Content {
	var dbContent models.ContentDB

	if err := db.Where("UUID = ?", uuid).First(&dbContent).Error; err != nil {
		fmt.Println("not found: ", err)
	}
	return dbContent.FromDB()
}

// CreateContent creates a new content
func CreateContent(content models.Content) models.Content {
	var contentDB = content.ToDB()
	db.Create(&contentDB)
	return contentDB.FromDB()
}

// UpdateContent updates the content with the given uuid to the given data and return the updated content
func UpdateContent(uuid string, content models.Content) models.Content {
	var dbContent models.ContentDB
	db.Where("UUID = ?", uuid).First(&dbContent).Updates(content.ToDB())
	return dbContent.FromDB()
}

// DeleteContent delete the content with the given uuid
func DeleteContent(uuid string) {
	var dbContent models.ContentDB
	db.Where("UUID = ?", uuid).First(&dbContent).Delete(&dbContent)
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
