package types

type Course struct {
	CourseID   int    `gorm:"primaryKey;autoIncrement"`
	CourseName string `gorm:"column:course_name"`
}
