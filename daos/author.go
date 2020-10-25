package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// SetAuthors create new author record in the db
func SetAuthors(workshopUUID string, authors []models.Author) []error {
	authorsDB := []models.AuthorDB{}
	errs := []error{}
	for _, author := range authors {
		authorsDB = append(authorsDB, author.ToDB(workshopUUID))
		authorDB := author.ToDB(workshopUUID)
		err := db.Create(&authorDB).Error
		errs = append(errs, err)
	}
	return errs
}

// UpdateAuthor updates the author record for the given workshop
func UpdateAuthor(workshopUUID string, update models.Author) (models.Author, error) {
	var authorDB models.AuthorDB
	updateDB := update.ToDB(workshopUUID)
	err := db.Where("workshop_UUID = ? AND user_UUID = ?", workshopUUID, updateDB.UserUUID).First(&authorDB).Updates(updateDB).Error
	return authorDB.FromDB(), err
}

// RemoveAuthor delete the author record
func RemoveAuthor(workshopUUID string, author models.Author) error {
	authorDB := author.ToDB(workshopUUID)
	err := db.Where("workshop_UUID = ? AND user_UUID = ?", workshopUUID, authorDB.UserUUID).First(&authorDB).Delete(&authorDB).Error
	return err
}

// GetAuthorsFromWorkshop returns all authors belonging to the workshop with the given UUID
func GetAuthorsFromWorkshop(workshopUUID string) []models.Author {
	authorDB := []models.AuthorDB{}
	db.Where("workshop_UUID = ?", workshopUUID).Find(&authorDB)

	authors := []models.Author{}
	for _, value := range authorDB {
		authors = append(authors, value.FromDB())
	}
	return authors
}

// OverrideAuthors delete all existing authors for this workshop and set new ones
func OverrideAuthors(workshopUUID string, authors []models.Author) []error {
	var authorDB models.AuthorDB
	errDelete := db.Where("workshop_UUID = ?", workshopUUID).Delete(&authorDB).Error
	errSet := SetAuthors(workshopUUID, authors)
	errs := []error{errDelete}
	errs = append(errs, errSet...)
	return errs
}
