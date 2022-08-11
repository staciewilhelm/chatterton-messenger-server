CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    first_name VARCHAR (100) NOT NULL,
    last_name VARCHAR (100) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    username VARCHAR (50) UNIQUE NOT NULL
);

INSERT INTO users (first_name, last_name, email, username) VALUES ('Caroline', 'Brewster', 'caro@brewster.com', 'brewster');
INSERT INTO users (first_name, last_name, email, username) VALUES ('Barb', 'Klein', 'bklein@kleinllc.com', 'bklein');
INSERT INTO users (first_name, last_name, email, username) VALUES ('Mark', 'Culper', 'mark@culper.com', 'culpertino');
