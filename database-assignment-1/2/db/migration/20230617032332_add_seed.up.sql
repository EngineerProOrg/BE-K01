-- Insert records into the professor table
INSERT INTO professor (prof_lname, prof_fname)
VALUES ('Smith', 'John'),
       ('Johnson', 'Lisa'),
       ('Williams', 'Michael'),
       ('Brown', 'Jessica'),
       ('Jones', 'David'),
       ('Davis', 'Emily'),
       ('Anderson', 'Andrew'),
       ('Taylor', 'Olivia'),
       ('Martinez', 'Daniel'),
       ('Miller', 'Sophia');

-- Insert records into the course table
INSERT INTO course (course_name)
VALUES ('Mathematics'),
       ('Physics'),
       ('History'),
       ('English'),
       ('Chemistry'),
       ('Computer Science'),
       ('Biology'),
       ('Art'),
       ('Economics'),
       ('Psychology');

-- Insert records into the room table
INSERT INTO room (room_loc, room_cap)
VALUES ('Building A, Room 101', '30'),
       ('Building B, Room 202', '25'),
       ('Building C, Room 303', '40'),
       ('Building D, Room 404', '20'),
       ('Building E, Room 505', '35'),
       ('Building F, Room 606', '15'),
       ('Building G, Room 707', '50'),
       ('Building H, Room 808', '28'),
       ('Building I, Room 909', '32'),
       ('Building J, Room 1010', '18');

-- Insert records into the class table
INSERT INTO class (class_name, prof_id, course_id, room_id)
VALUES ('Class A', 1, 1, 1),
       ('Class B', 2, 2, 2),
       ('Class C', 3, 3, 3),
       ('Class D', 4, 4, 4),
       ('Class E', 5, 5, 5),
       ('Class F', 6, 6, 6),
       ('Class G', 7, 7, 7),
       ('Class H', 8, 8, 8),
       ('Class I', 9, 9, 9),
       ('Class J', 10, 10, 10);

-- Insert records into the student table
INSERT INTO student (stud_fname, stud_lname, stud_street, stud_city, stud_zip)
VALUES ('Emma', 'Davis', '123 Main St', 'New York', '10001'),
       ('James', 'Johnson', '456 Elm St', 'Los Angeles', '90001'),
       ('Sophia', 'Smith', '789 Oak St', 'Chicago', '60601'),
       ('Oliver', 'Wilson', '234 Pine St', 'Houston', '77001'),
       ('Ava', 'Brown', '567 Maple St', 'Phoenix', '85001'),
       ('Noah', 'Garcia', '890 Cedar St', 'Philadelphia', '19019'),
       ('Isabella', 'Lopez', '901 Walnut St', 'San Antonio', '78201'),
       ('Lucas', 'Lee', '345 Oak St', 'San Diego', '92101'),
       ('Mia', 'Miller', '678 Elm St', 'Dallas', '75201'),
       ('Liam', 'Clark', '912 Pine St', 'San Francisco', '94101');

-- Insert records into the enroll table
INSERT INTO enroll (stud_id, class_id, grade)
VALUES (1, 1, 'A'),
       (2, 2, 'B'),
       (3, 3, 'A'),
       (4, 4, 'B'),
       (5, 5, 'C'),
       (6, 6, 'A'),
       (7, 7, 'B'),
       (8, 8, 'C'),
       (9, 9, 'A'),
       (10, 10, 'B');
