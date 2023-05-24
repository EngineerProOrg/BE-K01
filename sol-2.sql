-- Create tables
drop table if exists Enroll;

drop table if exists Class;

drop table if exists Professor;

drop table if exists Student;

drop table if exists Course;

drop table if exists Room;

create table Professor (
    prof_id int,
    prof_fname varchar(50),
    prof_lname varchar(50),
    primary key (prof_id)
);

create table Student (
    stud_id int,
    stud_fname varchar(50),
    stud_lname varchar(50),
    stud_street varchar(255),
    stud_city varchar(50),
    stud_zip varchar(10),
    primary key (stud_id)
);

create table Course (
    course_id int,
    course_name varchar(255),
    primary key (course_id)
);

create table Room (
    room_id int,
    room_loc varchar(50),
    room_cap varchar(50),
    primary key (room_id)
);

create table Class (
    class_id int,
    class_name varchar(255),
    prof_id int,
    course_id int,
    room_id int unique,
    primary key (class_id),
    foreign key (prof_id) references Professor(prof_id),
    foreign key (course_id) references Course(course_id),
    foreign key (room_id) references Room(room_id)
);

create table Enroll (
    stud_id int,
    class_id int,
    grade varchar(3),
    primary key (stud_id, class_id),
    foreign key (stud_id) references Student(stud_id),
    foreign key (class_id) references Class(class_id)
);

-- Add records

-- Insert to Professor
insert into Professor values (1, 'Albus', 'Dumbledore');
insert into Professor values (2, 'Severus', 'Snape');
insert into Professor values (3, 'Alastor', 'Moody');

-- Insert to Student
insert into Student values (1, "Harry", "Potter", "Privet Drive 1", "London", "100000");
insert into Student values (2, "Hermione", "Granger", "Privet Drive 2", "London", "200000");
insert into Student values (3, "Ron", "Weasly", "Privet Drive 3", "London", "300000");

-- Insert to Course
insert into Course values (1, "Physical Education");
insert into Course values (2, "Quidditch");
insert into Course values (3, "Defence Against the Dark Arts");
insert into Course values (4, "Potion");

-- Insert to Room
insert into Room values (1, "001", "100");
insert into Room values (2, "002", "200");
insert into Room values (3, "003", "300");
insert into Room values (4, "004", "400");
insert into Room values (5, "005", "500");

-- Insert to Class
insert into Class values (1, "PE1", 1, 1, 1);
insert into Class values (2, "QD1", 2, 2, 2);
insert into Class values (3, "DA1", 3, 3, 3);
insert into Class values (4, "DA2", 3, 3, 4);
insert into Class values (5, "PT1", 3, 4, 5);

-- Insert to Enroll
insert into Enroll values (1, 2, "B");
insert into Enroll values (2, 1, "A");
insert into Enroll values (2, 2, "A");
insert into Enroll values (2, 3, "A");
insert into Enroll values (3, 1, "C");
insert into Enroll values (3, 3, "D");
insert into Enroll values (1, 4, "D");
insert into Enroll values (2, 4, "D");
insert into Enroll values (3, 4, "D");
