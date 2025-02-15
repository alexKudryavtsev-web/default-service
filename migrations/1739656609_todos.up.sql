CREATE TABLE
    IF NOT EXISTS todos (
        id serial PRIMARY KEY,
        task VARCHAR(255)
    );