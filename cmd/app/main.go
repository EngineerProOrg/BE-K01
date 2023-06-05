package main

import (
	"assignment_1_gorm"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// main connects to database and does the exercises in sol-1.sql and sol-2.sql
func main() {
	// Connect to the MySQL database
	dsn := "quangmx:2511@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	initDatabase(db)

	// Những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
	related_stud_prof(db)

	// Những course (distinct) mà 1 professor cụ thể đang dạy
	teaching_courses(db, 1)

	// // TODO: Implement this
	// // Những course (distinct) mà 1 student cụ thể đang học
	// studying_courses(db, 1)

	// // TODO: Implement this
	// // điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
	// convert_grade(db)

	// // TODO: Implement this
	// // điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
	// stud_avg_grade(db, 1)

	// // TODO: Implement this
	// // điểm số trung bình của các class (quy ra lại theo chữ cái)
	// class_avg_grade(db)

	// // TODO: Implement this
	// // điểm số trung bình của các course (quy ra lại theo chữ cái)
	// course_avg_grade(db)

}

// initDatabase creates tables and adds records
func initDatabase(db *gorm.DB) {
	// Delete tables if existed
	db.Migrator().DropTable(&assignment_1_gorm.Professor{})
	db.Migrator().DropTable(&assignment_1_gorm.Student{})
	db.Migrator().DropTable(&assignment_1_gorm.Course{})
	db.Migrator().DropTable(&assignment_1_gorm.Class{})
	db.Migrator().DropTable(&assignment_1_gorm.Enroll{})
	fmt.Println("Tables deleted successfully...")

	// Create tables
	db.Migrator().CreateTable(&assignment_1_gorm.Professor{})
	db.Migrator().CreateTable(&assignment_1_gorm.Student{})
	db.Migrator().CreateTable(&assignment_1_gorm.Course{})
	db.Migrator().CreateTable(&assignment_1_gorm.Class{})
	db.Migrator().CreateTable(&assignment_1_gorm.Enroll{})
	fmt.Println("Tables created successfully...")

	// Create records
	var result *gorm.DB
	professors := []assignment_1_gorm.Professor{
		{FirstName: "Albus", LastName: "Dumbledore"},
		{FirstName: "Severus", LastName: "Snape"},
		{FirstName: "Alastor", LastName: "Moody"},
	}
	result = db.Create(&professors)
	if result.Error != nil {
		panic(result.Error)
	}

	students := []assignment_1_gorm.Student{
		{FirstName: "Harry", LastName: "Potter", Street: "Privet Drive 1", City: "London", Zip: "100000"},
		{FirstName: "Hermione", LastName: "Granger", Street: "Privet Drive 2", City: "London", Zip: "200000"},
		{FirstName: "Ron", LastName: "Weasly", Street: "Privet Drive 3", City: "London", Zip: "300000"},
	}
	result = db.Create(&students)
	if result.Error != nil {
		panic(result.Error)
	}

	courses := []assignment_1_gorm.Course{
		{Name: "Physical Education"},
		{Name: "Quidditch"},
		{Name: "Defence Against the Dark Arts"},
		{Name: "Potion"},
	}
	result = db.Create(&courses)
	if result.Error != nil {
		panic(result.Error)
	}

	classes := []assignment_1_gorm.Class{
		{Name: "PE1", ProfessorID: 1, CourseID: 1, RoomLoc: "001", RoomCap: "100"},
		{Name: "QD1", ProfessorID: 2, CourseID: 2, RoomLoc: "002", RoomCap: "200"},
		{Name: "DA1", ProfessorID: 3, CourseID: 3, RoomLoc: "003", RoomCap: "300"},
		{Name: "DA2", ProfessorID: 3, CourseID: 3, RoomLoc: "004", RoomCap: "400"},
		{Name: "PT1", ProfessorID: 3, CourseID: 4, RoomLoc: "005", RoomCap: "500"},
	}
	result = db.Create(&classes)
	if result.Error != nil {
		panic(result.Error)
	}

	enrolls := []assignment_1_gorm.Enroll{
		{StudentID: 1, ClassID: 2, Grade: "B"},
		{StudentID: 2, ClassID: 1, Grade: "A"},
		{StudentID: 2, ClassID: 2, Grade: "A"},
		{StudentID: 2, ClassID: 3, Grade: "A"},
		{StudentID: 3, ClassID: 1, Grade: "C"},
		{StudentID: 3, ClassID: 3, Grade: "D"},
		{StudentID: 1, ClassID: 4, Grade: "D"},
		{StudentID: 2, ClassID: 4, Grade: "D"},
		{StudentID: 3, ClassID: 4, Grade: "D"},
	}
	result = db.Create(&enrolls)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Records created successfully...")
}

// Những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
func related_stud_prof(db *gorm.DB) {
	type Record struct {
		StudentID   int `gorm:"column:stud_id"`
		ProfessorID int `gorm:"column:prof_id"`
		Count       int `gorm:"column:num_class"`
	}
	var records []Record
	query := `
select
    distinct stud_id,
    prof_id,
    count(*) as num_class
from
    Class
    join Enroll using (class_id)
group by
    stud_id,
    prof_id
order by
    stud_id;`
	db.Raw(query).Scan(&records)

	// Print to console
	fmt.Println("\nNhững cặp student-professor có dạy học nhau và số lớp mà họ có liên quan")
	for _, record := range records {
		fmt.Printf("%+v\n", record)
	}
}

// Những course (distinct) mà 1 professor cụ thể đang dạy
func teaching_courses(db *gorm.DB, prof_id int) {
	type Record struct {
		CourseID int `gorm:"column:course_id"`
	}
	var records []Record
	query := `
select
    distinct course_id
from
    Class
where
    prof_id = ?;`
	db.Raw(query, prof_id).Scan(&records)

	// Print to console
	fmt.Printf("\nNhững course (distinct) mà professor %d đang dạy\n", prof_id)
	for _, record := range records {
		fmt.Printf("%+v\n", record)
	}
}
