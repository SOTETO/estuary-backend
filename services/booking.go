package services

import (
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAll : get all known bookings
func GetAll() []models.Booking {
	return daos.GetAll()
}

// GetByID returns booking with the given id
func GetByID(id int) models.Booking {
	return daos.GetById(id)
}

// Create creates a new booking
func Create(booking models.Booking) {
	daos.Create(booking)
}

// Update the booking with the given id to the given data
func Update(id int, booking models.Booking) models.Booking {
	return daos.Update(id, booking)
}

// Delete the booking with the given id
func Delete(id int) {
	daos.Delete(id)
}
