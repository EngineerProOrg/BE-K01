package types

type Enroll struct {
	StudID  int     `gorm:"primaryKey"`
	ClassID int     `gorm:"primaryKey"`
	Grade   string  `gorm:"column:grade;index"`
	Student Student `gorm:"foreignKey:StudID"`
	Class   Class   `gorm:"foreignKey:ClassID"`
}
