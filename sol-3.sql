'''
Hãy viết câu query để tìm:
những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
những course (distinct) mà 1 professor cụ thể đang dạy
những course (distinct) mà 1 student cụ thể đang học
điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
điểm số trung bình của các class (quy ra lại theo chữ cái)
điểm số trung bình của các course (quy ra lại theo chữ cái)
'''

--những cặp student-professor có dạy học nhau và số lớp mà họ có liên quan
select pr.PROF_ID, st.STUD_ID, cl.CLASS_ID
from PROFESSOR pr inner join CLASS cl on pr.PROF_ID = cl.PROF_ID
                inner join ENROLL en on cl.CLASS_ID = en.CLASS_ID
                inner join STUDENT st on en.STUD_ID = st.STUD_ID

--những course (distinct) mà 1 professor cụ thể đang dạy. VD professor có id = 1
select distinct co.COURSE_NAME
from PROFESSOR pr inner join CLASS cl on pr.PROF_ID = cl.PROF_ID
                inner join COURSE co on cl.COURSE_ID = co.COURSE_ID
where pr.PROF_ID = 1

--những course (distinct) mà 1 student cụ thể đang học. VD student có id = 1
select distinct co.COURSE_NAME
from STUDENT st inner join ENROLL en on st.STUD_ID = en.STUD_ID
                inner join CLASS cl on en.CLASS_ID = cl.CLASS_ID
                inner join COURSE co on cl.COURSE_ID = co.COURSE_ID
where st.STUD_ID = 1

--điểm số là A, B, C, D, E, F tương đương với 10, 8, 6, 4, 2, 0
select (
    case 
        when GRADE = 'A' then 10
        when GRADE = 'B' then 8
        when GRADE = 'C' then 6
        when GRADE = 'D' then 4
        when GRADE = 'E' then 2
        when GRADE = 'F' then 0
    end
) as GRADE
from STUDENT st inner join ENROLL en on st.STUD_ID = en.STUD_ID

--điểm số trung bình của 1 học sinh cụ thể (quy ra lại theo chữ cái, và xếp loại học lực (weak nếu avg < 5, average nếu >=5 < 8, good nếu >=8 )
select (
    case 
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) < 5 then 'weak'
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) >= 5 and avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) < 8 then 'average'
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) >= 8 then 'good'
    end
) as GRADE
from STUDENT st inner join ENROLL en on st.STUD_ID = en.STUD_ID
group by st.STUD_ID

--điểm số trung bình của các class (quy ra lại theo chữ cái)
select cl.CLASS_ID, (
    case 
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) < 5 then 'weak'
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) >= 5 and avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) < 8 then 'average'
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) >= 8 then 'good'
    end
) as GRADE
from CLASS cl inner join ENROLL en on cl.CLASS_ID = en.CLASS_ID 
group by cl.CLASS_ID

--điểm số trung bình của các course (quy ra lại theo chữ cái)

select co.COURSE_NAME, (
    case 
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) < 5 then 'weak'
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) >= 5 and avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) < 8 then 'average'
        when avg(
            case 
                when GRADE = 'A' then 10
                when GRADE = 'B' then 8
                when GRADE = 'C' then 6
                when GRADE = 'D' then 4
                when GRADE = 'E' then 2
                when GRADE = 'F' then 0
            end
        ) >= 8 then 'good'
    end
) as GRADE
from COURSE co inner join CLASS cl on co.COURSE_ID = cl.COURSE_ID
                inner join ENROLL en on cl.CLASS_ID = en.CLASS_ID
group by co.COURSE_ID