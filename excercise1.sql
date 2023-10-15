/*Database

Mỗi bài tập sẽ lưu trên file sol-{number}.sql

1. Giải bài tập trên leetcode: 

https://leetcode.com/problems/capital-gainloss/description/ (gợi ý: sử dụng CASE)*/
--------------------------------------------------------------------------------------
select new.stock_name, sum(newprice) as capital_gain_loss from (
    select old.stock_name, 
    CASE
    WHEN operation = 'Buy' THEN -price
    ELSE price
    END AS newprice
    from Stocks as old
) as new group by stock_name;
--------------------------------------------------------------------------------------
/*https://leetcode.com/problems/count-salary-categories/ (ngoài các cách trên leetcode, hãy nghĩ cách để giúp câu query này nhanh hơn, kể cả thay đổi cấu trúc bản)*/
--------------------------------------------------------------------------------------
select 
new.category, count(new.category)-1 as accounts_count
from (
    select
    CASE
    WHEN old.income > 50000 THEN 'High Salary'
    WHEN old.income < 20000 THEN 'Low Salary'
    ELSE 'Average Salary'
    END AS category
    from Accounts as old
    UNION ALL
    SELECT 'High Salary' as category
    UNION ALL
    SELECT 'Low Salary' as category
    UNION ALL
    SELECT 'Average Salary' as category
) as new group by new.category;
--------------------------------------------------------------------------------------
/*2. Bạn hãy viết một script để tạo các bản cho hệ thống với cấu trúc ở dưới

![img.png](img.png)

hệ thống bao gồm:

- class
- professor: quan hệ one-many với class
- student: quan hệ many-many với class
- course: quan hệ one-many với class
- room: quan hệ one-one với class*/
--------------------------------------------------------------------------------------
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';
--
DROP DATABASE IF EXISTS `SCHOOLDB`;
CREATE DATABASE IF NOT EXISTS `SCHOOLDB`;
USE `SCHOOLDB`;
DROP SCHEMA IF EXISTS `COLLEGE`;
CREATE SCHEMA `COLLEGE` DEFAULT CHARACTER SET utf8;
USE `COLLEGE`;
--
DROP TABLE IF EXISTS `COLLEGE`.`COURSE`;

CREATE TABLE IF NOT EXISTS `COLLEGE`.`COURSE` (
 `COURSE_ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
 `COURSE_NAME` VARCHAR(255) NOT NULL,
 PRIMARY KEY (`COURSE_ID`));
--
DROP TABLE IF EXISTS `COLLEGE`.`PROFESSOR`;

CREATE TABLE IF NOT EXISTS `COLLEGE`.`PROFESSOR` (
 `PROF_ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
 `PROF_LNAME` VARCHAR(50) NOT NULL,
 `PROF_FNAME` VARCHAR(50) NOT NULL,
 PRIMARY KEY (`PROF_ID`));
--
DROP TABLE IF EXISTS `COLLEGE`.`STUDENT`;

CREATE TABLE IF NOT EXISTS `COLLEGE`.`STUDENT` (
 `STUD_ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
 `STUD_FNAME` VARCHAR(50) NOT NULL,
 `STUD_LNAME` VARCHAR(50) NOT NULL,
 `STUD_STREET` VARCHAR(255) NOT NULL,
 `STUD_CITY` VARCHAR(50) NOT NULL,
 `STUD_ZIP` VARCHAR(10) NOT NULL,
 PRIMARY KEY (`STUD_ID`));
--
DROP TABLE IF EXISTS `COLLEGE`.`CLASS`;

CREATE TABLE IF NOT EXISTS `COLLEGE`.`CLASS` (
 `CLASS_ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
 `CLASS_NAME` VARCHAR(255) NOT NULL,
 `PROF_ID` INT NOT NULL,
 `COURSE_ID` INT NOT NULL,
 `ROOM_ID` INT NOT NULL,
 PRIMARY KEY (`CLASS_ID`),
 INDEX `PROF_ID_IDX` (`PROF_ID` ASC),
 INDEX `COURSE_ID_IDX` (`COURSE_ID` ASC),
 INDEX `ROOM_ID_IDX` (`ROOM_ID` ASC),
 CONSTRAINT `PROF_CLASS_FK` FOREIGN KEY (`PROF_ID`) REFERENCES `COLLEGE`.`PROFESSOR` (`PROF_ID`)
 ON DELETE RESTRICT
 ON UPDATE RESTRICT,
 CONSTRAINT `COURSE_CLASS_FK` FOREIGN KEY (`COURSE_ID`) REFERENCES `COLLEGE`.`COURSE` (`COURSE_ID`)
 ON DELETE RESTRICT
 ON UPDATE RESTRICT,
 CONSTRAINT `ROOM_CLASS_FK` FOREIGN KEY (`ROOM_ID`) REFERENCES `COLLEGE`.`ROOM` (`ROOM_ID`)
 ON DELETE RESTRICT
 ON UPDATE RESTRICT);
--
 DROP TABLE IF EXISTS `COLLEGE`.`ENROLL`;

CREATE TABLE IF NOT EXISTS `COLLEGE`.`ENROLL` (
 `STUD_ID` INT NOT NULL,
 `CLASS_ID` INT NOT NULL,
 `GRADE` VARCHAR(3) NOT NULL,
 CONSTRAINT ENROLL_PK PRIMARY KEY (`STUD_ID`, `CLASS_ID`),
 CONSTRAINT `STUD_ENROLL_FK` FOREIGN KEY (`STUD_ID`) REFERENCES `COLLEGE`.`STUDENT` (`STUD_ID`)
 ON DELETE RESTRICT
 ON UPDATE RESTRICT,
 CONSTRAINT `CLASS_ENROLL_FK` FOREIGN KEY (`CLASS_ID`) REFERENCES `COLLEGE`.`CLASS` (`CLASS_ID`)
 ON DELETE RESTRICT
 ON UPDATE RESTRICT);
 --
DROP TABLE IF EXISTS `COLLEGE`.`ROOM`;

CREATE TABLE IF NOT EXISTS `COLLEGE`.`ROOM` (
 `ROOM_ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
 `ROOM_LOC` VARCHAR(50) NOT NULL,
 `ROOM_CAP` VARCHAR(50) NOT NULL,
 `CLASS_ID` INT NOT NULL,
 PRIMARY KEY (`ROOM_ID`),
 CONSTRAINT `CLASS_ROOM_FK` FOREIGN KEY (`CLASS_ID`) REFERENCES `COLLEGE`.`CLASS` (`CLASS_ID`)
 ON DELETE RESTRICT
 ON UPDATE RESTRICT);
--
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
--------------------------------------------------------------------------------------
/*3. Hãy viết câu query để tìm:
a những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
b những course (distinct) mà 1 professor cụ thể đang dạy
c những course (distinct) mà 1 student cụ thể đang học
d điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
e điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
f điểm số trung bình của các class (quy ra lại theo chữ cái)
g điểm số trung bình của các course (quy ra lại theo chữ cái)*/
--------------------------------------------------------------------------------------
a-
select p.PROF_ID, p.PROF_FNAME, s.STUD_ID, s.STUD_FNAME, c.CLASS_ID, c.CLASS_NAME from PROFESSOR as p 
    left join CLASS as c on p.PROF_ID = c.PROF_ID
    left join ENROLL as e on c.CLASS_ID = e.CLASS_ID
    left join STUDENT as s on e.STUD_ID = s.STUD_ID;
------------------------------------
b-
select distinct p.PROF_ID, p.PROF_FNAME, co.COURSE_ID, co.COURSE_NAME from PROFESSOR as p
    left join CLASS as c on p.PROF_ID = c.PROF_ID
    left join COURSE as co on c.COURSE_ID = co.COURSE_ID;
------------------------------------
c-
select distinct s.STUD_ID, p.STUD_FNAME, co.COURSE_ID, co.COURSE_NAME from STUDENT as s
    left join ENROLL as e on s.STUD_ID = e.STUD_ID
    left join CLASS as c on e.CLASS_ID = c.CLASS_ID
    left join COURSE as co on c.COURSE_ID = co.COURSE_ID;
------------------------------------
d-/*ENROLL as e*/
CASE
    WHEN e.GRADE = 'A' THEN 10
    WHEN e.GRADE = 'B' THEN 8
    WHEN e.GRADE = 'C' THEN 6
    WHEN e.GRADE = 'D' THEN 4
    WHEN e.GRADE = 'E' THEN 2
    ELSE 0
END
------------------------------------
e-
select new.STUD_ID, new.STUD_FNAME, (
    CASE
        WHEN average(new.number_grade) >= 8 THEN "Good"
        WHEN average(new.number_grade) >= 5 AND average(new.number_grade) <8 THEN "Average"
        ELSE "Weak"
    END as average_rank
) from (
    select s.STUD_ID, s.STUD_FNAME
    CASE
        WHEN e.GRADE = 'A' THEN 10
        WHEN e.GRADE = 'B' THEN 8
        WHEN e.GRADE = 'C' THEN 6
        WHEN e.GRADE = 'D' THEN 4
        WHEN e.GRADE = 'E' THEN 2
        ELSE 0
    END as number_grade
    FROM STUDENT as s left join ENROLL as e on s.STUD_ID = e.STUD_ID
    where s.STUD_ID = xxx
) as new group by new.STUD_ID;
------------------------------------
f-
select new.CLASS_ID, new.CLASS_NAME, (
    CASE
        WHEN average(new.number_grade) >= 8 THEN "Good"
        WHEN average(new.number_grade) >= 5 AND average(new.number_grade) <8 THEN "Average"
        ELSE "Weak"
    END as average_rank
) from (
    select c.CLASS_ID, c.CLASS_NAME
    CASE
        WHEN e.GRADE = 'A' THEN 10
        WHEN e.GRADE = 'B' THEN 8
        WHEN e.GRADE = 'C' THEN 6
        WHEN e.GRADE = 'D' THEN 4
        WHEN e.GRADE = 'E' THEN 2
        ELSE 0
    END as number_grade
    FROM CLASS as c left join ENROLL as e on c.CLASS_ID = e.CLASS_ID
) as new group by new.CLASS_ID;
------------------------------------
g-
select new.COURSE_ID, new.COURSE_NAME, (
    CASE
        WHEN average(new.number_grade) >= 8 THEN "Good"
        WHEN average(new.number_grade) >= 5 AND average(new.number_grade) <8 THEN "Average"
        ELSE "Weak"
    END as average_rank
) from (
    select co.COURSE_ID, co.COURSE_NAME
    CASE
        WHEN e.GRADE = 'A' THEN 10
        WHEN e.GRADE = 'B' THEN 8
        WHEN e.GRADE = 'C' THEN 6
        WHEN e.GRADE = 'D' THEN 4
        WHEN e.GRADE = 'E' THEN 2
        ELSE 0
    END as number_grade
    FROM COURSE as co
    left join CLASS as c on co.COURSE_ID = c.COURSE__ID
    left join ENROLL as e on c.CLASS_ID = e.CLASS_ID
) as new group by new.COURSE_ID;