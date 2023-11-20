-- những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
select  
	*,
	count(t1.stud_id) as num_class
from 
(
	select
		e.stud_id,
		CONCAT(s.stud_lname, ' ', s.stud_fname) as student_name,
		CONCAT(p.prof_lname, ' ',p.prof_fname) as prof_name
	FROM enrolls e
	inner join students s on s.stud_id = e.stud_id
	inner join classes c on c.class_id = e.class_id
	inner join professors p on p.prof_id = c.prof_id
) as t1
GROUP BY t1.stud_id, t1.student_name, t1.prof_name;

-- những course (distinct) mà 1 professor cụ thể đang dạy
select CONCAT(p.prof_lname, ' ',p.prof_fname) as prof_name,  c2.course_name  FROM 
professors p 
inner join classes c on c.prof_id = p.prof_id 
inner join courses c2 on c2.course_id = c.course_id 

-- những course (distinct) mà 1 student cụ thể đang học
select 
	s.course_id , s.course_name 
FROM courses s 
inner join classes c on c.course_id  = s.course_id  
inner join enrolls e on e.class_id = c.class_id 
inner join students s2 on e.stud_id = s2.stud_id 
group by s.course_id ,s.course_name 

-- điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
SELECT s.stud_lname , s.stud_fname , c.class_name ,
CASE 
	WHEN grade = 'A' THEN 10
	WHEN grade = 'B' THEN 8
	WHEN grade = 'C' THEN 6
	WHEN grade = 'D' THEN 4
	WHEN grade = 'E' THEN 2
	WHEN grade = 'F' THEN 0
END as point
from students s 
inner join enrolls e on e.stud_id = s.stud_id 
inner JOIN classes c on e.class_id = c.class_id 
-- điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
SELECT s.stud_id , 
CASE 
	WHEN ROUND(AVG(t1.point), 1) < 5 THEN 'weak'
	WHEN ROUND(AVG(t1.point), 1) > 5 AND ROUND(AVG(t1.point), 1) < 8 THEN 'weak'
	ELSE 'good'
END AS ranking

from students s
inner join 
(
	SELECT class_id,stud_id ,
		CASE 
			WHEN grade = 'A' THEN 10
			WHEN grade = 'B' THEN 8
			WHEN grade = 'C' THEN 6
			WHEN grade = 'D' THEN 4
			WHEN grade = 'E' THEN 2
			WHEN grade = 'F' THEN 0
		END as point
	FROM enrolls e 
) as t1
on s.stud_id = t1.stud_id
GROUP BY s.stud_id 
-- điểm số trung bình của các class (quy ra lại theo chữ cái)
select t2.class_id,c.class_name , AVG(point) as avg_point  from (
	SELECT t1.class_id, point 
	FROM 
		(
			SELECT class_id,
				CASE 
					WHEN grade = 'A' THEN 10
					WHEN grade = 'B' THEN 8
					WHEN grade = 'C' THEN 6
					WHEN grade = 'D' THEN 4
					WHEN grade = 'E' THEN 2
					WHEN grade = 'F' THEN 0
				END as point
			FROM enrolls e 
		) as t1
) as t2
inner join classes c on c.class_id  = t2.class_id
group by t2.class_id

-- điểm số trung bình của các course (quy ra lại theo chữ cái)
SELECT  c.course_id ,course_name , 
CASE 
	WHEN ROUND(AVG(t1.point), 1) < 5 THEN 'weak'
	WHEN ROUND(AVG(t1.point), 1) > 5 AND ROUND(AVG(t1.point), 1) < 8 THEN 'weak'
	ELSE 'good'
END AS ranking
FROM courses c 
INNER JOIN classes c2 ON c.course_id = c2.course_id  
INNER JOIN 
(
	SELECT class_id,
				CASE 
					WHEN grade = 'A' THEN 10
					WHEN grade = 'B' THEN 8
					WHEN grade = 'C' THEN 6
					WHEN grade = 'D' THEN 4
					WHEN grade = 'E' THEN 2
					WHEN grade = 'F' THEN 0
				END as point
			FROM enrolls e 
) as t1
on t1.class_id = c2.class_id  
GROUP BY c.course_id 