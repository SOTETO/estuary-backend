package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tombiers/estuary-backend/models"
	"github.com/tombiers/estuary-backend/services"
)

func CreateNewBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: new-booking POST")
	// get the body of the POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var booking models.BookingDTO
	json.Unmarshal(reqBody, &booking)
	services.Create(models.BookingFromDTO(booking))
}

func ReturnAllBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBookings")
	bookings := services.GetAll()
	bookingsDTO := []models.BookingDTO{}

	for i := 0; i < len(bookings); i++ {
		bookingsDTO = append(bookingsDTO, bookings[i].ToDTO())
	}

	json.NewEncoder(w).Encode(bookingsDTO)
}

func ReturnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Booking No: ", key)

	s, err := strconv.Atoi(key)
	if err == nil {
		var booking = services.GetByID(s)
		fmt.Println(booking)
		json.NewEncoder(w).Encode(booking.ToDTO())
	}
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking models.BookingDTO
	json.Unmarshal(reqBody, &booking)

	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Update Booking No: ", key)

	s, err := strconv.Atoi(key)
	if err == nil {
		// only update if id from path and json match
		if s == booking.ID {
			var newBooking = services.Update(s, models.BookingFromDTO(booking))
			json.NewEncoder(w).Encode(newBooking.ToDTO())
		}
	}
}

// delete a booking with the given id
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Delete Booking No: ", key)
	s, err := strconv.Atoi(key)
	if err == nil {
		services.Delete(s)
	}
}
