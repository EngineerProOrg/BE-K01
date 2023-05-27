-- những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan

select pr.prof_id,
		pr.prof_fname+' '+ pr.prof_lname 'prof name',
		st.stud_id, 
		count(*) as
from student as st
join enroll as en 
	on st.stud_id = en.stud_id
join class as cl 
	on cl.class_id = en.class_id
join progessor as pr
	on pr.prof_id = cl.prof_id
group by pr.prof_id,st.stud_id,pr.prof_fname+' '+ pr.prof_lname

-- những course (distinct) mà 1 professor cụ thể đang dạy
select distinct pr.prof_id, co.course_id, co.course_name 
from progessor as pr
join class as cl 
	on pr.prof_id = cl.prof_id
join course as co
	on cl.course_id = co.course_id
where pr.prof_id = 1

-- cau 3 những course (distinct) mà 1 student cụ thể đang học
select distinct st.stud_id,co.course_id, course_name
from student as st
join enroll as en
	on st.stud_id = en.stud_id
join class as cl
	on cl.class_id = en.class_id
join course as co 
	on cl.course_id = co.course_id
where st.stud_id = 1

-- điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
select *,
CASE
    WHEN grade = 'A' THEN 10
    WHEN grade = 'B' THEN 8
    WHEN grade = 'C' THEN 6
	WHEN grade = 'D' THEN 4
	WHEN grade = 'E' THEN 2
    ELSE 0
END
as score
from enroll

-- điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
select st.stud_id, 
	CASE
            WHEN grade.score < 5 THEN 'Weak'
            WHEN grade.score >= 5 AND grade.score < 8 THEN 'Average'
            WHEN grade.score >= 8 THEN 'Good'
	END
from student as st
join (
	select stud_id,avg(
	CASE
		WHEN grade = 'A' THEN 10
		WHEN grade = 'B' THEN 8
		WHEN grade = 'C' THEN 6
		WHEN grade = 'D' THEN 4
		WHEN grade = 'E' THEN 2
		ELSE 0
	END
	) as 'score'
	from enroll
	group by stud_id
) as grade 
on st.stud_id = grade.stud_id

-- điểm số trung bình của các class (quy ra lại theo chữ cái)
select cl.class_id, 
	case 
		WHEN grade.score < 5 THEN 'Weak'
        WHEN grade.score >= 5 AND grade.score < 8 THEN 'Average'
        WHEN grade.score >= 8 THEN 'Good'
	end
from class as cl
join (
select class_id, avg(
	case
		WHEN grade = 'A' THEN 10
		WHEN grade = 'B' THEN 8
		WHEN grade = 'C' THEN 6
		WHEN grade = 'D' THEN 4
		WHEN grade = 'E' THEN 2
		ELSE 0
	end
) as score
from enroll
group by class_id) as grade
on cl.class_id = grade.class_id


-- điểm số trung bình của các course (quy ra lại theo chữ cái)
select cl.class_id, 
	case 
		WHEN grade.score < 5 THEN 'Weak'
        WHEN grade.score >= 5 AND grade.score < 8 THEN 'Average'
        WHEN grade.score >= 8 THEN 'Good'
	end
from class as cl
join (
select class_id, avg(
	case
		WHEN grade = 'A' THEN 10
		WHEN grade = 'B' THEN 8
		WHEN grade = 'C' THEN 6
		WHEN grade = 'D' THEN 4
		WHEN grade = 'E' THEN 2
		ELSE 0
	end
) as score
from enroll
group by class_id) as grade
on cl.class_id = grade.class_id
