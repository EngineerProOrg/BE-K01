-- 1
select concat(p.fname, ' ', p.lname) as 'Professor Name',
      concat(s.fname, ' ', s.lname) as 'Student Name',
      count(c.id) as 'Number of Related Classes'
from Professors p
join Classes c on c.prof_id = p.id
join Enrolls e on e.class_id = c.id
join Students s on s.id = e.stud_id
group by p.id, s.id;

-- 2
select distinct(co.id)
from Professors p
join Classes cl on cl.prof_id = p.id
join Courses co on co.class_id = cl.id
group by p.id;

-- 3
select distinct(co.id)
from Students s
join Enrolls e on e.stud_id = s.id
join Classes cl on cl.id = e.class_id
join Courses co on co.id = cl.course_id
group by s.id;

-- 4
create or replace view Grade2Score as
select stud_id, class_id, 
(
case
  when grade = 'A' then 10
  when grade = 'B' then 8
  when grade = 'C' then 6
  when grade = 'D' then 4
  when grade = 'E' then 2
  else 0
end
) as score
from Enrolls;

select * from Grade2Score;

-- 5
select
avg(score) as avg_score,
(
case
  when avg(score) >= 8 then 'Good'
  when avg(score) >= 5 then 'Average'
  else 'Weak'
end
) as classification
from Grade2Score
where stud_id = 1;

-- 6
select class_id,
avg(score) as avg_score,
(
case
  when avg(score) >= 9 then 'A'
  when avg(score) >= 8 then 'B'
  when avg(score) >= 6 then 'C'
  when avg(score) >= 4 then 'D'
  when avg(score) >= 2 then 'E'
  else 'F'
end
) as classification
from Grade2Score
group by class_id;

-- 7
select co.id,
avg(gs.score) as avg_score,
(
case
  when avg(gs.score) >= 9 then 'A'
  when avg(gs.score) >= 8 then 'B'
  when avg(gs.score) >= 6 then 'C'
  when avg(gs.score) >= 4 then 'D'
  when avg(gs.score) >= 2 then 'E'
  else 'F'
end
) as classification
from Courses co
join Classes cl on cl.course_id = co.id
join Grade2Score gs on gs.class_id = cl.id
group by co.id;