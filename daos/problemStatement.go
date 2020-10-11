package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetAllProblemStatements returns all problemStatements
func GetAllProblemStatements() ([]models.ProblemStatement, error) {
	problemStatementsDB := []models.ProblemStatementDB{}
	problemStatements := []models.ProblemStatement{}
	err := db.Find(&problemStatementsDB).Error
	if err == nil {
		for _, ps := range problemStatementsDB {
			problemStatements = append(problemStatements, ps.FromDB())
		}
	}
	return problemStatements, err
}

// GetProblemStatementByUUID returns problemStatement with the given id
func GetProblemStatementByUUID(uuid string) (models.ProblemStatement, error) {
	var dbProblemStatement models.ProblemStatementDB
	err := db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Error
	return dbProblemStatement.FromDB(), err
}

// CreateProblemStatement creates a new problemStatement
func CreateProblemStatement(problemStatement models.ProblemStatement) (models.ProblemStatement, error) {
	dbProblemStatement := problemStatement.ToDB()
	err := db.Create(&dbProblemStatement).Error
	return dbProblemStatement.FromDB(), err
}

// UpdateProblemStatement updates the problemStatement with the given id to the given data and return the updated problemStatement
func UpdateProblemStatement(uuid string, update models.ProblemStatement) (models.ProblemStatement, error) {
	var dbProblemStatement models.ProblemStatementDB
	dbUpdate := update.ToDB()
	err := db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Updates(dbUpdate).Error
	return dbProblemStatement.FromDB(), err
}

// DeleteProblemStatement delete the problemStatement with the given id
func DeleteProblemStatement(uuid string) error {
	var dbProblemStatement models.ProblemStatementDB
	err := db.Where("Content_UUID = ?", uuid).First(&dbProblemStatement).Delete(&dbProblemStatement).Error
	return err
}
