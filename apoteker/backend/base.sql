
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    fullname VARCHAR(60),
    phone_number VARCHAR(13),
    role_id INT REFERENCES role(id),
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE doctor(
    id SERIAL PRIMARY KEY,
    name VARCHAR(80),
    spesialis VARCHAR(80),
    layanan int,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO doctor (name, spesialis) VALUES
('dr. Andi', 'Bidan'),
('dr. Budi', 'Dokter Anak'),
('dr. Citra', 'Trapis'),
('dr. Dedi', 'Anak');
CREATE TABLE role(
    id SERIAL PRIMARY KEY,
    name VARCHAR(80)
);

INSERT INTO role (name) VALUES
('admin'),
('user');

DROP TABLE users CASCADE;

CREATE TABLE reserve(
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(80),
    phone_number VARCHAR(15),
    age VARCHAR(3),
    date DATE,
    complaint VARCHAR(255),
    user_id INT REFERENCES users(id),
    doctor_id INT REFERENCES doctor(id)
);

SELECT * FROM users;

select * from role;

select * from reserve;