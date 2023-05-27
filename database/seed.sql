# My data examples for the assignment
# Use at your own risk
INSERT INTO
  professor (prof_id, prof_lname, prof_fname)
VALUES
  (1, 'Smith', 'John'),
  (2, 'Doe', 'Jane'),
  (3, 'Johnson', 'Michael'),
  (4, 'Williams', 'Emily'),
  (5, 'Jones', 'David'),
  (6, 'Garcia', 'Sophia');

INSERT INTO
  student (
    stud_id,
    stud_lname,
    stud_fname,
    stud_street,
    stud_city,
    stud_zip
  )
VALUES
  (
    1,
    'Brown',
    'Alice',
    '123 Main St',
    'Springfield',
    '12345'
  ),
  (
    2,
    'Davis',
    'Bob',
    '456 Elm St',
    'Shelbyville',
    '54321'
  ),
  (
    3,
    'Miller',
    'Charlie',
    '789 Oak St',
    'Capital City',
    '67890'
  ),
  (
    4,
    'Rodriguez',
    'Daniel',
    '246 Maple St',
    'Ogdenville',
    '24680'
  ),
  (
    5,
    'Wilson',
    'Emma',
    '369 Pine St',
    'North Haverbrook',
    '36912'
  ),
  (
    6,
    'Martinez',
    'Olivia',
    '159 Cedar St',
    'Brockway',
    '15975'
  );

INSERT INTO
  class (class_id, class_name, prof_id, course_id)
VALUES
  (1, 'Intro to Computer Science', 1, 1),
  (2, 'Calculus I', 2, 2),
  (3, 'English Literature', 3, 3),
  (4, 'Chemistry I', 4, 4);

INSERT INTO
  room (room_id, room_loc, room_cap)
VALUES
  (101, 'Building A', '30'),
  (102, 'Building B', '40'),
  (103, 'Building C', '35'),
  (104, 'Building D', '45');

UPDATE
  room
set
  class_id = 1
where
  room_id = 101;

UPDATE
  room
set
  class_id = 2
where
  room_id = 102;

UPDATE
  room
set
  class_id = 3
where
  room_id = 103;

UPDATE
  room
set
  class_id = 4
where
  room_id = 104;

INSERT INTO
  course (course_id, course_name)
VALUES
  (1, 'Computer Science'),
  (2, 'Mathematics'),
  (3, 'English'),
  (4, 'Chemistry');

INSERT INTO
  enroll (stud_id, class_id, grade)
VALUES
  (1, 1, 'A'),
  (1, 2, 'B'),
  (2, 1, 'C'),
  (3, 2, 'A'),
  (4, 1, 'B'),
  (4, 2, 'A'),
  (6, 2, 'C'),
  (1, 3, 'A'),
  (1, 4, 'B'),
  (2, 3, 'C'),
  (2, 4, 'A'),
  (5, 1, 'C'),
  (5, 2, 'A'),
  (5, 6, 'A');

INSERT INTO
  room (room_id, room_loc, room_cap)
VALUES
  (105, 'Building E', '50'),
  (106, 'Building F', '55');

INSERT INTO
  class (
    class_id,
    class_name,
    prof_id,
    course_id,
    room_id
  )
VALUES
  (5, 'History of Art', 5, 5, 105),
  (6, 'Physics I', 6, 6, 106);