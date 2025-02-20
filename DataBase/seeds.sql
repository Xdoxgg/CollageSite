-- Вставка данных в таблицу groups
INSERT INTO groups (group_name) VALUES
('Group 1'),
('Group 2'),
('Group 3'),
('Group 4');

-- Вставка данных в таблицу students
INSERT INTO students (group_id, student_name) VALUES
(1, 'Student 1.1'),
(1, 'Student 1.2'),
(1, 'Student 1.3'),
(1, 'Student 1.4'),
(1, 'Student 1.5'),
(2, 'Student 2.1'),
(2, 'Student 2.2'),
(2, 'Student 2.3'),
(2, 'Student 2.4'),
(2, 'Student 2.5'),
(3, 'Student 3.1'),
(3, 'Student 3.2'),
(3, 'Student 3.3'),
(3, 'Student 3.4'),
(3, 'Student 3.5'),
(4, 'Student 4.1'),
(4, 'Student 4.2'),
(4, 'Student 4.3'),
(4, 'Student 4.4'),
(4, 'Student 4.5');

-- Вставка данных в таблицу lessons
DO $$
DECLARE
    group_id INTEGER;
    lesson_date DATE;
    lesson_day INTEGER;
    lesson_number INTEGER;
    lesson_title TEXT;
    place INTEGER;
    place_counter INTEGER := 1;
BEGIN
    -- Перебираем группы
    FOR group_id IN (SELECT id FROM groups) LOOP
        lesson_date := '2025-02-17'; -- Начальная дата
        lesson_day := 1; -- День недели
        -- Для каждой группы создаем уроки на 6 дней недели
        WHILE lesson_day <= 6 LOOP
            -- Для каждого дня создаем 6 уроков
            FOR lesson_number IN 1..6 LOOP
                lesson_title := 'Lesson ' || lesson_number || ' for Group ' || group_id;
                place := place_counter; -- Уникальный номер аудитории
                INSERT INTO lessons (title, lesson_number, lesson_day, place, group_id, lesson_date)
                VALUES (lesson_title, lesson_number, lesson_day, place, group_id, lesson_date);
                place_counter := place_counter + 1; -- Увеличиваем счетчик аудиторий
            END LOOP;
            lesson_day := lesson_day + 1; -- Переход к следующему дню
            lesson_date := lesson_date + INTERVAL '1 day'; -- Увеличиваем дату
        END LOOP;
    END LOOP;
END $$;