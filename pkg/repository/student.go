package repository

import (
	"gorm.io/gorm"
)

type ClassGrade struct {
	ClassName string  `gorm:"column:Class Name"`
	AvgGrade  float64 `gorm:"column:Average Grade"`
	Grade     string  `gorm:"column:Grade"`
}

type StudentRepository interface {
	GetAvgGradeOfStudent() ([]StudentGrade, error)
	GetAvgGradeOfClass() ([]ClassGrade, error)
	GetAvgGradeOfCourse() ([]CourseGrade, error)
}

type StudentGrade struct {
	StudentName string  `gorm:"column:Student Name"`
	AvgGrade    float64 `gorm:"column:Average Grade"`
	Grade       string  `gorm:"column:Grade"`
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &StudentRepo{
		db: db,
	}
}

type StudentRepo struct {
	db *gorm.DB
}

type CourseGrade struct {
	CourseName string  `gorm:"column:Course Name"`
	AvgGrade   float64 `gorm:"column:Average Grade"`
	Grade      string  `gorm:"column:Grade"`
}

func (s StudentRepo) GetAvgGradeOfCourse() ([]CourseGrade, error) {
	var results []CourseGrade

	query := `
		SELECT c.course_name AS 'Course Name',
			AVG(G.Grade) AS 'Average Grade',
			CASE
				WHEN AVG(G.Grade) < 5 THEN 'Weak'
				WHEN AVG(G.Grade) >= 5 AND AVG(G.Grade) < 8 THEN 'Average'
				WHEN AVG(G.Grade) >= 8 THEN 'Good'
			END AS 'Grade'
		FROM courses c
		JOIN classes C2 ON c.course_id = C2.course_id
		JOIN Grade G ON C2.class_id = G.class_id
		GROUP BY c.course_id, c.course_name
	`

	err := s.db.Raw(query).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s StudentRepo) GetAvgGradeOfClass() ([]ClassGrade, error) {
	var results []ClassGrade

	query := `
		select c.class_name AS 'Class Name',
       AVG(G.Grade) AS 'Average Grade',
       CASE
           WHEN AVG(G.Grade) < 5 THEN 'Weak'
           WHEN AVG(G.Grade) >= 5 AND AVG(G.Grade) < 8 THEN 'Average'
           WHEN AVG(G.Grade) >= 8 THEN 'Good'
           END      AS 'Grade'
from classes c
         join Grade G on c.class_id = G.class_id
group by c.class_id, c.class_name
	`

	err := s.db.Raw(query).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s StudentRepo) GetAvgGradeOfStudent() ([]StudentGrade, error) {
	var results []StudentGrade

	query := `
		SELECT CONCAT(s.stud_fname, ' ', s.stud_lname) AS 'Student Name',
			AVG(G.Grade) AS 'Average Grade',
			CASE
				WHEN AVG(G.Grade) < 5 THEN 'Weak'
				WHEN AVG(G.Grade) >= 5 AND AVG(G.Grade) < 8 THEN 'Average'
				WHEN AVG(G.Grade) >= 8 THEN 'Good'
			END AS 'Grade'
		FROM students s
		JOIN Grade G ON s.stud_id = G.stud_id
		GROUP BY s.stud_id
	`

	err := s.db.Raw(query).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
