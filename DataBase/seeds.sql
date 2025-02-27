-- Сиды для таблицы groups
INSERT INTO groups (group_name) VALUES 
('Group A'),
('Group B'),
('Group C');

-- Сиды для таблицы teachers
INSERT INTO teachers (teacher_name) VALUES 
('John Smith'),
('Jane Doe'),
('Emily Johnson');

-- Сиды для таблицы students
INSERT INTO students (group_id, student_name) VALUES 
(1, 'Alice Brown'),
(1, 'Bob White'),
(2, 'Charlie Green'),
(2, 'Diana Black'),
(3, 'Eve Blue');


-- Чёткий парень

INSERT INTO students (group_id, student_name, student_date) VALUES 
(1, 'Крутой тип', '2007-05-29');


-- Сиды для таблицы lessons
INSERT INTO lessons (title, lesson_number, lesson_day, place, group_id, teacher_id) VALUES 
-- День 1
('Mathematics', 1, 1, 101, 1, 1),
('Physics', 2, 1, 102, 1, 2),
('Chemistry', 3, 1, 103, 1, 3),
('Biology', 4, 1, 104, 1, 1),
('History', 5, 1, 105, 1, 2),
('History', 6, 1, 105, 1, 2),
('History', 7, 1, 105, 1, 2),
('History', 8, 1, 105, 1, 2),

-- День 2
('Mathematics', 1, 2, 101, 1, 1),
('Physics', 2, 2, 102, 1, 2),
('Chemistry', 3, 2, 103, 1, 3),
('History', 5, 2, 105, 1, 2),
('History', 6, 2, 105, 1, 2),
('History', 7, 2, 105, 1, 2),
('History', 8, 2, 105, 1, 2),

-- День 3
('Mathematics', 1, 3, 101, 1, 1),
('Physics', 2, 3, 102, 1, 2),
('Chemistry', 3, 3, 103, 1, 3),
('Biology', 4, 3, 104, 1, 1),
('History', 5, 3, 105, 1, 2),
('History', 6, 3, 105, 1, 2),
('History', 7, 3, 105, 1, 2),
('History', 8, 3, 105, 1, 2),

-- День 4
('Mathematics', 1, 4, 101, 1, 1),
('Physics', 2, 4, 102, 1, 2),
('Chemistry', 3, 4, 103, 1, 3),
('Biology', 4, 4, 104, 1, 1),
('History', 5, 4, 105, 1, 2),


-- День 5

('Biology', 4, 5, 104, 1, 1),
('History', 5, 5, 105, 1, 2),
('History', 6, 5, 105, 1, 2),
('History', 7, 5, 105, 1, 2),
('History', 8, 5, 105, 1, 2),

-- День 6
('Mathematics', 1, 6, 101, 1, 1),
('Physics', 2, 6, 102, 1, 2),
('Chemistry', 3, 6, 103, 1, 3),
('Biology', 4, 6, 104, 1, 1),
('History', 5, 6, 105, 1, 2),

('History', 8, 6, 105, 1, 2);

-- Сиды для таблицы mark
INSERT INTO mark (mark_value, student_id, discipline) VALUES 
(85, 1, 'Mathematics'),
(90, 2, 'Physics'),
(78, 3, 'Chemistry'),
(88, 4, 'Biology'),
(92, 5, 'History');