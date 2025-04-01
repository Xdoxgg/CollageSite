-- Сиды для таблицы groups
INSERT INTO groups (group_name)
VALUES ('Group A'),
       ('Group B'),
       ('Group C');

-- Сиды для таблицы teachers
INSERT INTO teachers (teacher_name)
VALUES ('John Smith'),
       ('Jane Doe'),
       ('Emily Johnson');

-- Сиды для таблицы students
INSERT INTO students (group_id, student_name)
VALUES (1, 'Alice Brown'),
       (1, 'Bob White'),
       (2, 'Charlie Green'),
       (2, 'Diana Black'),
       (3, 'Eve Blue'),
       (1, 'Крутой тип');  -- Убедитесь, что student_date не требуется

-- Сиды для таблицы lessons
INSERT INTO lessons (title, lesson_number, lesson_day, place, group_id, teacher_id)
VALUES
-- День 1
('Mathematics', 1, 1, 101, 1, 1),
('Physics', 2, 1, 102, 1, 2),
('Chemistry', 3, 1, 103, 1, 3),
('Biology', 4, 1, 104, 1, 1),
('History', 5, 1, 105, 1, 2);

-- Сиды для таблицы mark
INSERT INTO mark (mark_value, student_id, discipline, mark_date)
VALUES
    (5, 1, 'Mathematics', '2025-03-10'),
    (9, 1, 'Physics', '2025-03-10'),
    (8, 1, 'Chemistry', '2025-03-10'),
    (8, 1, 'Biology', '2025-03-10'),
    (2, 1, 'History', '2025-04-10');

-- Убедитесь, что student_id в mark_to_student существует
INSERT INTO mark_to_student (student_id, mark_id)
VALUES (1, 1),  -- Замените на существующий student_id
       (1, 2),
       (1, 3),
       (1, 4);

-- Сиды для таблицы news
INSERT INTO news (title, data, img, post_date)
VALUES ('test news1', 'testdata', 'img1', '2025-03-22'),
       ('test new2', 'testdata', 'img2', '2025-03-21');