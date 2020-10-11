package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tombiers/estuary-backend/models"
	"github.com/tombiers/estuary-backend/services"
)

type uuidList struct {
	UUIDs []string `json:"UUIDs"`
}

// AllContentByUUID get all requested content by their uuid
func AllContentByUUID(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var uuids uuidList
	json.Unmarshal(reqBody, &uuids)
	var requestedContent models.ContentCollection
	requestedContent.ProblemStatements = services.GetMultipleProblemStatementByUUID(uuids.UUIDs)

	json.NewEncoder(w).Encode(requestedContent)
	/*
		problemStatement, err := services.GetProblemStatementByUUID(key)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(problemStatement)
	*/

}
