package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tombiers/estuary-backend/services"
)

// FindTags finds all tags containing a query
func FindTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Find Tags GET")
	vars := mux.Vars(r)
	query := vars["query"]
	tags := services.FindTags(query)
	json.NewEncoder(w).Encode(tags)
}
