CREATE TABLE role ( id SERIAL PRIMARY KEY, name VARCHAR(80) );

INSERT INTO role (name) VALUES ('admin'), ('user');

SELECT * FROM users;

DROP TABLE reserve;

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

DROP TABLE users CASCADE;

select * from reserve;

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

INSERT INTO
    reserve (
        fullname,
        phone_number,
        age,
        date,
        doctor,
        complaint,
        user_id,
        status,
        rec_medic,
        updated_at
    )
VALUES (
        'Andi Pratama',
        '081234567890',
        '30',
        '2025-05-25',
        'Dr. Siti Aulia',
        'Demam tinggi dan batuk',
        1,
        'pending',
        'Paracetamol 500mg, istirahat total',
        '2025-05-25 10:00:00'
    ),
    (
        'Budi Santoso',
        '081298765432',
        '45',
        '2025-05-26',
        'Dr. Rudi Hartono',
        'Sakit perut dan mual',
        1,
        'confirmed',
        'Antasida, hindari makanan pedas',
        '2025-05-26 09:30:00'
    ),
    (
        'Citra Lestari',
        '082112345678',
        '28',
        '2025-05-27',
        'Dr. Nina Kurnia',
        'Pusing dan kelelahan',
        1,
        'completed',
        'Multivitamin, olahraga ringan',
        '2025-05-27 14:15:00'
    ),
    (
        'Dedi Haryanto',
        '081377788899',
        '52',
        '2025-05-28',
        'Dr. Siti Aulia',
        'Nyeri sendi lutut',
        1,
        'pending',
        'Obat antiinflamasi, kompres hangat',
        '2025-05-28 08:45:00'
    ),
    (
        'Eka Yulianti',
        '085678901234',
        '35',
        '2025-05-29',
        'Dr. Rudi Hartono',
        'Sesak napas ringan',
        1,
        'cancelled',
        'Periksa ulang paru-paru, hindari asap',
        '2025-05-29 11:20:00'
    ),
    (
        'Fajar Nugroho',
        '082233344455',
        '40',
        '2025-05-30',
        'Dr. Nina Kurnia',
        'Tensi tinggi',
        1,
        'confirmed',
        'Amlodipine 5mg tiap pagi',
        '2025-05-30 10:10:00'
    ),
    (
        'Gita Prameswari',
        '081355566677',
        '22',
        '2025-06-01',
        'Dr. Siti Aulia',
        'Sakit tenggorokan dan pilek',
        1,
        'pending',
        'Antibiotik ringan dan air hangat',
        '2025-06-01 09:50:00'
    ),
    (
        'Hendra Wijaya',
        '081234567123',
        '50',
        '2025-06-02',
        'Dr. Rudi Hartono',
        'Kolesterol tinggi',
        1,
        'completed',
        'Simvastatin 20mg malam hari',
        '2025-06-02 13:05:00'
    ),
    (
        'Intan Permata',
        '085611223344',
        '29',
        '2025-06-03',
        'Dr. Nina Kurnia',
        'Migrain berkala',
        1,
        'confirmed',
        'Ibuprofen jika nyeri, hindari stres',
        '2025-06-03 15:30:00'
    ),
    (
        'Joko Susilo',
        '082155566677',
        '38',
        '2025-06-04',
        'Dr. Siti Aulia',
        'Sakit pinggang',
        1,
        'pending',
        'Fisioterapi ringan, pereda nyeri',
        '2025-06-04 08:40:00'
    );

SELECT * FROM users;

select * from role;

SELECT * FROM doctor;

DROP Table reserve;

select * from reserve;