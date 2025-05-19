CREATE TABLE role ( id SERIAL PRIMARY KEY, name VARCHAR(80) );

INSERT INTO role (name) VALUES ('admin'), ('user');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    fullname VARCHAR(60),
    phone_number VARCHAR(13),
    role_id INT REFERENCES role (id),
    image VARCHAR(255),
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
)

DROP TABLE users CASCADE;

CREATE TABLE reserve (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(80),
    phone_number VARCHAR(15),
    age VARCHAR(3),
    date DATE,
    doctor VARCHAR(60),
    complaint VARCHAR(255),
    user_id INT REFERENCES users (id),
    status VARCHAR(20)
);

SELECT * FROM users;

select * from role;

SELECT * FROM doctor;

DROP Table reserve;

select * from reserve;