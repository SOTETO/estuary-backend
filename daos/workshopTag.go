package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// AddTag add a tag to the workshop
func AddTag(workshopUUID string, tagUUID string) error {
	workshopTagDB := models.WorkshopTagDB{
		WorkshopUUID: workshopUUID,
		TagUUID:      tagUUID,
	}
	err := db.Create(&workshopTagDB).Error
	return err
}

// RemoveTag delete the tag from the workshop
func RemoveTag(workshopUUID string, tagUUID string) error {
	var workshopTagDB models.WorkshopTagDB
	err := db.Where("workshop_UUID = ? AND tags_UUID = ?", workshopUUID, tagUUID).First(&workshopTagDB).Delete(&workshopTagDB).Error
	return err
}

// GetTagsFromWorkshop returns all tags belonging to the workshop with the given UUID
func GetTagsFromWorkshop(workshopUUID string) ([]models.Tag, error) {
	workshopTagsDB := []models.WorkshopTagDB{}
	err := db.Where("workshop_UUID = ?", workshopUUID).Find(&workshopTagsDB).Error

	tags := []models.Tag{}
	for _, value := range workshopTagsDB {
		var tagDB models.TagDB
		db.Where("UUID = ?", value.TagUUID).Find(&tagDB)
		tags = append(tags, tagDB.FromDB())
	}
	return tags, err
}

// OverrideTags delete all existing tags for this workshop and set new ones
func OverrideTags(workshopUUID string, tags []string) []error {
	var authorDB models.AuthorDB
	var errs []error
	errDelete := db.Where("workshop_UUID = ?", workshopUUID).Delete(&authorDB).Error
	errs = append(errs, errDelete)
	for _, tag := range tags {
		err := AddTag(workshopUUID, tag)
		errs = append(errs, err)
	}
	return errs
}
