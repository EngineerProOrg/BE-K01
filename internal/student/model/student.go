// create table STUDENT (
//     STUD_ID int primary key,
//     STUD_LNAME varchar(50),
//     STUD_FNAME varchar(50),
//     STUD_STREET varchar(255),
//     STUD_CITY varchar(50),
//     STUD_ZIP varchar(10)
// )

package studentmodel

type Student struct { 
	Stud_id int `json:"stud_id" gorm:"column:STUD_ID;primaryKey"`
	Stud_lname string `json:"stud_lname" gorm:"column:STUD_LNAME"`
	Stud_fname string `json:"stud_fname" gorm:"column:STUD_FNAME"`
	Stud_street string `json:"stud_street" gorm:"column:STUD_STREET"`
	Stud_city string `json:"stud_city" gorm:"column:STUD_CITY"`
	Stud_zip string `json:"stud_zip" gorm:"column:STUD_ZIP"`
}

func (Student) TableName() string {
	return "STUDENT"
}