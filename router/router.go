package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tombiers/estuary-backend/controllers"
	"github.com/tombiers/estuary-backend/views"
)

// say hi
func SayHi() {
	println("Hello World!")
}

func HandleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", views.HomePage)
	// bookings
	myRouter.HandleFunc("/new-booking", controllers.CreateNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", controllers.ReturnAllBookings).Methods("GET")
	myRouter.HandleFunc("/booking/{id}", controllers.ReturnSingleBooking).Methods("GET")
	myRouter.HandleFunc("/update-post/{id}", controllers.UpdateBooking).Methods("PUT")
	myRouter.HandleFunc("/delete/{id}", controllers.DeleteBooking).Methods("DELETE")
	// workshops
	myRouter.HandleFunc("/new-workshop", controllers.CreateNewWorkshop).Methods("POST")
	myRouter.HandleFunc("/all-workshops", controllers.AllWorkshops).Methods("GET")
	myRouter.HandleFunc("/workshop/{uuid}", controllers.SingleWorkshop).Methods("GET")
	myRouter.HandleFunc("/update-workshop/{uuid}", controllers.UpdateWorkshop).Methods("PUT")
	myRouter.HandleFunc("/delete-workshop/{uuid}", controllers.DeleteWorkshop).Methods("DELETE")
	// problemStatement
	myRouter.HandleFunc("/new-problemStatement", controllers.CreateNewProblemStatement).Methods("POST")
	myRouter.HandleFunc("/all-problemStatements", controllers.AllProblemStatements).Methods("GET")
	myRouter.HandleFunc("/problemStatement/{uuid}", controllers.SingleProblemStatement).Methods("GET")
	myRouter.HandleFunc("/update-problemStatement/{uuid}", controllers.UpdateProblemStatement).Methods("PUT")
	myRouter.HandleFunc("/delete-problemStatement/{uuid}", controllers.DeleteProblemStatement).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
