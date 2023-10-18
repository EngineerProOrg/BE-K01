create table PROFESSOR (
    PROF_ID int primary key,
    PROF_LNAME varchar(50),
    PROF_FNAME varchar(50)
)

create table STUDENT (
    STUD_ID int primary key,
    STUD_LNAME varchar(50),
    STUD_FNAME varchar(50),
    STUD_STREET varchar(255),
    STUD_CITY varchar(50),
    STUD_ZIP varchar(10)
)

create table CLASS (
    CLASS_ID int primary key,
    CLASS_NAME varchar(255),,
    PROF_ID int,
    COURSE_ID int,
    ROOM_ID int,
    foreign key (PROF_ID) references PROFESSOR(PROF_ID),
    foreign key (COURSE_ID) references COURSE(COURSE_ID),
    foreign key (ROOM_ID) references ROOM(ROOM_ID)
)

create table ENROLL(
    STUD_ID int,
    CLASS_ID int,
    GRADE varchar(3),
    foreign key (STUD_ID) references STUDENT(STUD_ID),
    foreign key (CLASS_ID) references CLASS(CLASS_ID),
    PRIMARY KEY (STUD_ID, CLASS_ID)
)


create table COURSE (
    COURSE_ID int primary key,
    COURSE_NAME varchar(255)
)

create table ROOM (
    ROOM_ID int primary key,
    ROOM_LOC varchar(50),
    ROOM_CAP varchar(50),
    CLASS_ID int
)

ALTER TABLE ROOM
ADD CONSTRAINT FK_ROOM_CLASS FOREIGN KEY (CLASS_ID) REFERENCES CLASS(CLASS_ID)