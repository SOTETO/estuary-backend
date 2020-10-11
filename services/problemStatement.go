package services

import (
	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllProblemStatements return all ProblemStatements
func GetAllProblemStatements() []models.ProblemStatement {
	var problemStatementsDB = daos.GetAllProblemStatements()
	var problemStatements []models.ProblemStatement
	for _, ps := range problemStatementsDB {
		var problemStatement = ps.FromDB()
		problemStatement.Content = daos.GetContentByUUID(problemStatement.UUID).FromDB()
		problemStatements = append(problemStatements, problemStatement)
	}
	return problemStatements
}

// GetProblemStatementByUUID return ProblemStatement with the given uuid
func GetProblemStatementByUUID(uuid string) models.ProblemStatement {
	var problemStatement = daos.GetProblemStatementByUUID(uuid).FromDB()
	problemStatement.Content = daos.GetContentByUUID(uuid).FromDB()
	return problemStatement
}

// CreateProblemStatement create a new ProblemStatement
func CreateProblemStatement(problemStatement models.ProblemStatement) models.ProblemStatement {
	problemStatement.UUID = guuid.New().String()
	var content = daos.CreateContent(problemStatement.Content.ToDB()).FromDB()
	var newProblemStatement = daos.CreateProblemStatement(problemStatement.ToDB()).FromDB()
	newProblemStatement.Content = content
	return newProblemStatement
}

// UpdateProblemStatement update the problemStatement with the given uuid to the given data
func UpdateProblemStatement(uuid string, update models.ProblemStatement) models.ProblemStatement {
	var problemStatement = daos.UpdateProblemStatement(uuid, update.ToDB()).FromDB()
	problemStatement.Content = daos.UpdateContent(uuid, problemStatement.Content.ToDB()).FromDB()
	return problemStatement
}

// DeleteProblemStatement delete the ProblemStatement with the given uuid
func DeleteProblemStatement(uuid string) {
	daos.DeleteProblemStatement(uuid)
	daos.DeleteContent(uuid)
}
