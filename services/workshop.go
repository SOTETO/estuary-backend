package services

import (
	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllWorkshops return all workshops
func GetAllWorkshops() []models.Workshop {
	workshops, _ := daos.GetAllWorkshops()
	for i := range workshops {
		workshops[i].ContentUUIDs = FetchContentUUIDs(workshops[i].UUID)
	}
	// TODO: gather data from linked tables:
	// tags
	// likes
	// authors
	return workshops
}

// GetWorkshopByUUID return workshop with the given id
func GetWorkshopByUUID(uuid string) models.Workshop {
	workshop, _ := daos.GetWorkshopByUUID(uuid)
	// TODO: gather data from linked tables:
	// tags
	// likes
	// authors
	workshop.ContentUUIDs = FetchContentUUIDs(workshop.UUID)
	return workshop
}

// CreateWorkshop create a new workshop
func CreateWorkshop(workshop models.Workshop) models.Workshop {
	workshop.UUID = guuid.New().String()
	newWorkshop, _ := daos.CreateWorkshop(workshop)
	// TODO: create linked data:
	// tags
	// authors
	return newWorkshop
}

// UpdateWorkshop update the workshop with the given id to the given data
func UpdateWorkshop(uuid string, update models.Workshop) models.Workshop {
	var workshop, _ = daos.UpdateWorkshop(uuid, update)
	// TODO: update linked data
	// tags
	// authors
	workshop.ContentUUIDs = FetchContentUUIDs(workshop.UUID)
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
