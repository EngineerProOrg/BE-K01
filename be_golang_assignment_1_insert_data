use be_golang;

INSERT INTO professor (prof_lname,prof_fname)
VALUES ("phong", "tran"), ("albert", "einstein"), ("stephen", "hawking");

INSERT INTO student (stud_fname,stud_lname,stud_street,stud_city,stud_zip)
VALUES
("cristiano","ronaldo","ronaldo street","Lisbon","1000"),
("lionel","messi","messi street","Rosario","2000"),
("harry","maguire","maguire street","Sheffield","3000");

INSERT INTO course (course_name)
VALUES ("football"), ("nuclear physics"), ("golang backend"), ("advanced course: how to deal with women");

START TRANSACTION;
SET FOREIGN_KEY_CHECKS=0;
INSERT INTO class(class_name,prof_id,course_id,room_id)
VALUES
("basic", 2, 2, 2),
("advanced", 3, 3, 4),
("optional", 1, 1, 1),
("critical", 1, 3, 3);

INSERT INTO room(room_loc,class_id)
VALUES ("mars", 1), ("earth", 2), ("manu fan's cave", 3), ("engineer pro", 4);
SET FOREIGN_KEY_CHECKS=1;
COMMIT;

INSERT INTO enroll(stud_id,class_id,grade)
VALUES
(1, 1, "1st"),
(1, 2, "2nd"),
(1, 3, "3rd"),
(1, 4, "1st"),
(2, 1, "2nd"),
(2, 2, "3rd"),
(2, 3, "1st"),
(2, 4, "2nd"),
(3, 1, "3rd"),
(3, 2, "1st"),
(3, 3, "2nd"),
(3, 4, "3rd");
