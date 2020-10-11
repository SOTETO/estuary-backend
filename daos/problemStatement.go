package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetAllProblemStatements returns all problemStatements
func GetAllProblemStatements() []models.ProblemStatement {
	problemStatementsDB := []models.ProblemStatementDB{}
	db.Find(&problemStatementsDB)

	problemStatements := []models.ProblemStatement{}
	for _, ps := range problemStatementsDB {
		problemStatements = append(problemStatements, ps.FromDB())
	}

	return problemStatements
}

// GetProblemStatementByUUID returns problemStatement with the given id
func GetProblemStatementByUUID(uuid string) (models.ProblemStatement, error) {
	var dbProblemStatement models.ProblemStatementDB
	err := db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Error
	return dbProblemStatement.FromDB(), err
}

// CreateProblemStatement creates a new problemStatement
func CreateProblemStatement(problemStatement models.ProblemStatement) models.ProblemStatement {
	dbProblemStatement := problemStatement.ToDB()
	db.Create(&dbProblemStatement)
	return dbProblemStatement.FromDB()
}

// UpdateProblemStatement updates the problemStatement with the given id to the given data and return the updated problemStatement
func UpdateProblemStatement(uuid string, update models.ProblemStatement) models.ProblemStatement {
	var dbProblemStatement models.ProblemStatementDB
	dbUpdate := update.ToDB()
	db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Updates(dbUpdate)
	return dbProblemStatement.FromDB()
}

// DeleteProblemStatement delete the problemStatement with the given id
func DeleteProblemStatement(uuid string) {
	var dbProblemStatement models.ProblemStatementDB
	db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Delete(&dbProblemStatement)
}
