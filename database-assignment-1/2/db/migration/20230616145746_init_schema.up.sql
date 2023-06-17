CREATE TABLE IF NOT EXISTS professor(
	prof_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	prof_lname VARCHAR(50),
	prof_fname VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS course(
	course_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	course_name VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS class(
	class_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	class_name VARCHAR(255),
	prof_id INT,
	course_id INT,
	room_id INT
);


CREATE TABLE IF NOT EXISTS room(
	room_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	room_loc VARCHAR(50),
	room_cap VARCHAR(50),
	class_id INT
);


CREATE TABLE IF NOT EXISTS student(
	stud_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	stud_fname VARCHAR(50),
	stud_lname VARCHAR(50),
	stud_street VARCHAR(255),
	stud_city VARCHAR(50),
	stud_zip VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS enroll(
	stud_id INT,
	class_id INT,
	grade VARCHAR(3),
	
	CONSTRAINT pk_enroll PRIMARY KEY (stud_id, class_id)
);

-- ADD FOREIGN KEY for talbe class
ALTER TABLE class
ADD CONSTRAINT fk_class_professor
FOREIGN KEY (prof_id) REFERENCES professor(prof_id);

ALTER TABLE class
ADD CONSTRAINT fk_class_course
FOREIGN KEY (course_id) REFERENCES course(course_id);

ALTER TABLE class
ADD CONSTRAINT fk_class_room
FOREIGN KEY (room_id) REFERENCES room(room_id);

-- ADD FOREIGN KEY for talbe room
ALTER TABLE room
ADD CONSTRAINT fk_room_class
FOREIGN KEY (class_id) REFERENCES class(class_id);

-- ADD FOREIGN KEY for table enroll
ALTER TABLE enroll
ADD CONSTRAINT fk_enroll_student
FOREIGN KEY (stud_id) REFERENCES student(stud_id);


ALTER TABLE enroll
ADD CONSTRAINT fk_enroll_class
FOREIGN KEY (class_id) REFERENCES class(class_id);
