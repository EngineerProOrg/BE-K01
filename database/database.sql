create table if not exists professor (
  prof_id integer primary key,
  prof_lname varchar(50),
  prof_fname varchar(50)
);

create table if not exists student (
  stud_id integer primary key,
  stud_lname varchar(50),
  stud_fname varchar(50),
  stud_street varchar(255),
  stud_city varchar(50),
  stud_zip varchar(10)
);

create table if not exists enroll (
  stud_id integer,
  class_id integer,
  grade enum ('A', 'B', 'C', 'D', 'F')
);

create table if not exists class (
  class_id integer primary key,
  class_name varchar(255),
  prof_id integer,
  course_id integer,
  room_id integer
);

create table if not exists room (
  room_id integer primary key,
  room_loc varchar(50),
  room_cap varchar(50),
  class_id integer
);

create table if not exists course (
  course_id integer primary key,
  course_name varchar(255)
);

alter table
  enroll
add
  constraint fk_student_enroll foreign key (stud_id) references student(stud_id);

alter table
  enroll
add
  constraint fk_class_enroll foreign key (class_id) references class(class_id);

alter table
  enroll
add
  primary key (stud_id, class_id);

alter table
  class
add
  constraint fk_prof_class foreign key (prof_id) references professor(prof_id);

alter table
  class
add
  constraint fk_course_class foreign key (course_id) references course(course_id);

alter table
  class
add
  constraint fk_room_class foreign key (room_id) references room(room_id);

alter table
  room
add
  constraint fk_class_room foreign key (class_id) references class(class_id)