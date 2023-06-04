package types

type Student struct {
	StudID     int    `gorm:"primaryKey;autoIncrement"`
	StudFName  string `gorm:"column:stud_fname"`
	StudLName  string `gorm:"column:stud_lname"`
	StudStreet string `gorm:"column:stud_street"`
	StudCity   string `gorm:"column:stud_city"`
	StudZip    string `gorm:"column:stud_zip"`
}
