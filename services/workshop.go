package services

import (
	"fmt"

	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllWorkshops return all workshops
func GetAllWorkshops() []models.Workshop {
	var workshopsDB = daos.GetAllWorkshops()
	var workshops []models.Workshop
	for _, workshop := range workshopsDB {
		workshops = append(workshops, workshop.FromDB())
	}
	// TODO: gather data from linked tables:
	// tags
	// likes
	// authors
	// content IDs
	return workshops
}

// GetWorkshopByUUID return workshop with the given id
func GetWorkshopByUUID(uuid string) models.Workshop {
	var workshop = daos.GetWorkshopByUUID(uuid).FromDB()
	// TODO: gather data from linked tables:
	// tags
	// likes
	// authors
	// content IDs
	return workshop
}

// CreateWorkshop create a new workshop
func CreateWorkshop(workshop models.Workshop) {
	fmt.Println("Service reporting, making a new Workshop with UUID: ", workshop.UUID, " date: ", workshop.Date, " teaser: ", workshop.Teaser, " authors: ", workshop.Authors)
	daos.CreateWorkshop(workshop.ToDB())
	// TODO: create linked data:
	// tags
	// authors
}

// UpdateWorkshop update the workshop with the given id to the given data
func UpdateWorkshop(uuid string, update models.Workshop) models.Workshop {
	var workshop = daos.UpdateWorkshop(uuid, update.ToDB()).FromDB()
	// TODO: update linked data
	// tags
	// authors
	// content IDs
	return workshop
}

// DeleteWorkshop delete the booking with the given id
func DeleteWorkshop(uuid string) {
	daos.DeleteWorkshop(uuid)
	// orphaned rows in tags, authors, likes are deleted cascadingly in the db
	// Content (ProblemStatements) from the workshop is NOT deleted
}
