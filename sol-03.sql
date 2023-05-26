USE be01_ass1;
-- bai 1
SELECT c.CLASS_NAME, concat(s.STUD_LNAME, '  ', s.STUD_FNAME) as STUD_FULLNAME, concat(p.PROF_FNAME, '  ' , p.PROF_LNAME) as PROF_FULLNAME
FROM enroll as e
JOIN student as s
ON e.STUD_ID = s.STUD_ID
JOIN class as c
ON c.CLASS_ID = e.CLASS_ID
JOIN professor p
ON p.PROF_ID = c.PROF_ID;

-- bai 2
select distinct CONCAT(p.PROF_FNAME, '  ', p.PROF_LNAME) AS PROF_FULLNAME, COURSE_NAME
from professor as p
join class as c on p.PROF_ID= c.PROF_ID
join course on c.COURSE_ID = course.COURSE_ID
group by p.PROF_ID, c.COURSE_ID;

-- bai 3
SELECT distinct concat(s.STUD_FNAME, '  ', s.STUD_LNAME) as fullname, cs.COURSE_NAME
FROM student as s
Join enroll as e
on s.STUD_ID = e.STUD_ID
join class as c
on c.CLASS_ID = e.CLASS_ID
join course as cs
on cs.COURSE_ID = c.COURSE_ID;

-- bai 4,5
SELECT *, 
    CASE
        WHEN tempTable.AVG_GRADE = 10 THEN 'A'
        WHEN tempTable.AVG_GRADE >= 9 THEN 'B'
        WHEN tempTable.AVG_GRADE >= 6 THEN 'C'
        WHEN tempTable.AVG_GRADE >= 4 THEN 'D'
        WHEN tempTable.AVG_GRADE >= 2 THEN 'E'
        ELSE 'F'
    END as NUMBER_GRADE
FROM 
(SELECT e.STUD_ID, concat(s.STUD_FNAME,'  ',s.STUD_LNAME) as FULLNAME, avg(CAST(e.GRADE as DECIMAL)) as AVG_GRADE
FROM enroll as e
JOIN student as s
ON e.STUD_ID = s.STUD_ID
GROUP BY e.STUD_ID) as tempTable;

-- bai 6
select *, 
    CASE
        WHEN tempTable.avg_grade = 10 THEN 'A'
        WHEN tempTable.avg_grade >= 9 THEN 'B'
        WHEN tempTable.avg_grade >= 6 THEN 'C'
        WHEN tempTable.avg_grade >= 4 THEN 'D'
        WHEN tempTable.avg_grade >= 2 THEN 'E'
        ELSE 'F'
    END as NUMBER_GRADE
from (
select c.CLASS_ID, avg(Cast(e.GRADE as DECIMAL)) as avg_grade
from class as c
join enroll as e
on c.CLASS_ID = e.CLASS_ID
group by c.CLASS_ID) as tempTable
order by CLASS_ID;

-- bai 7
select *, 
    CASE
        WHEN tempTable.avg_grade = 10 THEN 'A'
        WHEN tempTable.avg_grade >= 9 THEN 'B'
        WHEN tempTable.avg_grade >= 6 THEN 'C'
        WHEN tempTable.avg_grade >= 4 THEN 'D'
        WHEN tempTable.avg_grade >= 2 THEN 'E'
        ELSE 'F'
    END as NUMBER_GRADE
from (
select c.COURSE_ID, avg(Cast(e.GRADE as DECIMAL)) as avg_grade
from class as c
join enroll as e
on c.CLASS_ID = e.CLASS_ID
join course as cs
on cs.COURSE_ID = c.COURSE_ID
group by c.COURSE_ID) as tempTable
order by COURSE_ID;