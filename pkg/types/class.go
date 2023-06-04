package types

type Class struct {
	ClassID   int       `gorm:"primaryKey;autoIncrement"`
	ClassName string    `gorm:"column:class_name"`
	ProfID    int       `gorm:"column:prof_id"`
	CourseID  int       `gorm:"column:course_id"`
	RoomID    int       `gorm:"column:room_id"`
	Professor Professor `gorm:"foreignKey:ProfID"`
	Course    Course    `gorm:"foreignKey:CourseID"`
	Room      Room      `gorm:"foreignKey:RoomID"`
}
