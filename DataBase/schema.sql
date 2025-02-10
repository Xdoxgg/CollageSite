


CREATE TABLE IF NOT EXISTS students(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,

); 


CREATE TABLE IF NOT EXISTS group(
    id SERIAL PRIMARY KEY
    student_id INTEGER NOT NULL,
    group_numbrt VARCHAR(255) NOT NULL,
     
    FOREIGN KEY 

);
