package services

import (
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

// FindTags returns all Tags containing the query in their name
func FindTags(query string) []models.Tag {

	tags, _ := daos.FindTags(query)
	return tags
}
