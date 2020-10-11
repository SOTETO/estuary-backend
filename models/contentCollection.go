package models

// ContentCollection a collection to hold all types of contents,
type ContentCollection struct {
	ProblemStatements []ProblemStatement `json:"ProblemStatements"`
	// add new content types here
}
