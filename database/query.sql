## 1
## những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
select
  concat(p.prof_fname, ' ', p.prof_lname) prof_name,
  concat(s.stud_fname, ' ', s.stud_lname) stud_name,
  count(*) number_of_classes
from
  student s
  join enroll e on s.stud_id = e.stud_id
  join class c on c.class_id = e.class_id
  join professor p on p.prof_id = c.prof_id
group by
  prof_name,
  stud_name;

## 2
## những course (distinct) mà 1 professor cụ thể đang dạy
select
  distinct *
from
  course
  join class c on course.course_id = c.course_id
  join professor p on c.prof_id = p.prof_id
where
  p.prof_id = 1;

## 3
## những course (distinct) mà 1 student cụ thể đang học
select
  distinct course.course_id,
  course_name
from
  course
  join class c on course.course_id = c.course_id
  join enroll e on c.class_id = e.class_id
  join student s on e.stud_id = s.stud_id
where
  e.stud_id = 5;

## 4
## điểm số trung bình của 1 học sinh cụ thể, quy ra lại theo chữ cái và học lực
select
  case
    when avg(score) >= 9 then 'A'
    when avg(score) >= 7 then 'B'
    when avg(score) >= 5 then 'C'
    when avg(score) >= 3 then 'D'
    else 'F'
  end as avg_score,
  case
    when avg(score) >= 8 then 'GOOD'
    when avg(score) < 5 then 'WEAK'
    else 'AVERAGE'
  end as ranked
from
  enroll e1
  join (
    select
      stud_id,
      case
        when grade = 'A' then 10
        when grade = 'B' then 8
        when grade = 'C' then 6
        when grade = 'D' then 4
        when grade = 'E' then 2
        when grade = 'F' then 0
      end as score
    from
      enroll
  ) e2 on e1.stud_id = e2.stud_id
where
  e1.stud_id = 2
group by
  e1.stud_id;

## 5
## điểm số trung bình của các class (quy ra lại theo chữ cái)
select
  e1.class_id class_id,
  case
    when avg(score) >= 9 then 'A'
    when avg(score) >= 7 then 'B'
    when avg(score) >= 5 then 'C'
    when avg(score) >= 3 then 'D'
    else 'F'
  end as avg_score
from
  enroll e1
  join (
    select
      class_id,
      case
        when grade = 'A' then 10
        when grade = 'B' then 8
        when grade = 'C' then 6
        when grade = 'D' then 4
        when grade = 'E' then 2
        when grade = 'F' then 0
      end as score
    from
      enroll
  ) e2 on e1.class_id = e2.class_id
group by
  e1.class_id;

## 6
## điểm số trung bình của các course (quy ra lại theo chữ cái)
select
  cre.course_id,
  cre.course_name,
  case
    when avg(score) >= 9 then 'A'
    when avg(score) >= 7 then 'B'
    when avg(score) >= 5 then 'C'
    when avg(score) >= 3 then 'D'
    else 'F'
  end as avg_score
from
  enroll e1
  join (
    select
      class_id,
      case
        when grade = 'A' then 10
        when grade = 'B' then 8
        when grade = 'C' then 6
        when grade = 'D' then 4
        when grade = 'E' then 2
        when grade = 'F' then 0
      end as score
    from
      enroll
  ) e2 on e1.class_id = e2.class_id
  join class c on c.class_id = e1.class_id
  join course cre on cre.course_id = c.course_id
group by
  cre.course_id,
  cre.course_name