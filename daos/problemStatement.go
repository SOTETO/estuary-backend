package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetAllProblemStatements returns all problemStatements
func GetAllProblemStatements() []models.ProblemStatementDB {
	problemStatements := []models.ProblemStatementDB{}
	db.Find(&problemStatements)
	return problemStatements
}

// GetProblemStatementByUUID returns problemStatement with the given id
func GetProblemStatementByUUID(uuid string) models.ProblemStatementDB {
	var dbProblemStatement models.ProblemStatementDB
	db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement)
	return dbProblemStatement
}

// CreateProblemStatement creates a new problemStatement
func CreateProblemStatement(problemStatement models.ProblemStatementDB) models.ProblemStatementDB {
	db.Create(&problemStatement)
	return problemStatement
}

// UpdateProblemStatement updates the problemStatement with the given id to the given data and return the updated problemStatement
func UpdateProblemStatement(uuid string, problemStatement models.ProblemStatementDB) models.ProblemStatementDB {
	var dbProblemStatement models.ProblemStatementDB
	db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Updates(problemStatement)
	return dbProblemStatement
}

// DeleteProblemStatement delete the problemStatement with the given id
func DeleteProblemStatement(uuid string) {
	var dbProblemStatement models.ProblemStatementDB
	db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Delete(&dbProblemStatement)
}
