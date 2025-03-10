DROP TABLE IF EXISTS mark CASCADE;
DROP TABLE IF EXISTS students CASCADE;
DROP TABLE IF EXISTS lessons CASCADE;
DROP TABLE IF EXISTS teachers CASCADE;
DROP TABLE IF EXISTS groups CASCADE;
DROP TABLE IF EXISTS mark_to_student CASCADE;

CREATE TABLE IF NOT EXISTS groups
(
    id         SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS students
(
    id           SERIAL PRIMARY KEY,
    student_date DATE,
    group_id     INTEGER      NOT NULL,
    student_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (id)
);

CREATE TABLE IF NOT EXISTS teachers
(
    id           SERIAL PRIMARY KEY,
    teacher_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS lessons
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR(255) NOT NULL,
    lesson_number INTEGER,
    lesson_day    INTEGER,
    place         INTEGER,
    group_id      INTEGER,
    teacher_id    INTEGER,
    FOREIGN KEY (group_id) REFERENCES groups (id),
    FOREIGN KEY (teacher_id) REFERENCES teachers (id)
);

CREATE TABLE IF NOT EXISTS mark
(
    id         SERIAL PRIMARY KEY,
    mark_value INTEGER,
    student_id INTEGER      NOT NULL,
    discipline VARCHAR(255) NOT NULL,
    mark_date  TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students (id)
);

CREATE TABLE IF NOT EXISTS mark_to_student
(
    id         SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL,
    mark_id    INTEGER NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students (id),
    FOREIGN KEY (mark_id) REFERENCES mark (id)
);