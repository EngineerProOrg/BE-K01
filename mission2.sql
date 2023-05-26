
-- create tables  
CREATE TABLE IF NOT EXISTS professor (
  id INT PRIMARY KEY,
  firstName VARCHAR(50) NOT NULL,
  lastName VARCHAR(50) NOT NULL
  
);

CREATE TABLE IF NOT EXISTS course (
  id INT PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);


CREATE TABLE IF NOT EXISTS student (
  id INT PRIMARY KEY,
  firstName VARCHAR(50) NOT NULL,
  lastName VARCHAR(50) NOT NULL,
  street varchar(255),
  city varchar(50),
  zip varchar(10)
);

CREATE TABLE IF NOT EXISTS class (
  id INT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  professor_id INT NOT NULL,
  course_id INT NOT NULL,
  room_id INT NOT NULL,
  FOREIGN KEY (professor_id) REFERENCES professor(id),
  FOREIGN KEY (course_id) REFERENCES course(id)
  -- FOREIGN KEY (room_id) REFERENCES room(id)  
);


CREATE TABLE IF NOT EXISTS room (
  id INT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  loc varchar(50),
  cap varchar(50),
  class_id INT UNIQUE NOT NULL,
  FOREIGN KEY (class_id) REFERENCES class(id)
);

-- vì đây là quan hệ n-n nên không khai báo fk trong cả 2 đc vì chạy k đc, nên bỏ 1 trong 2 tạo bảng rồi mới add fk  
alter table class 
add constraint class_room
foreign key(room_id) references room(id);


CREATE TABLE IF NOT EXISTS class_student (
  class_id INT NOT NULL,
  student_id INT NOT NULL,
  
  grade varchar(3),
  PRIMARY KEY (class_id, student_id),
  FOREIGN KEY (class_id) REFERENCES class(id),
  FOREIGN KEY (student_id) REFERENCES student(id)
);

-- những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan 
-- tạo thêm table class_professor để tối ưu query 

CREATE TABLE IF NOT EXISTS class_professor (
  class_id INT NOT NULL,
  professor_id INT NOT NULL,
  PRIMARY KEY (class_id, professor_id),
  FOREIGN KEY (class_id) REFERENCES class(id),
  FOREIGN KEY (professor_id) REFERENCES professor(id)
);


select t.id as student_id, t.lastName as student_name,
	   p.id as professor_id, p.lastName as professor_name,
       count(distinct(class_id)) as numClass
from student as t
join class_student as c_s
on t.id= c_s.student_id
join class_professor as c_f
on c_s.class_id= c_f.class_id
join professor as p 
on c_f.professor_id= p.id
group by student_id,professor_id;


-- những course (distinct) mà 1 professor cụ thể đang dạy

select p.id, c.course_id

from professor as p
join class as c on p.id= c.professor_id
group by p.id, c.course_id;


-- những course (distinct) mà 1 student cụ thể đang học 
SELECT s.id AS student_id, c.name AS course_name
FROM class c
JOIN class_student cs ON c.id = cs.class_id
JOIN student s ON cs.student_id = s.id
GROUP BY student_id, course_name;


-- điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0 
-- điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )

SELECT cs.student_id,
	   AVG(CASE
             WHEN cs.grade =10 THEN 'A'
             WHEN cs.grade >= 8 THEN 'B'
             WHEN cs.grade >= 6 THEN 'C'
             WHEN cs.grade >= 4 THEN 'D'
             WHEN cs.grade >= 2 THEN 'E'
             ELSE 'F'
           END) AS letter_grade,
       CASE
         WHEN AVG(cs.grade) < 5 THEN 'weak'
         WHEN AVG(cs.grade) < 8 THEN 'average'
         ELSE 'good'
       END AS performance_level
FROM class_student cs
group by cs.student_id;

--  điểm số trung bình của các class (quy ra lại theo chữ cái) 

SELECT c.id AS class_id, c.name AS class_name,
       AVG(CASE
			 WHEN cs.grade =10 THEN 'A'
             WHEN cs.grade >= 8 THEN 'B'
             WHEN cs.grade >= 6 THEN 'C'
             WHEN cs.grade >= 4 THEN 'D'
             WHEN cs.grade >= 2 THEN 'E'
             ELSE 'F'
           END) AS letter_grade
FROM class c
JOIN class_student cs ON c.id = cs.class_id
GROUP BY class_id, class_name;

-- điểm số trung bình của các course (quy ra lại theo chữ cái)

SELECT c.course_id AS course_id,
       AVG(CASE
			 WHEN cs.grade =10 THEN 'A'
             WHEN cs.grade >= 8 THEN 'B'
             WHEN cs.grade >= 6 THEN 'C'
             WHEN cs.grade >= 4 THEN 'D'
             WHEN cs.grade >= 2 THEN 'E'
             ELSE 'F'
           END) AS letter_grade
FROM class c
JOIN class_student cs ON c.id = cs.class_id
GROUP BY course_id
 












