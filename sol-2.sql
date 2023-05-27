-- DROP DATABASE ep_asm1_2;

CREATE DATABASE IF NOT EXISTS ep_asm1_2;
USE ep_asm1_2;

DROP TABLE IF EXISTS `professors`;
CREATE TABLE `professors`(
  `id` INT NOT NULL AUTO_INCREMENT,
  `fname` VARCHAR(50),
  `lname` VARCHAR(50),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `students`;
CREATE TABLE `students`(
  `id` INT NOT NULL AUTO_INCREMENT,
  `fname` VARCHAR(50),
  `lname` VARCHAR(50),
  `street` VARCHAR(255),
  `city` VARCHAR(50),
  `zip` VARCHAR(10),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses`(
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes`(
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `prof_id` INT,
  `course_id` INT,
  `room_id` INT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms`(
  `id` INT NOT NULL AUTO_INCREMENT,
  `loc` VARCHAR(50),
  `cap` VARCHAR(50),
  `class_id` INT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `enrolls`;
CREATE TABLE `enrolls`(
  `stud_id` INT NOT NULL,
  `class_id` INT NOT NULL,
  `grade` VARCHAR(3),
  PRIMARY KEY (`stud_id`, `class_id`)
) ENGINE=InnoDB;


-- Add Foreign Key
-- Classes 
ALTER TABLE Classes
ADD FOREIGN KEY (`prof_id`) REFERENCES Professors(`id`);
ALTER TABLE Classes
ADD FOREIGN KEY (`course_id`) REFERENCES Courses(`id`);
ALTER TABLE Classes
ADD FOREIGN KEY (`room_id`) REFERENCES Rooms(`id`);

-- Rooms 
ALTER TABLE Rooms
ADD FOREIGN KEY (`class_id`) REFERENCES Classes(`id`);

-- Enrols
ALTER TABLE Enrolls
ADD FOREIGN KEY (`stud_id`) REFERENCES Students(`id`);
ALTER TABLE Enrolls
ADD FOREIGN KEY (`class_id`) REFERENCES Classes(`id`);

-- Insert to Professor
insert into Professors values (1, 'Albert', 'Einstein');
insert into Professors values (2, 'Nicolas', 'Tesla');
insert into Professors values (3, 'Andrew', 'Ng');

-- Insert to Student
insert into Students values (1, 'A', 'Nguyen', 'ABC 1', 'HCM', '70000');
insert into Students values (2, 'B', 'Tran', 'DEF 1', 'HCM', '70000');
insert into Students values (3, 'C', 'Pham', 'HIJ 1', 'HCM', '70000');

-- Insert to Course
insert into Courses values (1, 'PE');
insert into Courses values (2, 'Math');
insert into Courses values (3, 'Physics');
insert into Courses values (4, 'Henshin');

-- Insert to Class
insert into Classes values (1, 'PE1', 1, 1, 1);
insert into Classes values (2, 'MA1', 2, 2, 2);
insert into Classes values (3, 'PH1', 3, 3, 3);
insert into Classes values (4, 'HE1', 3, 3, 4);
insert into Classes values (5, 'HE2', 3, 4, 5);

-- Insert to Room
insert into Rooms values (1, '001', '10', 1);
insert into Rooms values (2, '002', '10', 2);
insert into Rooms values (3, '003', '10', 3);
insert into Rooms values (4, '004', '10', 4);
insert into Rooms values (5, '005', '10', 5);

-- Insert to Enroll
insert into Enrolls values (1, 2, 'C');
insert into Enrolls values (2, 1, 'F');
insert into Enrolls values (2, 2, 'A');
insert into Enrolls values (2, 3, 'D');
insert into Enrolls values (3, 1, 'B');
insert into Enrolls values (3, 3, 'E');
insert into Enrolls values (1, 4, 'A');
insert into Enrolls values (2, 4, 'B');
insert into Enrolls values (3, 4, 'C');
