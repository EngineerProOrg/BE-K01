-- Những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
select
    distinct stud_id,
    prof_id,
    count(*) as num_class
from
    Class
    join Enroll using (class_id)
group by
    stud_id,
    prof_id
order by
    stud_id;

-- Những course (distinct) mà 1 professor cụ thể đang dạy
select
    distinct course_id
from
    Class
where
    prof_id = 3;

-- Những course (distinct) mà 1 student cụ thể đang học
select
    distinct class_id
from
    Enroll
    join Class using (class_id)
where
    stud_id = 3;

-- điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
create
or replace view Enroll_Grade as
select
    stud_id,
    class_id,
    (
        case
            when grade = "A" then 10
            when grade = "B" then 8
            when grade = "C" then 6
            when grade = "D" then 4
            when grade = "E" then 2
            else 0
        end
    ) as score
from
    Enroll;

select
    *
from
    Enroll_Grade;

-- điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
select
    stud_id,
    avg(score) as average_score,
    (
        case
            when avg(score) >= 8 then "Good"
            when avg(score) >= 5 then "Average"
            else "Weak"
        end
    ) as classification
from
    Enroll_Grade
where
    stud_id = 2;

-- điểm số trung bình của các class (quy ra lại theo chữ cái)
select
    class_id,
    class_name,
    avg(score) as average_score,
    (
        case
            when avg(score) >= 9 then "A"
            when avg(score) >= 8 then "B"
            when avg(score) >= 6 then "C"
            when avg(score) >= 4 then "D"
            when avg(score) >= 2 then "E"
            else "F"
        end
    ) as average_grade
from
    Enroll_Grade
    join Class using (class_id)
group by
    class_id;

-- điểm số trung bình của các course (quy ra lại theo chữ cái)
select
    course_id,
    course_name,
    avg(score) as average_score,
    (
        case
            when avg(score) >= 9 then "A"
            when avg(score) >= 8 then "B"
            when avg(score) >= 6 then "C"
            when avg(score) >= 4 then "D"
            when avg(score) >= 2 then "E"
            else "F"
        end
    ) as average_grade
from
    Enroll_Grade
    join Class using (class_id)
    join Course using (course_id)
group by
    course_id;