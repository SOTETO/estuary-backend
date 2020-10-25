package services

import (
	"fmt"

	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllWorkshops return all workshops
func GetAllWorkshops() []models.Workshop {
	workshops, _ := daos.GetAllWorkshops()
	for i := range workshops {
		workshops[i].ContentUUIDs = FetchContentUUIDs(workshops[i].UUID)
		workshops[i].Authors = daos.GetAuthorsFromWorkshop(workshops[i].UUID)
	}
	// TODO: gather data from linked tables:
	// tags
	// likes
	return workshops
}

// GetWorkshopByUUID return workshop with the given id
func GetWorkshopByUUID(uuid string) (models.Workshop, error) {
	workshop, err := daos.GetWorkshopByUUID(uuid)
	if err != nil {
		return workshop, err
	}
	// TODO: tags, likes

	workshop.Authors = daos.GetAuthorsFromWorkshop(uuid)
	workshop.ContentUUIDs = FetchContentUUIDs(workshop.UUID)
	return workshop, nil
}

// CreateWorkshop create a new workshop
func CreateWorkshop(workshop models.Workshop) models.Workshop {
	fmt.Println("Create a new workshop: ", workshop)
	workshop.UUID = guuid.New().String()
	newWorkshop, _ := daos.CreateWorkshop(workshop)
	// TODO: tags

	// authors
	errAuthors := daos.SetAuthors(workshop.UUID, workshop.Authors)
	if errAuthors == nil {
		newWorkshop.Authors = workshop.Authors
	}
	return newWorkshop
}

// UpdateWorkshop update the workshop with the given id to the given data
func UpdateWorkshop(uuid string, update models.Workshop) models.Workshop {
	fmt.Println("Update workshop:", update)
	var workshop, _ = daos.UpdateWorkshop(uuid, update)
	// TODO: tags

	// authors
	err1, err2 := daos.OverrideAuthors(uuid, update.Authors)
	if err1 == nil && err2 == nil {
		workshop.Authors = update.Authors
	}
	workshop.ContentUUIDs = FetchContentUUIDs(uuid)
	return workshop
}

// DeleteWorkshop delete the booking with the given id
func DeleteWorkshop(uuid string) {
	daos.DeleteWorkshop(uuid)
	// orphaned rows in tags, authors, likes are deleted cascadingly in the db
	// Content (ProblemStatements) from the workshop is NOT deleted
}

// FetchContentUUIDs fetch the UUIDs of contained contents
func FetchContentUUIDs(workshopUUID string) []string {
	var contents = daos.GetContentsFromWorkshop(workshopUUID)
	var UUIDs []string
	for _, content := range contents {
		UUIDs = append(UUIDs, content.UUID)
	}
	return UUIDs
}
