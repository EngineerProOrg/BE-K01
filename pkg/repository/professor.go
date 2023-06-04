package repository

import (
	"github.com/EngineerProOrg/BE-K01/pkg/types"
	"gorm.io/gorm"
)

type ProfessorRepository interface {
	GetProfessorStudentClassNames() ([]QueryResult, error)
}

type ProfessorRepoDb struct {
	db *gorm.DB
}

type ProfessorStudentClass struct {
	ProfName    string `gorm:"column:professors Name"`
	StudentName string `gorm:"column:Student Name"`
	ClassName   string `gorm:"column:classes Name"`
}

type QueryResult struct {
	ProfessorsName string `gorm:"column:Professor Name"`
	StudentName    string `gorm:"column:Student Name"`
	ClassName      string `gorm:"column:Class Name"`
}

func (p ProfessorRepoDb) GetProfessorStudentClassNames() ([]QueryResult, error) {
	var result []QueryResult

	p.db.Model(&types.Professor{}).
		Select("CONCAT(professors.prof_fname, ' ', professors.prof_lname) AS 'Professor Name'",
			"CONCAT(students.stud_fname, ' ', students.stud_lname) AS 'Student Name'",
			"classes.class_name AS 'Class Name'").
		Joins("JOIN classes ON professors.prof_id = classes.prof_id").
		Joins("JOIN enrolls ON classes.class_id = enrolls.class_id").
		Joins("JOIN students ON enrolls.stud_id = students.stud_id").
		Scan(&result)

	if err := p.db.Error; err != nil {
		return nil, err
	}
	return result, nil
}

func NewProfessorRepository(db *gorm.DB) ProfessorRepository {
	return &ProfessorRepoDb{
		db: db,
	}
}
