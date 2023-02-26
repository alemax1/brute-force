CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name varchar(50),
    username varchar(50) UNIQUE NOT NULL,
    password varchar(100) NOT NULL
);

DROP TABLE IF EXISTS users;