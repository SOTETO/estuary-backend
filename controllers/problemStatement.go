package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tombiers/estuary-backend/models"
	"github.com/tombiers/estuary-backend/services"
)

// CreateNewProblemStatement create a new problemStatement
func CreateNewProblemStatement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateNewProblemStatement POST")
	// get the body of the POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var problemStatement models.ProblemStatement
	json.Unmarshal(reqBody, &problemStatement)
	var newProblemStatement = services.CreateProblemStatement(problemStatement)
	json.NewEncoder(w).Encode(newProblemStatement)
}

// AllProblemStatements returns all known problemStatements
func AllProblemStatements(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: AllProblemStatements")
	problemStatements := services.GetAllProblemStatements()
	json.NewEncoder(w).Encode(problemStatements)
}

// SingleProblemStatement returns a single problemStatement
func SingleProblemStatement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uuid"]
	fmt.Println("Endpoint Hit: ProblemStatement No: ", key)

	var problemStatement = services.GetProblemStatementByUUID(key)
	fmt.Println(problemStatement)
	json.NewEncoder(w).Encode(problemStatement)
}

// UpdateProblemStatement updates the problemStatement with the given uuid
// reuqested uuid and uuid in the update data have to match
func UpdateProblemStatement(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var problemStatement models.ProblemStatement
	json.Unmarshal(reqBody, &problemStatement)

	vars := mux.Vars(r)
	key := vars["uuid"]
	fmt.Println("Endpoint Hit: Update ProblemStatement No: ", key)

	// only update if id from path and json match
	if key == problemStatement.UUID {
		var newProblemStatement = services.UpdateProblemStatement(key, problemStatement)
		json.NewEncoder(w).Encode(newProblemStatement)
	}

}

// DeleteProblemStatement deletes the problemStatement with the given uui
func DeleteProblemStatement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uuid"]
	fmt.Println("Endpoint Hit: Delete ProblemStatement No: ", key)
	services.DeleteProblemStatement(key)
}
