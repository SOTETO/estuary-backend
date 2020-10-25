package daos

import (
	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/models"
)

// CreateTag create a new tag, returns UUID
func CreateTag(name string) error {
	tagDB := models.TagDB{
		UUID: guuid.New().String(),
		Name: name,
	}
	err := db.Create(&tagDB).Error
	return err
}

// DeleteTag delete the tag with the given UUID from the db
func DeleteTag(tagUUID string) error {
	var tagDB models.TagDB
	err := db.Where("UUID = ?", tagUUID).First(&tagDB).Delete(&tagDB).Error
	return err
}

// GetTag get the Tag with the given UUID
func GetTag(tagUUID string) (models.Tag, error) {
	var tagDB models.TagDB
	err := db.Where("UUID = ?", tagUUID).First(&tagDB).Error
	return tagDB.FromDB(), err
}

// UpdateTag updates the Tag with the given UUID
func UpdateTag(tagUUID string, update models.Tag) (models.Tag, error) {
	var tagDB models.TagDB
	updateDB := update.ToDB()
	err := db.Where("UUID = ?", tagUUID).First(&tagDB).Updates(updateDB).Error
	return tagDB.FromDB(), err
}
