package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

func CreateNewBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: new-booking POST")
	// get the body of the POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var booking models.Booking
	json.Unmarshal(reqBody, &booking)
	fmt.Printf("%+v\n", booking)
	daos.Create(booking)
}

func ReturnAllBookings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBookings")
	bookings := daos.GetAll()
	json.NewEncoder(w).Encode(bookings)
}

func ReturnSingleBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Booking No: ", key)

	s, err := strconv.Atoi(key)
	if err == nil {
		var booking = daos.GetById(s)
		fmt.Println(booking)
		json.NewEncoder(w).Encode(booking)
	}
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking models.Booking
	json.Unmarshal(reqBody, &booking)

	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Update Booking No: ", key)

	s, err := strconv.Atoi(key)
	if err == nil {
		// only update if id from path and json match
		if s == booking.ID {
			var newBooking = daos.Update(s, booking)
			json.NewEncoder(w).Encode(newBooking)
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
		daos.Delete(s)
	}
}
