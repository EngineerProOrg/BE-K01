CREATE TABLE professors (
    prof_id int NOT NULL AUTO_INCREMENT,
    prof_lname VARCHAR(50),
    prof_fname VARCHAR(50),
    PRIMARY KEY (prof_id)
);

CREATE TABLE students (
    stud_id int NOT NULL AUTO_INCREMENT,
    stud_lname VARCHAR(50),
    stud_fname VARCHAR(50),
    stud_street VARCHAR(255),
    stud_city VARCHAR(50),
    stud_zip VARCHAR(10),
    PRIMARY KEY (stud_id)
);


CREATE TABLE rooms (
    room_id int NOT NULL AUTO_INCREMENT,,
    room_loc VARCHAR(50),
    room_cap VARCHAR(50),
    PRIMARY KEY (room_id)
);

CREATE TABLE courses (
    course_id int NOT NULL AUTO_INCREMENT,,
    course_name VARCHAR(50),
    PRIMARY KEY (course_id)
);

CREATE TABLE classes (
    class_id int NOT NULL AUTO_INCREMENT,,
    class_name VARCHAR(255),
    prof_id int,
    course_id int,
    room_id int,
    PRIMARY KEY (class_id),
    UNIQUE (room_id),
    CONSTRAINT fk_professor_class FOREIGN KEY (prof_id) REFERENCES professors(prof_id),
    CONSTRAINT fk_room_class FOREIGN KEY (room_id) REFERENCES rooms(room_id),
    CONSTRAINT fk_courses_class FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

CREATE TABLE enrolls (
    stud_id int,
    class_id int,
    grade VARCHAR(3),
    PRIMARY KEY (stud_id, class_id),
    CONSTRAINT fk_student_enroll FOREIGN KEY (stud_id) REFERENCES students(stud_id),
    CONSTRAINT fk_class_enroll FOREIGN KEY (class_id) REFERENCES classes(class_id)
);