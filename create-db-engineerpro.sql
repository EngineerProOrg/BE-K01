USE engineerpro;
CREATE TABLE Professor (
	prof_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    prof_lname VARCHAR(50),
    prof_fname VARCHAR(50)
);
CREATE TABLE Student (
	stud_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    stud_fname VARCHAR(50) NOT NULL, 
    stud_lname VARCHAR(50) NOT NULL,
    stud_street VARCHAR(255), 
    stud_city VARCHAR(50), 
    stud_zip VARCHAR(10)
);
CREATE TABLE Course(
	course_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    course_name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS Room(
	room_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    room_loc VARCHAR(50) NOT NULL,
    room_cap VARCHAR(50) NOT NULL,
    class_id INT NULL
);
CREATE TABLE IF NOT EXISTS Class (
	class_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    class_name VARCHAR(255) NOT NULL,
    prof_id INT NOT NULL, 
    course_id INT NOT NULL,
    room_id INT NOT NULL-- add constraint later
);
CREATE TABLE Enroll(
	stud_id INT NOT NULL,
    class_id INT NOT NULL,
    grade VARCHAR(3) NOT NULL,
    PRIMARY KEY (stud_id, class_id)
);
-- Add constraint to Class table
ALTER TABLE Class ADD CONSTRAINT fk_class_professor_prof_id
FOREIGN KEY (prof_id) REFERENCES Professor (prof_id);

ALTER TABLE Class ADD CONSTRAINT fk_class_course_course_id
FOREIGN KEY (course_id) REFERENCES Course (course_id);

ALTER TABLE Class ADD CONSTRAINT fk_class_room_room_id
FOREIGN KEY (room_id) REFERENCES Room (room_id);
-- end


-- add constraint to Enroll table
ALTER TABLE Enroll ADD CONSTRAINT fk_enroll_student_stud_id
FOREIGN KEY (stud_id) REFERENCES Student (stud_id);

ALTER TABLE Enroll ADD CONSTRAINT fk_enroll_class_class_id
FOREIGN KEY (class_id) REFERENCES Class (class_id);
-- end

-- add constraint to Room table
ALTER TABLE Room ADD CONSTRAINT fk_room_class_class_id
FOREIGN KEY (class_id) REFERENCES Class (class_id);
-- end

-- INSERT 
INSERT INTO Professor (prof_lname, prof_fname)
VALUES
	('Smith', 'John'),
	('Johnson', 'Emily'),
	('Brown', 'Michael'),
	('Davis', 'Sarah'),
	('Miller', 'David'),
	('Wilson', 'Jennifer'),
	('Anderson', 'Christopher'),
	('Thomas', 'Jessica'),
	('Taylor', 'Matthew'),
	('Moore', 'Amanda');

INSERT INTO Student (stud_fname, stud_lname, stud_street, stud_city, stud_zip) 
VALUES
	('Emma', 'Johnson', '123 Main St', 'New York', '10001'),
	('Oliver', 'Williams', '456 Elm St', 'Los Angeles', '90001'),
	('Sophia', 'Jones', '789 Oak St', 'Chicago', '60601'),
	('Liam', 'Brown', '234 Maple St', 'Houston', '77001'),
	('Ava', 'Davis', '567 Pine St', 'Philadelphia', '19101'),
	('Noah', 'Wilson', '890 Cedar St', 'Phoenix', '85001'),
	('Isabella', 'Lee', '321 Birch St', 'San Antonio', '78201'),
	('Mason', 'Taylor', '654 Willow St', 'San Diego', '92101'),
	('Charlotte', 'Clark', '987 Spruce St', 'Dallas', '75201'),
	('Elijah', 'Moore', '654 Rose St', 'Austin', '78701');

INSERT INTO Course (course_name) 
VALUES
	('Mathematics'),
	('English'),
	('History'),
	('Science'),
	('Art'),
	('Computer Science'),
	('Physics'),
	('Biology'),
	('Chemistry'),
	('Music');
    
-- set class id null than update later on
INSERT INTO Room (room_loc, room_cap, class_id) 
VALUE
	('Building A, Room 101', '30',NULL),
	('Building A, Room 102', '25',NULL),
	('Building B, Room 201', '40',NULL),
	('Building B, Room 202', '35',NULL),
	('Building C, Room 301', '20',NULL),
	('Building C, Room 302', '15',NULL),
	('Building D, Room 401', '30',NULL),
	('Building D, Room 402', '25',NULL),
	('Building E, Room 501', '40',NULL),
	('Building E, Room 502', '35',NULL);
    
INSERT INTO Class (class_name, prof_id, course_id, room_id) 
VALUES
	('Class 1', 1, 1, 1),
	('Class 2', 2, 2, 2),
	('Class 3', 3, 3, 3),
	('Class 4', 4, 4, 4),
	('Class 5', 5, 5, 5),
	('Class 6', 6, 6, 6),
	('Class 7', 7, 7, 7),
	('Class 8', 8, 8, 8),
	('Class 9', 9, 9, 9),
	('Class 10',10,10,10);

INSERT INTO Enroll (stud_id, class_id, grade) 
VALUES
	(1, 1, 'A'),
	(2, 2, 'B'),
	(3, 3, 'C'),
	(4, 4, 'A'),
	(5, 5, 'B'),
	(6, 6, 'C'),
	(7, 7, 'A'),
	(8, 8, 'B'),
	(9, 9, 'C'),
	(10, 10, 'A');


