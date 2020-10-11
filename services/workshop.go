package services

import (
	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllWorkshops return all workshops
func GetAllWorkshops() []models.Workshop {
	var workshopsDB = daos.GetAllWorkshops()
	var workshops []models.Workshop
	for _, workshopDB := range workshopsDB {
		var workshop = workshopDB.FromDB()
		workshop.ContentUUIDs = FetchContentUUIDs(workshop)
		workshops = append(workshops, workshop)
	}
	// TODO: gather data from linked tables:
	// tags
	// likes
	// authors
	return workshops
}

// GetWorkshopByUUID return workshop with the given id
func GetWorkshopByUUID(uuid string) models.Workshop {
	var workshop = daos.GetWorkshopByUUID(uuid).FromDB()
	// TODO: gather data from linked tables:
	// tags
	// likes
	// authors
	workshop.ContentUUIDs = FetchContentUUIDs(workshop)
	return workshop
}

// CreateWorkshop create a new workshop
func CreateWorkshop(workshop models.Workshop) models.Workshop {
	workshop.UUID = guuid.New().String()
	var workshopDB = daos.CreateWorkshop(workshop.ToDB())
	// TODO: create linked data:
	// tags
	// authors
	return workshopDB.FromDB()
}

// UpdateWorkshop update the workshop with the given id to the given data
func UpdateWorkshop(uuid string, update models.Workshop) models.Workshop {
	var workshop = daos.UpdateWorkshop(uuid, update.ToDB()).FromDB()
	// TODO: update linked data
	// tags
	// authors
	workshop.ContentUUIDs = FetchContentUUIDs(workshop)
	return workshop
}

// DeleteWorkshop delete the booking with the given id
func DeleteWorkshop(uuid string) {
	daos.DeleteWorkshop(uuid)
	// orphaned rows in tags, authors, likes are deleted cascadingly in the db
	// Content (ProblemStatements) from the workshop is NOT deleted
}

// FetchContentUUIDs fetch the UUIDs of contained contents
func FetchContentUUIDs(workshop models.Workshop) []string {
	var contents = daos.GetContentsFromWorkshop(workshop.UUID)
	var UUIDs []string
	for _, content := range contents {
		UUIDs = append(UUIDs, content.UUID)
	}
	return UUIDs
}
