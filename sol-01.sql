-----------------------------------------------------------------------------
/*
1. Giải bài tập trên leetcode: 

https://leetcode.com/problems/capital-gainloss/description/ (gợi ý: sử dụng CASE)
*/
----------------------------------------------------------------------------

/*
    My way
*/
SELECT stock_name, SUM(
  IF(operation = 'Buy', -price, price)
) as capital_gain_loss 
FROM Stocks 
GROUP BY stock_name

/*
    Submitted way
*/
SELECT stock_name, SUM(
  CASE
      WHEN operation = 'Buy' THEN -price
      ELSE price
  END
) AS capital_gain_loss 
FROM Stocks 
GROUP BY stock_name

-----------------------------------------------------------------------------
/*
https://leetcode.com/problems/count-salary-categories/
*/
-----------------------------------------------------------------------------

/*
    My way, but change the structure of the table, if there is no account of a category, it will not be shown in the result table
*/
SELECT (
    CASE
        WHEN income < 20000 THEN "Low Salary"
        WHEN income > 50000 THEN "High Salary"
        ELSE "Average Salary"
    END
) AS category, COUNT(1) AS accounts_count
FROM Accounts
GROUP BY category

/*
    Submitted way
*/
SELECT "Low Salary" AS category, COUNT(1) AS accounts_count
FROM Accounts
WHERE income < 20000
UNION
SELECT "Average Salary" AS category, COUNT(1)
FROM Accounts
WHERE income >= 20000 AND income <= 50000
UNION
SELECT "High Salary" AS category, COUNT(1)
FROM Accounts
WHERE income > 50000

-----------------------------------------------------------------------------
/*
2. Bạn hãy viết một script để tạo các bản cho hệ thống với cấu trúc ở dưới: 

- class
- professor: quan hệ one-many với class
- student: quan hệ many-many với class
- course: quan hệ one-many với class
- room: quan hệ one-one với class
*/
----------------------------------------------------------------------------

DROP DATABASE IF EXISTS `UNIVERSITY`;
CREATE DATABASE `UNIVERSITY`;
USE `UNIVERSITY`;
CREATE TABLE IF NOT EXISTS PROFESSOR (
    PROF_ID     INT             PRIMARY KEY AUTO_INCREMENT,
    PROF_LNAME  VARCHAR(50)     NOT NULL,
    PROF_FNAME  VARCHAR(50)     NOT NULL
)

CREATE TABLE IF NOT EXISTS STUDENT (
    STUD_ID     INT             PRIMARY KEY AUTO_INCREMENT,
    STUD_LNAME  VARCHAR(50)     NOT NULL,
    STUD_FNAME  VARCHAR(50)     NOT NULL,
    STUD_STREET VARCHAR(255)    NOT NULL,
    STUD_CITY   VARCHAR(50)     NOT NULL,
    STUD_ZIP    VARCHAR(10)     NOT NULL
)

CREATE TABLE IF NOT EXISTS COURSE (
    COURSE_ID   INT             PRIMARY KEY AUTO_INCREMENT,
    COURSE_NAME VARCHAR(255)    NOT NULL
)

CREATE TABLE IF NOT EXISTS CLASS (
    CLASS_ID    INT             PRIMARY KEY AUTO_INCREMENT,
    CLASS_NAME  VARCHAR(255)    NOT NULL,
    PROF_ID     INT             NOT NULL,
    COURSE_ID   INT             NOT NULL,
    ROOM_ID     INT             NOT NULL,
    FOREIGN KEY (COURSE_ID)     REFERENCES COURSE(COURSE_ID),
    FOREIGN KEY (PROF_ID)       REFERENCES PROFESSOR(PROF_ID),
    FOREIGN KEY (ROOM_ID)       REFERENCES ROOM(ROOM_ID)
)

CREATE TABLE IF NOT EXISTS ROOM (
    ROOM_ID     INT             PRIMARY KEY AUTO_INCREMENT,
    ROOM_LOC    VARCHAR(50)     NOT NULL,
    ROOM_CAP    VARCHAR(50)     NOT NULL,
    CLASS_ID    INT             NULL,
    FOREIGN KEY (CLASS_ID)      REFERENCES CLASS(CLASS_ID)
)

CREATE TABLE IF NOT EXISTS ENROLL (
    STUD_ID     INT             NOT NULL,
    CLASS_ID    INT             NOT NULL,
    GRADE       VARCHAR(3)      NOT NULL,
    FOREIGN KEY (STUD_ID)       REFERENCES STUDENT(STUD_ID),
    FOREIGN KEY (CLASS_ID)      REFERENCES CLASS(CLASS_ID),
    PRIMARY KEY (STUD_ID, CLASS_ID)
)

-------------------------------------------------------------------------------
/*
3. Hãy viết câu query để tìm:
a. những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
b. những course (distinct) mà 1 professor cụ thể đang dạy
c. những course (distinct) mà 1 student cụ thể đang học
d. điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
e. điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
f. điểm số trung bình của các class (quy ra lại theo chữ cái)
g. điểm số trung bình của các course (quy ra lại theo chữ cái)
*/
--------------------------------------------------------------------------------

/* a. những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
*/
SELECT ENROLL.`STUD_ID`, PROFESSOR.`PROF_ID`, COUNT(1) AS COMMON_CLASS FROM `ENROLL` 
INNER JOIN `CLASS` ON CLASS.`CLASS_ID` = ENROLL.`CLASS_ID`
INNER JOIN `PROFESSOR` ON PROFESSOR.`PROF_ID` = CLASS.`PROF_ID`
GROUP BY ENROLL.`STUD_ID`, PROFESSOR.`PROF_ID`;

/* b. những course (distinct) mà 1 professor cụ thể đang dạy
*/
SELECT DISTINCT PROF.`PROF_ID`, COURSE.`COURSE_ID` FROM (
    SELECT PROF_ID
    FROM `PROFESSOR`
    WHERE PROFESSOR.`PROF_ID` = 1
) PROF
INNER JOIN `CLASS` ON CLASS.`PROF_ID` = PROF.`PROF_ID`
INNER JOIN `COURSE` ON COURSE.`COURSE_ID` = CLASS.`COURSE_ID`;

/* c. những course (distinct) mà 1 student cụ thể đang học
*/
SELECT DISTINCT STU.`STUD_ID`, CLASS.`COURSE_ID` FROM (
    SELECT STUD_ID
    FROM `STUDENT`
    WHERE STUDENT.`STUD_ID` = 1
) STU
INNER JOIN `ENROLL` ON ENROLL.`STUD_ID` = STU.`STUD_ID`
INNER JOIN `CLASS` ON CLASS.`CLASS_ID` = ENROLL.`CLASS_ID`;

/* d. điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
*/
SELECT `STUD_ID`, `CLASS_ID`, (CASE 
    WHEN GRADE = 'F' THEN 0
    WHEN GRADE = 'E' THEN 2
    WHEN GRADE = 'D' THEN 4
    WHEN GRADE = 'C' THEN 6
    WHEN GRADE = 'B' THEN 8
    ELSE 10
END) AS SCORE
FROM `ENROLL`;

/* e. điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
*/
SELECT `STUD_ID`, AVG_SCORE, (
    CASE
        WHEN AVG_SCORE < 5 THEN 'weak'
        WHEN AVG_SCORE >= 8 THEN 'good'
        ELSE 'average'
    END
) AS `TYPE`
FROM (
    SELECT `STUD_ID`, AVG(
        CASE 
            WHEN GRADE = 'F' THEN 0
            WHEN GRADE = 'E' THEN 2
            WHEN GRADE = 'D' THEN 4
            WHEN GRADE = 'C' THEN 6
            WHEN GRADE = 'B' THEN 8
            ELSE 10
        END
    ) AS AVG_SCORE
    FROM `ENROLL`
    WHERE `STUD_ID` = 1
    GROUP BY `STUD_ID`
) STU;

/* f. điểm số trung bình của các class (quy ra lại theo chữ cái)
*/
SELECT `CLASS_ID`, `CLASS_NAME`, AVG_SCORE, (
    CASE
        WHEN AVG_SCORE < 2 THEN 'F'
        WHEN AVG_SCORE >= 2 AND AVG_SCORE < 4 THEN 'E'
        WHEN AVG_SCORE >= 4 AND AVG_SCORE < 6 THEN 'D'
        WHEN AVG_SCORE >= 6 AND AVG_SCORE < 8 THEN 'C'
        WHEN AVG_SCORE >= 8 AND AVG_SCORE < 10 THEN 'B'
        ELSE 'A'
    END
) AS AVG_GRADE
FROM (
    SELECT CLASS.`CLASS_ID`, CLASS.`CLASS_NAME`, AVG(
        CASE 
            WHEN ENROLL.`GRADE` = 'F' THEN 0
            WHEN ENROLL.`GRADE` = 'E' THEN 2
            WHEN ENROLL.`GRADE` = 'D' THEN 4
            WHEN ENROLL.`GRADE` = 'C' THEN 6
            WHEN ENROLL.`GRADE` = 'B' THEN 8
            ELSE 10
        END
    ) AS AVG_SCORE
    FROM `CLASS` 
    INNER JOIN `ENROLL` ON ENROLL.`CLASS_ID` = CLASS.`CLASS_ID`
    GROUP BY CLASS.`CLASS_ID`
) TB;

/* điểm số trung bình của các course (quy ra lại theo chữ cái)
*/
SELECT `COURSE_ID`, `COURSE_NAME`, AVG_SCORE, (
    CASE
        WHEN AVG_SCORE < 2 THEN 'F'
        WHEN AVG_SCORE >= 2 AND AVG_SCORE < 4 THEN 'E'
        WHEN AVG_SCORE >= 4 AND AVG_SCORE < 6 THEN 'D'
        WHEN AVG_SCORE >= 6 AND AVG_SCORE < 8 THEN 'C'
        WHEN AVG_SCORE >= 8 AND AVG_SCORE < 10 THEN 'B'
        ELSE 'A'
    END
) AS AVG_GRADE
FROM (
    SELECT COURSE.`COURSE_ID`, COURSE.`COURSE_NAME`, AVG(
        CASE 
            WHEN ENROLL.`GRADE` = 'F' THEN 0
            WHEN ENROLL.`GRADE` = 'E' THEN 2
            WHEN ENROLL.`GRADE` = 'D' THEN 4
            WHEN ENROLL.`GRADE` = 'C' THEN 6
            WHEN ENROLL.`GRADE` = 'B' THEN 8
            ELSE 10
        END
    ) AS AVG_SCORE
    FROM `COURSE`
    INNER JOIN `CLASS` ON CLASS.`COURSE_ID` = COURSE.`COURSE_ID`
    INNER JOIN `ENROLL` ON ENROLL.`CLASS_ID` = CLASS.`CLASS_ID`
    GROUP BY COURSE.`COURSE_ID`
) TB;


