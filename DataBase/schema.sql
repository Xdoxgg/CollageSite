

CREATE TABLE IF NOT EXISTS lessons(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    lesson_number INTEGER,
    lesson_date VARCHAR(255) NOT NULL
);