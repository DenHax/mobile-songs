CREATE TABLE Song (
    id SERIAL PRIMARY KEY,
    song VARCHAR(255) NOT NULL,
    group_name VARCHAR(255) NOT NULL,
    lyrics TEXT
);
