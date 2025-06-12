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
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
)

INSERT INTO
    users (
        email,
        password,
        fullname,
        phone_number,
        role_id,
        image
    )
VALUES (
        'admin123@gmail.com',
        '$argon2i$v=19$m=16,t=2,p=1$VlZOb1gySGlXRkZJVktvSQ$eDQU7tuYb+pHMwJSVINfPw', --Admin123
        'Admin User',
        '1234567890',
        1,
        'https://example.com/image.jpg'
    )

DELETE FROM users WHERE id = 20;

select * FROM users;

DROP Table users CASCADE;

CREATE TABLE reserve (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(80),
    phone_number VARCHAR(15),
    age VARCHAR(3),
    date DATE,
    doctor VARCHAR(60),
    complaint VARCHAR(255),
    user_id INT REFERENCES users (id),
    status VARCHAR(20),
    rec_medic VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);