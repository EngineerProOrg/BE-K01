-- Create tables
drop table if exists Enroll;

drop table if exists Class;

drop table if exists Professor;

drop table if exists Student;

drop table if exists Course;

create table Professor (
    prof_id int auto_increment,
    prof_fname varchar(50),
    prof_lname varchar(50),
    primary key (prof_id)
);

create table Student (
    stud_id int auto_increment,
    stud_fname varchar(50),
    stud_lname varchar(50),
    stud_street varchar(255),
    stud_city varchar(50),
    stud_zip varchar(10),
    primary key (stud_id)
);

create table Course (
    course_id int auto_increment,
    course_name varchar(255),
    primary key (course_id)
);

-- Class and Room are joined into 1 table Class
create table Class (
    class_id int auto_increment,
    room_id int not NULL unique,
    class_name varchar(255),
    prof_id int,
    course_id int,
    room_loc varchar(50),
    room_cap varchar(50),
    primary key (class_id),
    foreign key (prof_id) references Professor(prof_id),
    foreign key (course_id) references Course(course_id)
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
insert into
    Professor
values
    (NULL, 'Albus', 'Dumbledore');

insert into
    Professor
values
    (NULL, 'Severus', 'Snape');

insert into
    Professor
values
    (NULL, 'Alastor', 'Moody');

-- Insert to Student
insert into
    Student
values
    (
        NULL,
        "Harry",
        "Potter",
        "Privet Drive 1",
        "London",
        "100000"
    );

insert into
    Student
values
    (
        NULL,
        "Hermione",
        "Granger",
        "Privet Drive 2",
        "London",
        "200000"
    );

insert into
    Student
values
    (
        NULL,
        "Ron",
        "Weasly",
        "Privet Drive 3",
        "London",
        "300000"
    );

-- Insert to Course
insert into
    Course
values
    (NULL, "Physical Education");

insert into
    Course
values
    (NULL, "Quidditch");

insert into
    Course
values
    (NULL, "Defence Against the Dark Arts");

insert into
    Course
values
    (NULL, "Potion");

-- Insert to Class
insert into
    Class
values
    (NULL, 1, "PE1", 1, 1, "001", "100");

insert into
    Class
values
    (NULL, 2, "QD1", 2, 2, "002", "200");

insert into
    Class
values
    (NULL, 3, "DA1", 3, 3, "003", "300");

insert into
    Class
values
    (NULL, 4, "DA2", 3, 3, "004", "400");

insert into
    Class
values
    (NULL, 5, "PT1", 3, 4, "005", "500");

-- Insert to Enroll
insert into
    Enroll
values
    (1, 2, "B");

insert into
    Enroll
values
    (2, 1, "A");

insert into
    Enroll
values
    (2, 2, "A");

insert into
    Enroll
values
    (2, 3, "A");

insert into
    Enroll
values
    (3, 1, "C");

insert into
    Enroll
values
    (3, 3, "D");

insert into
    Enroll
values
    (1, 4, "D");

insert into
    Enroll
values
    (2, 4, "D");

insert into
    Enroll
values
    (3, 4, "D");