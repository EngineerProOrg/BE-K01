create table if not exists Professor
(
    prof_id    int primary key auto_increment,
    prof_lname varchar(50),
    prof_fname varchar(50)
    );

create table if not exists Course
(
    course_id   int primary key auto_increment,
    course_name varchar(255)
    );


create table if not exists Room
(
    room_id  int primary key auto_increment,
    room_loc varchar(50),
    room_cap varchar(50),
    class_id int
    );

create table if not exists Class
(
    class_id   int primary key auto_increment,
    class_name varchar(255),
    prof_id    int,
    course_id  int,
    room_id    int,
    FOREIGN KEY (prof_id) REFERENCES Professor (prof_id),
    FOREIGN KEY (course_id) REFERENCES Course (course_id),
    FOREIGN KEY (room_id) REFERENCES Room (room_id)
    );

alter table Room
    add FOREIGN KEY (class_id) REFERENCES Class (class_id);


create table if not exists Student
(
    stud_id     int primary key auto_increment,
    stud_fname  varchar(50),
    stud_lname  varchar(50),
    stud_street varchar(255),
    stud_city   varchar(50),
    stud_zip    varchar(10)
    );

create table if not exists Enroll
(
    stud_id  int,
    class_id int,
    grade    varchar(3),
    primary key (stud_id, class_id),
    FOREIGN KEY (stud_id) REFERENCES Student (stud_id),
    FOREIGN KEY (class_id) REFERENCES Class (class_id)
    );

INSERT INTO Professor (prof_lname, prof_fname)
VALUES ('Smith', 'John'),
       ('Johnson', 'Mary'),
       ('Davis', 'David');
INSERT INTO Course (course_name)
VALUES ('Mathematics'),
       ('English'),
       ('Science');


INSERT INTO Room (room_loc, room_cap)
VALUES ('Building A, Room 101', '30'),
       ('Building B, Room 201', '25'),
       ('Building C, Room 301', '35');
INSERT INTO Class (class_name, prof_id, course_id, room_id)
VALUES ('Class 1', 1, 1, 1),
       ('Class 2', 2, 2, 2),
       ('Class 3', 3, 3, 3);

update Room
set class_id=1
where room_id = 1;

update Room
set class_id=2
where room_id = 2;

update Room
set class_id=3
where room_id = 3;


INSERT INTO Student (stud_fname, stud_lname, stud_street, stud_city, stud_zip)
VALUES ('Alice', 'Johnson', '123 Main St', 'New York', '10001'),
       ('Bob', 'Smith', '456 Elm St', 'Los Angeles', '90001'),
       ('Charlie', 'Davis', '789 Oak St', 'Chicago', '60601');
INSERT INTO Enroll (class_id, grade)
VALUES (1, 'A'),
       (2, 'B'),
       (3, 'C');

Create OR REPLACE VIEW Grade AS
SELECT e.class_id,e.stud_id,
       CASE e.grade
           WHEN 'A' THEN 10
           WHEN 'B' THEN 8
           WHEN 'C' THEN 6
           WHEN 'D' THEN 4
           WHEN 'E' THEN 2
           WHEN 'F' THEN 0
           ELSE 0
           END
    AS 'Grade'
FROM Enroll e


#1
select CONCAT(p.prof_fname, ' ', p.prof_lname) as 'Professor Name', CONCAT(s.stud_fname, ' ', s.stud_lname) as 'Student Name', c.class_name as 'Class Name'
from Professor p
         join Class c on p.prof_id = c.prof_id
         join Enroll E on c.class_id = E.class_id
         join Student s on E.stud_id = s.stud_id
#2
select distinct(co.course_name)
from Course co
         join Class c on co.course_id = c.course_id
#3
select distinct(co.course_name)
from Course co
         join Class c on co.course_id = c.course_id
         join Enroll E on c.class_id = E.class_id

#4
select CONCAT(s.stud_fname, ' ', s.stud_lname) AS 'Student Name',
        AVG(G.Grade)                            AS 'Average Grade',
        CASE
            WHEN AVG(G.Grade) < 5 THEN 'Weak'
            WHEN AVG(G.Grade) >= 5 AND AVG(G.Grade) < 8 THEN 'Average'
            WHEN AVG(G.Grade) >= 8 THEN 'Good'
            END                                 AS 'Grade'
from Student s
         join Grade G on s.stud_id = G.stud_id
group by s.stud_id


#5
select c.class_name AS 'Class Name',
        AVG(G.Grade) AS 'Average Grade',
        CASE
            WHEN AVG(G.Grade) < 5 THEN 'Weak'
            WHEN AVG(G.Grade) >= 5 AND AVG(G.Grade) < 8 THEN 'Average'
            WHEN AVG(G.Grade) >= 8 THEN 'Good'
            END      AS 'Grade'
from Class c
         join Grade G on c.class_id = G.class_id
group by c.class_id, c.class_name


#6
select c.course_name as 'Course Name',
        AVG(G.Grade) AS 'Average Grade',
        CASE
            WHEN AVG(G.Grade) < 5 THEN 'Weak'
            WHEN AVG(G.Grade) >= 5 AND AVG(G.Grade) < 8 THEN 'Average'
            WHEN AVG(G.Grade) >= 8 THEN 'Good'
            END      AS 'Grade'
from Course c join Class C2 on c.course_id = C2.course_id
              join Grade G on C2.class_id = G.class_id
group by c.course_id, c.course_name

