package services

import (
	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllProblemStatements return all ProblemStatements
func GetAllProblemStatements() []models.ProblemStatement {
	problemStatements := daos.GetAllProblemStatements()
	for i := range problemStatements {
		problemStatements[i].Content = daos.GetContentByUUID(problemStatements[i].UUID)
	}
	return problemStatements
}

// GetProblemStatementByUUID return ProblemStatement with the given uuid
func GetProblemStatementByUUID(uuid string) models.ProblemStatement {
	problemStatement, _ := daos.GetProblemStatementByUUID(uuid)
	problemStatement.Content = daos.GetContentByUUID(uuid)
	return problemStatement
}

// CreateProblemStatement create a new ProblemStatement
func CreateProblemStatement(problemStatement models.ProblemStatement) models.ProblemStatement {
	problemStatement.UUID = guuid.New().String()
	var content = daos.CreateContent(problemStatement.Content)
	var newProblemStatement = daos.CreateProblemStatement(problemStatement)
	newProblemStatement.Content = content
	return newProblemStatement
}

// UpdateProblemStatement update the problemStatement with the given uuid to the given data
func UpdateProblemStatement(uuid string, update models.ProblemStatement) models.ProblemStatement {
	var problemStatement = daos.UpdateProblemStatement(uuid, update)
	problemStatement.Content = daos.UpdateContent(uuid, problemStatement.Content)
	return problemStatement
}

// DeleteProblemStatement delete the ProblemStatement with the given uuid
func DeleteProblemStatement(uuid string) {
	daos.DeleteProblemStatement(uuid)
	daos.DeleteContent(uuid)
}
