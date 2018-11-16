CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fname TEXT,
    lname TEXT,
    email TEXT NOT NULL
);

CREATE TABLE alarms (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    occur TIMESTAMP NOT NULL
);