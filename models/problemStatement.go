package models

// TableName overrides the table name used by ProblemStatementDB to `problemstatement`
func (ProblemStatementDB) TableName() string {
	return "problemstatement"
}

//ProblemStatementDB model for a record in the db
type ProblemStatementDB struct {
	UUID    string `gorm:"column:Content_UUID; primaryKey"`
	Iam     string `gorm:"column:iAM"`
	Iwant   string `gorm:"column:iWant"`
	But     string `gorm:"column:but"`
	Because string `gorm:"column:because"`
	Feel    string `gorm:"column:feel"`
}

// ProblemStatement the ProblemStatement Content type
type ProblemStatement struct {
	Content
	Iam     string `json:"iAM"`
	Iwant   string `json:"iWant"`
	But     string `json:"but"`
	Because string `json:"because"`
	Feel    string `json:"feel"`
}

// ToDB creates a ContentDB from a Content object
func (ps ProblemStatement) ToDB() ProblemStatementDB {
	var db ProblemStatementDB
	db.UUID = ps.Content.UUID
	db.Iam = ps.Iam
	db.Iwant = ps.Iwant
	db.But = ps.But
	db.Because = ps.Because
	db.Feel = ps.Feel
	return db
}

// FromDB returns a Content from a ContentDB
func (db ProblemStatementDB) FromDB() ProblemStatement {
	var ps ProblemStatement
	ps.Content.UUID = db.UUID
	ps.Iam = db.Iam
	ps.Iwant = db.Iwant
	ps.But = db.But
	ps.Because = db.Because
	ps.Feel = db.Feel
	return ps
}
