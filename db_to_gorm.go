package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	// if used gorm.Model
	// auto include created_at, deleted_at, id, ...
	StudID     uint   `gorm:"column:stud_id;primaryKey;autoIncrement"`
	StudFName  string `gorm:"column:stud_fname;type:varchar(50);not null"`
	StudLName  string `gorm:"column:stud_lname;type:varchar(50);not null"`
	StudStreet string `gorm:"column:stud_street;type:varchar(255)"`
	StudZip    string `gorm:"column:stud_zip;type:varchar(10)"`
	StudCity   string `gorm:"column:stud_city;type:varchar(50)"`
}
type Professor struct {
	profID    uint    `gorm:"column:prof_id;primaryKey;autoIncrement"`
	profLName string  `gorm:"column:prof_lname;not null;type:varchar(50)"`
	profFName string  `gorm:"column:prof_fname;not null;type:varchar(50)`
	Classes   []Class `gorm:"foreignKey:ProfessorID"`
}
type Room struct {
	roomID  uint    `gorm:"column:room_id;primaryKey;autoIncrement"`
	roomLOC string  `gorm:"column:room_loc;not null;type:varchar(50)"`
	roomCAP string  `gorm:"column:room_cap;type:varchar(50)"`
	ClassID uint    `gorm:"column:class_id"`
	Classes []Class `gorm:"foreignKey:RoomID"`
}
type Course struct {
	courseID   uint    `gorm:"column:course_id;primaryKey;autoIncrement"`
	courseName string  `gorm:"column:course_name;type:varchar(255);not null`
	Classes    []Class `gorm:"foreignKey:CourseID"`
}
type Class struct {
	classID     uint     `gorm:"column:class_id;primaryKey;autoIncrement"`
	className   string   `gorm:"column:class_name;type:varchar(255);not null`
	ProfessorID uint     `gorm:"column:prof_id;not null"`
	CourseID    uint     `gorm:"column:course_id;not null"`
	RoomID      uint     `gorm:"column:room_id;not null"`
	Enrolls     []Enroll `gorm:"foreignKey:ClassID"`
}
type Enroll struct {
	StudentID uint   `gorm:"column:stud_id;primaryKey"`
	ClassID   uint   `gorm:"column:class_id;primaryKey"`
	grade     string `gorm:"column:grade;varchar(3);not null"`
}

func (Student) TableName() string {
	return "Student"
}
func (Professor) TableName() string {
	return "Professor"
}
func (Enroll) TableName() string {
	return "Enroll"
}
func (Class) TableName() string {
	return "Class"
}
func (Course) TableName() string {
	return "Course"
}
func (Room) TableName() string {
	return "Room"
}

func main() {
	dsn := "root:thobeogalaxy257@tcp(127.0.0.1:3307)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	var result1 []map[string]interface{}
	var result2 []map[string]interface{}
	var result3 []map[string]interface{}
	var result4 []map[string]interface{}
	var result5 []map[string]interface{}
	var result6 []map[string]interface{}
	var result7 []map[string]interface{}

	// get the student and prof and class they are in
	db.Table("Professor").
		Select("Student.stud_id, CONCAT(stud_fname, ' ', stud_lname) as stud_fullname, CONCAT(prof_fname, ' ', prof_lname) as prof_fullname, class_name").
		Joins("JOIN Class ON Professor.prof_id = Class.prof_id").
		Joins("JOIN Enroll On Class.class_id = Enroll.class_id").
		Joins("JOIN Student ON Enroll.stud_id = Student.stud_id").
		Find(&result1)
	jsonData, err := json.MarshalIndent(result1, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 1: ")
	fmt.Println(string(jsonData))
	// get the distinct class of professor
	db.Table("Professor").
		Distinct("course_name"). // distinct above select
		Select("Professor.prof_id, CONCAT(prof_fname, ' ', prof_lname) as prof_fullname, course_name").
		Joins("JOIN Class ON Professor.prof_id = Class.prof_id").
		Joins("JOIN Course ON Course.course_id = Class.course_id").
		Find(&result2)
	jsonData, err = json.MarshalIndent(result2, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 2: ")
	fmt.Println(string(jsonData))

	// the distinct course that student study
	db.Table("Student").
		Distinct("course_name"). // distinct above select
		Select("Student.stud_id, CONCAT(stud_fname, ' ', stud_lname) as stud_fullname, course_name").
		Joins("JOIN Enroll ON Student.stud_id = Enroll.stud_id").
		Joins("JOIN Class ON Class.class_id = Enroll.class_id").
		Joins("JOIN Course ON Class.course_id = Course.course_id").
		Find(&result3)
	jsonData, err = json.MarshalIndent(result3, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 3: ")
	fmt.Println(string(jsonData))

	// make A,B,C,D,E,F as 10,8,6,4,2,0
	db.Table("Enroll").
		Select(`stud_id, 
		CASE 
			WHEN grade = 'A' THEN 10 
			WHEN grade = 'B' THEN 8 
			WHEN grade = 'C' THEN 6 
			WHEN grade = 'D' THEN 4 
			WHEN grade = 'E' THEN 2 
			WHEN grade = '7' THEN 0 
		END AS grade`).
		Order("stud_id").
		Find(&result4)
	jsonData, err = json.MarshalIndent(result4, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 4: ")
	fmt.Println(string(jsonData))

	// academic ability from AVG(grade)
	db.Table("Enroll").
		Select(`stud_id, 
		AVG(CASE 
			WHEN grade = 'A' THEN 10 
			WHEN grade = 'B' THEN 8 
			WHEN grade = 'C' THEN 6 
			WHEN grade = 'D' THEN 4 
			WHEN grade = 'E' THEN 2 
			WHEN grade = '7' THEN 0 
		END) AS academic_ability`).
		Order("stud_id").
		Group("stud_id").
		Find(&result5)
	jsonData, err = json.MarshalIndent(result5, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 5: ")
	fmt.Println(string(jsonData))

	// điểm trung bình của các class
	db.Table("Enroll").
		Select(`class_id, class_name, 
		AVG(CASE 
			WHEN grade = 'A' THEN 10 
			WHEN grade = 'B' THEN 8 
			WHEN grade = 'C' THEN 6 
			WHEN grade = 'D' THEN 4 
			WHEN grade = 'E' THEN 2 
			WHEN grade = '7' THEN 0 
		END) AS avg_class_grade`).
		Joins("JOIN Class ON Class.class_id = Enroll.class_id").
		Order("class_id").
		Group("class_id").
		Find(&result6)
	jsonData, err = json.MarshalIndent(result6, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 6: ")
	fmt.Println(string(jsonData))

	// avg_grade of courses
	db.Table("Enroll").
		Select(`Course.course_id, course_name, 
		AVG(CASE 
			WHEN grade = 'A' THEN 10 
			WHEN grade = 'B' THEN 8 
			WHEN grade = 'C' THEN 6 
			WHEN grade = 'D' THEN 4 
			WHEN grade = 'E' THEN 2 
			WHEN grade = '7' THEN 0 
		END) AS avg_course_grade`).
		Joins("JOIN Class ON Class.class_id = Enroll.class_id").
		Joins("JOIN Course On Course.course_id = Class.course_id").
		Order("Course.course_id").
		Group("Course.course_id").
		Find(&result7)
	jsonData, err = json.MarshalIndent(result7, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Query 7: ")
	fmt.Println(string(jsonData))
}
