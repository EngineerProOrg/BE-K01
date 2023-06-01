package assignment_1_gorm

// CREATE TABLE Professor (
//
//	prof_id INT AUTO_INCREMENT,
//	prof_fname VARCHAR(50),
//	prof_lname VARCHAR(50),
//	PRIMARY KEY (prof_id)
//
// );
type Professor struct {
	ID        uint32 `gorm:"column:prof_id;type:int(32);primary_key;auto_increment"`
	FirstName string `gorm:"column:prof_fname;type:varchar(50)"`
	LastName  string `gorm:"column:prof_lname;type:varchar(50)"`
}

func (Professor) TableName() string {
	return "Professor"
}

// CREATE TABLE Student (
//
//	stud_id INT AUTO_INCREMENT,
//	stud_fname VARCHAR(50),
//	stud_lname VARCHAR(50),
//	stud_street VARCHAR(255),
//	stud_city VARCHAR(50),
//	stud_zip VARCHAR(10),
//	PRIMARY KEY (stud_id)
//
// );
type Student struct {
	ID        uint32 `gorm:"column:stud_id;type:int(32);primary_key;auto_increment"`
	FirstName string `gorm:"column:stud_fname;type:varchar(50)"`
	LastName  string `gorm:"column:stud_lname;type:varchar(50)"`
	Street    string `gorm:"column:stud_street;type:varchar(255)"`
	City      string `gorm:"column:stud_city;type:varchar(50)"`
	Zip       string `gorm:"column:stud_zip;type:varchar(10)"`
}

func (Student) TableName() string {
	return "Student"
}

// CREATE TABLE Course (
//
//	course_id INT AUTO_INCREMENT,
//	course_name VARCHAR(255),
//	PRIMARY KEY (course_id)
//
// );
type Course struct {
	ID   uint32 `gorm:"column:course_id;type:int(32);primary_key;auto_increment"`
	Name string `gorm:"column:course_name;type:varchar(255)"`
}

func (Course) TableName() string {
	return "Course"
}

// -- Class and Room are joined into 1 table Class
// CREATE TABLE Class (
//
//	class_id INT AUTO_INCREMENT,
//	class_name VARCHAR(255),
//	prof_id INT,
//	course_id INT,
//	room_loc VARCHAR(50),
//	room_cap VARCHAR(50),
//	PRIMARY KEY (class_id),
//	FOREIGN KEY (prof_id) REFERENCES Professor(prof_id),
//	FOREIGN KEY (course_id) REFERENCES Course(course_id)
//
// );
type Class struct {
	ID          uint32 `gorm:"column:class_id;type:int(32);primary_key;auto_increment"`
	Name        string `gorm:"column:class_name;type:varchar(255)"`
	ProfessorID uint32 `gorm:"column:prof_id;type:int(32)"`
	CourseID    uint32 `gorm:"column:course_id;type:int(32)"`
	RoomLoc     string `gorm:"column:room_loc;type:varchar(50)"`
	RoomCap     string `gorm:"column:room_cap;type:varchar(50)"`

	Professor Professor `gorm:"foreign_key:prof_id;references:prof_id"`
	Course    Course    `gorm:"foreign_key:course_id;references:course_id"`
}

func (Class) TableName() string {
	return "Class"
}

// CREATE TABLE Enroll (
//
//	stud_id INT,
//	class_id INT,
//	grade VARCHAR(3),
//	PRIMARY KEY (stud_id, class_id),
//	FOREIGN KEY (stud_id) REFERENCES Student(stud_id),
//	FOREIGN KEY (class_id) REFERENCES Class(class_id)
//
// );
type Enroll struct {
	StudentID int    `gorm:"column:stud_id;type:int(32);primary_key"`
	ClassID   int    `gorm:"column:class_id;type:int(32);primary_key"`
	Grade     string `gorm:"column:grade;type:varchar(3)"`

	Student Student `gorm:"foreign_key:stud_id;references:stud_id"`
	Class   Class   `gorm:"foreign_key:class_id;references:class_id"`
}

func (Enroll) TableName() string {
	return "Enroll"
}
