package types

type Professor struct {
	ProfID    int    `gorm:"primaryKey;autoIncrement"`
	ProfLName string `gorm:"column:prof_lname"`
	ProfFName string `gorm:"column:prof_fname"`
}
