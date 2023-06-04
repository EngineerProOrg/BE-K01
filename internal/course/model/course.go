// create table COURSE (
//     COURSE_ID int primary key,
//     COURSE_NAME varchar(255)
// )


package coursemodel

type Course struct {
	Course_id int `json:"course_id" gorm:"column:COURSE_ID;primaryKey"`
	Course_name string `json:"course_name" gorm:"column:COURSE_NAME"`
}

func (Course) TableName() string {
	return "COURSE"
}