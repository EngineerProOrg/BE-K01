create table if not exists professors
(
    prof_id    int primary key AUTO_INCREMENT,
    prof_lname varchar(50),
    prof_fname varchar(50)
    );

create table if not exists courses
(
    course_id   int primary key AUTO_INCREMENT,
    course_name varchar(255)
    );


create table if not exists rooms
(
    room_id  int primary key AUTO_INCREMENT,
    room_loc varchar(50),
    room_cap varchar(50),
    class_id int
    );

create table if not exists classes
(
    class_id   int primary key AUTO_INCREMENT,
    class_name varchar(255),
    prof_id    int,
    course_id  int,
    room_id    int,
    FOREIGN KEY (prof_id) REFERENCES professors (prof_id),
    FOREIGN KEY (course_id) REFERENCES courses (course_id),
    FOREIGN KEY (room_id) REFERENCES rooms (room_id)
    );

alter table rooms
    add FOREIGN KEY (class_id) REFERENCES classes (class_id);


create table if not exists students
(
    stud_id     int primary key AUTO_INCREMENT,
    stud_fname  varchar(50),
    stud_lname  varchar(50),
    stud_street varchar(255),
    stud_city   varchar(50),
    stud_zip    varchar(10)
    );

create table if not exists enrolls
(
    stud_id  int,
    class_id int,
    grade    varchar(3),
    primary key (stud_id, class_id),
    FOREIGN KEY (stud_id) REFERENCES students (stud_id),
    FOREIGN KEY (class_id) REFERENCES classes (class_id)
    );
