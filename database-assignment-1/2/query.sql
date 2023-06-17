-- những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
SELECT CONCAT(professor.prof_fname, ' ', professor.prof_lname) as prof_name, CONCAT(student.stud_fname, ' ', student.stud_lname) as student_name , COUNT(*) as total_class  FROM student
INNER JOIN enroll ON student.stud_id = enroll.stud_id
INNER JOIN `class` ON enroll.class_id = `class`.class_id
INNER JOIN professor ON professor.prof_id = `class`.prof_id
GROUP BY prof_name, student_name;

-- những course (distinct) mà 1 professor cụ thể đang dạy
-- vi du nhu prof_id=1
SELECT DISTINCT course.course_id, course.course_name FROM `class`
INNER JOIN professor ON `class`.prof_id = professor.prof_id
INNER JOIN course ON `class`.course_id = course.course_id
WHERE professor.prof_id=1;

-- những course (distinct) mà 1 student cụ thể đang học
SELECT DISTINCT course.course_id, course_name from student
INNER JOIN enroll ON enroll.stud_id = student.stud_id
INNER JOIN class ON enroll.class_id = `class`.class_id
INNER JOIN course ON `class`.course_id = course.course_id
WHERE student.stud_id=8;

-- điểm số trung bình của 1 học sinh cụ thể 
-- (quy ra lại theo chữ cái, 
-- và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
SELECT
	temp.student_name,
	CASE WHEN ROUND(AVG(grade_num), 0) > 8 THEN
		'A'
	WHEN ROUND(AVG(grade_num), 0) > 6 THEN
		'B'
	WHEN ROUND(AVG(grade_num), 0) > 4 THEN
		'C'
	WHEN ROUND(AVG(grade_num), 0) > 2 THEN
		'D'
	WHEN ROUND(AVG(grade_num), 0) > 0 THEN
		'E'
	ELSE
		'F'
	END AS avg_grade,
	CASE WHEN ROUND(AVG(grade_num), 0) >= 8 THEN
		'good'
	WHEN ROUND(AVG(grade_num), 0) >= 5 THEN
		'average'
	ELSE
		'weak'
	END AS _rank
FROM (
	SELECT
		CONCAT(student.stud_fname, ' ', student.stud_lname) AS student_name,
		CASE WHEN grade = 'A' THEN
			10
		WHEN grade = 'B' THEN
			8
		WHEN grade = 'C' THEN
			6
		WHEN grade = 'D' THEN
			4
		WHEN grade = 'E' THEN
			2
		WHEN grade = 'F' THEN
			0
		END AS grade_num
	FROM
		student
		INNER JOIN enroll ON enroll.stud_id = student.stud_id) AS temp
GROUP BY
	temp.student_name;

-- điểm số trung bình của các class (quy ra lại theo chữ cái)
SELECT
	temp.class_id,
	temp.class_name,
	CASE WHEN ROUND(avg(temp.grade_num), 0) > 8 THEN
		'A'
	WHEN ROUND(avg(temp.grade_num), 0) > 6 THEN
		'B'
	WHEN ROUND(avg(temp.grade_num), 0) > 4 THEN
		'C'
	WHEN ROUND(avg(temp.grade_num), 0) > 2 THEN
		'D'
	WHEN ROUND(avg(temp.grade_num), 0) > 0 THEN
		'E'
	ELSE 'F'
	END avg_grade
FROM (
	SELECT
		`class`.class_id,
		`class`.class_name,
		CASE WHEN grade = 'A' THEN
			10
		WHEN grade = 'B' THEN
			8
		WHEN grade = 'C' THEN
			6
		WHEN grade = 'D' THEN
			4
		WHEN grade = 'E' THEN
			2
		WHEN grade = 'F' THEN
			0
		END AS grade_num
	FROM
		`class`
		INNER JOIN enroll ON enroll.class_id = `class`.class_id) AS temp
GROUP BY
	temp.class_id;

-- điểm số trung bình của các course (quy ra lại theo chữ cái)
SELECT temp.course_id, temp.course_name,
CASE
	WHEN ROUND(AVG(temp.grade), 0) > 8 THEN 'A'
	WHEN ROUND(AVG(temp.grade), 0) > 6 THEN 'B'
	WHEN ROUND(AVG(temp.grade), 0) > 4 THEN 'C'
	WHEN ROUND(AVG(temp.grade), 0) > 2 THEN 'D'
	WHEN ROUND(AVG(temp.grade), 0) > 0 THEN 'E'
	ELSE 'F'
END AS avg_grade
FROM
(
SELECT course.course_id, course.course_name, 
CASE
	WHEN enroll.grade = 'A' THEN 10
	WHEN enroll.grade = 'B' THEN 8
	WHEN enroll.grade = 'C' THEN 6
	WHEN enroll.grade = 'D' THEN 4
	WHEN enroll.grade = 'E' THEN 2
	ELSE 0
END AS grade
FROM `class`
INNER JOIN enroll ON enroll.class_id = `class`.class_id
INNER JOIN course ON course.course_id = `class`.course_id
) AS temp
GROUP BY temp.course_id;