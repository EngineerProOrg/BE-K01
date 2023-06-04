package repository

import (
	"github.com/EngineerProOrg/BE-K01/pkg/types"
	"gorm.io/gorm"
)

type CourseRepository interface {
	GetCourseWithProfessorTeaching() ([]string, error)
	GetCourseWithStudentStudying() ([]string, error)
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepo{
		db: db,
	}
}

type CourseRepo struct {
	db *gorm.DB
}

func (c CourseRepo) GetCourseWithStudentStudying() ([]string, error) {
	var courses []string
	c.db.Model(&types.Course{}).
		Select("DISTINCT(course_name)").
		Joins("JOIN Class ON Course.course_id = Class.course_id").
		Joins("JOIN Enroll ON Class.class_id = Enroll.class_id").
		Scan(&courses)
	if err := c.db.Error; err != nil {
		return nil, err
	}
	return courses, nil
}

type a struct {
	s int
}

func (c CourseRepo) GetCourseWithProfessorTeaching() ([]string, error) {
	var courses []string
	c.db.Model(&types.Course{}).
		Select("DISTINCT(course_name)").
		Joins("JOIN classes ON courses.course_id = classes.course_id").
		Scan(&courses)

	if err := c.db.Error; err != nil {
		return nil, err
	}
	return courses, nil
}
