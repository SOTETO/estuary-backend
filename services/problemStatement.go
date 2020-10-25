package services

import (
	guuid "github.com/google/uuid"
	"github.com/tombiers/estuary-backend/daos"
	"github.com/tombiers/estuary-backend/models"
)

//GetAllProblemStatements return all ProblemStatements
func GetAllProblemStatements() []models.ProblemStatement {
	problemStatements, _ := daos.GetAllProblemStatements()
	for i := range problemStatements {
		problemStatements[i].Content, _ = daos.GetContentByUUID(problemStatements[i].UUID)
	}
	return problemStatements
}

// GetProblemStatementByUUID return ProblemStatement with the given uuid
func GetProblemStatementByUUID(uuid string) (models.ProblemStatement, error) {
	problemStatement, err := daos.GetProblemStatementByUUID(uuid)
	problemStatement.Content, _ = daos.GetContentByUUID(uuid)
	return problemStatement, err
}

// GetMultipleProblemStatementByUUID get all problemStatements from a list of UUIDs, empty list if none is found
func GetMultipleProblemStatementByUUID(uuids []string) []models.ProblemStatement {
	problemStatements := []models.ProblemStatement{}
	for _, uuid := range uuids {
		ps, err := GetProblemStatementByUUID(uuid)
		if err == nil {
			problemStatements = append(problemStatements, ps)
		}
	}
	return problemStatements
}

// CreateProblemStatement create a new ProblemStatement
func CreateProblemStatement(problemStatement models.ProblemStatement) models.ProblemStatement {
	problemStatement.UUID = guuid.New().String()
	var content, _ = daos.CreateContent(problemStatement.Content)
	var newProblemStatement, _ = daos.CreateProblemStatement(problemStatement)
	newProblemStatement.Content = content
	return newProblemStatement
}

// UpdateProblemStatement update the problemStatement with the given uuid to the given data
func UpdateProblemStatement(uuid string, update models.ProblemStatement) models.ProblemStatement {
	var problemStatement, _ = daos.UpdateProblemStatement(uuid, update)
	problemStatement.Content, _ = daos.UpdateContent(uuid, update.Content)
	return problemStatement
}

// DeleteProblemStatement delete the ProblemStatement with the given uuid
func DeleteProblemStatement(uuid string) {
	daos.DeleteProblemStatement(uuid)
	daos.DeleteContent(uuid)
}
