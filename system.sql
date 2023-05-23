create table if not exists Professor
(
    prof_id    int primary key,
    prof_lname varchar(50),
    prof_fname varchar(50)
    );

create table if not exists Course
(
    course_id   int primary key,
    course_name varchar(255)
    );


create table if not exists Room
(
    room_id  int primary key,
    room_loc varchar(50),
    room_cap varchar(50),
    class_id int
    );

create table if not exists Class
(
    class_id   int primary key,
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
    stud_id     int primary key,
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

INSERT INTO Professor (prof_id, prof_lname, prof_fname)
VALUES (1, 'Smith', 'John'),
       (2, 'Johnson', 'Mary'),
       (3, 'Davis', 'David');
INSERT INTO Course (course_id, course_name)
VALUES (1, 'Mathematics'),
       (2, 'English'),
       (3, 'Science');


INSERT INTO Room (room_id, room_loc, room_cap)
VALUES (1, 'Building A, Room 101', '30'),
       (2, 'Building B, Room 201', '25'),
       (3, 'Building C, Room 301', '35');
INSERT INTO Class (class_id, class_name, prof_id, course_id, room_id)
VALUES (1, 'Class 1', 1, 1, 1),
       (2, 'Class 2', 2, 2, 2),
       (3, 'Class 3', 3, 3, 3);

update Room
set class_id=1
where room_id = 1;

update Room
set class_id=2
where room_id = 2;

update Room
set class_id=3
where room_id = 3;


INSERT INTO Student (stud_id, stud_fname, stud_lname, stud_street, stud_city, stud_zip)
VALUES (1, 'Alice', 'Johnson', '123 Main St', 'New York', '10001'),
       (2, 'Bob', 'Smith', '456 Elm St', 'Los Angeles', '90001'),
       (3, 'Charlie', 'Davis', '789 Oak St', 'Chicago', '60601');
INSERT INTO Enroll (stud_id, class_id, grade)
VALUES (1, 1, 'A'),
       (2, 2, 'B'),
       (3, 3, 'C');

#1
select CONCAT(p.prof_fname, ' ', p.prof_lname) as 'Professor Name',
        CONCAT(s.stud_fname, ' ', s.stud_lname) as 'Student Name',
        c.class_name                            as 'Class Name'
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
SELECT distinct(CONCAT(s.stud_fname, ' ', s.stud_lname)) AS 'Student Name',
        AvgGrade.Grade                            AS 'Average Grade',
        CASE
            WHEN AvgGrade.Grade < 5 THEN 'Weak'
            WHEN AvgGrade.Grade >= 5 AND AvgGrade.Grade < 8 THEN 'Average'
            WHEN AvgGrade.Grade >= 8 THEN 'Good'
            END                                   AS 'Grade'
FROM Student s
         JOIN Enroll e ON s.stud_id = e.stud_id
         JOIN (SELECT e.stud_id,
                      AVG(
                              CASE e.grade
                                  WHEN 'A' THEN 10
                                  WHEN 'B' THEN 8
                                  WHEN 'C' THEN 6
                                  WHEN 'D' THEN 4
                                  WHEN 'E' THEN 2
                                  WHEN 'F' THEN 0
                                  ELSE 0
                                  END
                          ) AS 'Grade'
               FROM Enroll e
               GROUP BY stud_id) AS AvgGrade ON s.stud_id = AvgGrade.stud_id;

#5
select c.class_name,
       AvgGrade.Grade as 'Grade',
        CASE
            WHEN AvgGrade.Grade < 5 THEN 'Weak'
            WHEN AvgGrade.Grade >= 5 AND AvgGrade.Grade < 8 THEN 'Average'
            WHEN AvgGrade.Grade >= 8 THEN 'Good'
            END AS 'Grade'
from Class c
         JOIN (SELECT e.class_id,
                      AVG(
                              CASE e.grade
                                  WHEN 'A' THEN 10
                                  WHEN 'B' THEN 8
                                  WHEN 'C' THEN 6
                                  WHEN 'D' THEN 4
                                  WHEN 'E' THEN 2
                                  WHEN 'F' THEN 0
                                  ELSE 0
                                  END
                          ) AS 'Grade'
               FROM Enroll e
               GROUP BY e.class_id) AS AvgGrade ON c.class_id = AvgGrade.class_id
#6
select co.course_name,
       AvgGrade.Grade as 'Grade',
        CASE
            WHEN AvgGrade.Grade < 5 THEN 'Weak'
            WHEN AvgGrade.Grade >= 5 AND AvgGrade.Grade < 8 THEN 'Average'
            WHEN AvgGrade.Grade >= 8 THEN 'Good'
            END AS 'Grade'
from Course co
         JOIN (SELECT coo.course_id,
                      AVG(
                              CASE e.grade
                                  WHEN 'A' THEN 10
                                  WHEN 'B' THEN 8
                                  WHEN 'C' THEN 6
                                  WHEN 'D' THEN 4
                                  WHEN 'E' THEN 2
                                  WHEN 'F' THEN 0
                                  ELSE 0
                                  END
                          ) AS 'Grade'
               FROM Enroll e join Course coo on e.class_id = coo.course_id
               GROUP BY coo.course_id) AS AvgGrade ON co.course_id = AvgGrade.course_id