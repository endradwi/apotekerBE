DROP table cinema_date, cinema_location, cinema_time, cinemas, movies;


SELECT profile.first_name, orders.id, cinemas.name , movies.tittle 
from profile 
full OUTER JOIN orders ON orders.profile_id = profile.id
full OUTER JOIN cinemas ON cinemas.id = orders.cinema_id
full OUTER JOIN movies ON movies.id = orders.movie_id
;

drop table cinemas;

INSERT INTO movie (tittle, genre, images, synopsis,
	author, actors, release_date, duration, tag) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, tittle, genre, images, synopsis,
	author, actors, release_date, duration, tag;


CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE profile(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(60),
    last_name VARCHAR(60),
    image VARCHAR(255),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

ALTER TABLE profile ADD phone_number int;
ALTER TABLE profile ADD point int;
SELECT * from profile;

drop table profile;

DROP TABLE orders;
CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    profile_id INT REFERENCES profile(id),
    movie_id INT REFERENCES movies(id),
    cinema_id INT REFERENCES cinema(id),
    seat_id INT REFERENCES seat(id),
    date DATE,
    qty INT,
    total_price INT,
    payment_id int REFERENCES payment_method(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

drop table ORDERS;
CREATE TABLE movies(
    id SERIAL PRIMARY KEY,
    tittle VARCHAR(60),
    genre VARCHAR(100),
    images VARCHAR(255),
    synopsis VARCHAR(255),
    author VARCHAR(60),
    actors VARCHAR(255),
    release_date DATE,
    duration TIME,
    tag VARCHAR(30),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);
CREATE Table cinema(
    id serial PRIMARY KEY,
    name VARCHAR(50),
    price INT,
    movies_id INT REFERENCES movies(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
)

INSERT INTO orders(profile_id, movie_id, cinema_id, total_price, qty) VALUES (154, 1, 1, 10000, 2);
INSERT INTO cinema (name, price, movies_id) VALUES
('XXI', 20000, 1),
('Blitzmegaplex', 15000, 5),
('CGV Cinemas Indonesia', 30000, 3),
('Cineplex 21 ', 10000, 2),
('Cinépolis Indonesia', 18000, 4),
('Dakota Cinema', 25000, 6),
('KOTA Cinema Mall', 10000, 10),
('New Star Cineplex', 17000, 9),
('Platinum Cineplex', 22000, 7),
('The Premiere', 100000, 8);

SELECT price FROM cinema WHERE id = $1;
drop Table cinema;

CREATE TABLE cinema_date(
    id SERIAL PRIMARY KEY,
    name_date DATE,
    cinema_id int REFERENCES cinema(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
)


CREATE TABLE cinema_time(
    id SERIAL PRIMARY KEY,
    name_time TIME,
    cinema_id int REFERENCES cinema(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP 
)

CREATE TABLE cinema_location(
    id SERIAL PRIMARY KEY,
    name_location VARCHAR(50),
    cinema_id int REFERENCES cinema(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO cinema_location (name_location, cinema_id) VALUES
('Bekasi', 1), ('Jakarta', 2), ('Depok', 3), ('Bogor', 4),
('Karawang', 5), ('Subang', 6), ('Semarang', 7), ('Jogja', 8),
('Solo', 9), ('Surabaya', 10);
drop Table cinema_location, cinema_date, cinema_date;
drop TABLE cinema;

SELECT profile.id, orders.profile_id, orders.movie_id, movies.tittle,
	movies.genre, movies.images, cinema.name, 
	FORM profile
	JOIN orders ON orders.profile_id = profile.id
	-- JOIN movies ON movies.id = orders.movie_id
	-- JOIN cinema ON cinema.id = ordes_cinema_id
	WHERE profile.id = 1;

SELECT profile.first_name, orders.profile_id, orders.movie_id,
movies.tittle, movies.genre, movies.images, cinema.name
from profile
JOIN orders ON orders.profile_id = profile.id
JOIN movies ON movies.id = orders.movie_id
JOIN cinema ON cinema.id = orders.cinema_id
WHERE profile.id = 1
;

CREATE TABLE seat (
    id serial PRIMARY KEY,
    name varchar(3),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);
DROP table seat;


SELECT count(seat_id) as seat_id  From orders 
WHERE id = $1;

SELECT * from seat;
TRUNCATE TABLE seat;
INSERT INTO seat (name) VALUES
('A1'),('A2'),('A3'),('A4'),('A5'),
('A6'),('A7'),('A8'),('A9'),('A10'),
('A11'),('A12'),('A13'),('A14'),('B1'),
('B2'),('B3'),('B4'),('B5'),('B6'),
('B7'),('B8'),('B9'),('B10'),('B11'),
('B12'),('B13'),('B14'),('C1'),('C2'),
('C3'),('C4'),('C5'),('C6'),('C7'),
('C8'),('C9'),('C10'),('C11'),('C12'),
('C13'),('C14'),('D1'),('D2'),('D3'),
('D4'),('D5'),('D6'),('D7'),('D8'),
('D9'),('D10'),('D11'),('D12'),('D13'),
('D14'),('E1'),('E2'),('E3'),('E4'),
('E5'),('E6'),('E7'),('E8'),('E9'),
('E10'),('E11'),('E12'),('E13'),('E14'),
('F1'),('F2'),('F3'),('F4'),('F5'),
('F6'),('F7'),('F8'),('F9'),('F10'),
('F11'),('F12'),('F13'),('F14'),('G1'),
('G2'),('G3'),('G4'),('G5'),('G6'),
('G7'),('G8'),('G9'),('G10'),('G10'),
('G11'),('G12'),('G13'),('G14');

CREATE TABLE payment_method (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO payment_method (name) VALUES
('Goggle Pay'), ('Visa'), ('Gopay'), ('Pay Pal'),
('Dana'), ('Bank BCA'), ('Bank BRI'), ('OVO');
UPDATE users
	SET email=$1, password=$2
	FROM profile
	WHERE users.id=$3
	AND users.id = profile.user_id
	RETURNING users.email, users.password;

SELECT profile.id, orders.id, movies.tittle, movies.images, movies.genre,
cinema.name, cinema_location.name_location FROM orders
JOIN profile ON orders.profile_id = profile.id 
JOIN movies ON movies.id = orders.movie_id 
JOIN cinema ON cinema.id = orders.cinema_id
JOIN cinema_location ON cinema_location.id = cinema.id
WHERE orders.id = $1;

SELECT  movies.tittle,movies.images,movies.genre, cinema.name, cinema_location.name_location from orders
-- JOIN profile ON profile.id = orders.profile_id
JOIN movies ON movies.id = orders.movie_id
JOIN cinema ON cinema.id = orders.cinema_id
JOIN cinema_location ON cinema_location.cinema_id = cinema.id
WHERE orders.id = 1;

SELECT  movies.tittle, movies.images, movies.genre,
cinema.name, cinema_location.name_location FROM orders
-- JOIN orders ON orders.profile_id = profile.id 
JOIN movies ON movies.id = orders.movie_id 
JOIN cinema ON cinema.id = orders.cinema_id
JOIN cinema_location ON cinema_location.id = cinema.id
WHERE orders.id = 100;

SELECT * from profile;
SELECT * from orders;

SELECT * FROM ;

SELECT users.id, users.email, users.password, profile.image FROM users 
JOIN profile ON profile.user_id = users.id
ORDER BY users.id ASC;

UPDATE users
SET email = COALESCE($1, email),
    password = COALESCE($2, password)
WHERE email = $3;

SELECT * from users;

drop table orders;
drop TABLE seat;

INSERT INTO users (email, password) VALUES
('john.doe@example.com', 'password123'),
('jane.smith@example.com', 'password123'),
('michael.brown@example.com', 'password123'),
('emily.jones@example.com', 'password123'),
('chris.davis@example.com', 'password123'),
('sarah.miller@example.com', 'password123'),
('david.moore@example.com', 'password123'),
('mary.taylor@example.com', 'password123'),
('robert.anderson@example.com', 'password123'),
('lisa.thomas@example.com', 'password123'),
('james.jackson@example.com', 'password123'),
('patricia.white@example.com', 'password123'),
('charles.harris@example.com', 'password123'),
('joseph.martin@example.com', 'password123'),
('jennifer.lee@example.com', 'password123'),
('mark.young@example.com', 'password123'),
('nancy.king@example.com', 'password123'),
('kevin.green@example.com', 'password123'),
('daniel.adams@example.com', 'password123'),
('susan.carter@example.com', 'password123'),
('william.hernandez@example.com', 'password12'),
('barbara.martinez@example.com', 'password123'),
('david.garcia@example.com', 'password123'),
('linda.rodriguez@example.com', 'password123'),
('george.wilson@example.com', 'password123'),
('elizabeth.moore@example.com', 'password123'),
('joseph.taylor@example.com', 'password123'),
('martha.jackson@example.com', 'password123'),
('nicholas.johnson@example.com', 'password123'),
('christopher.brown@example.com', 'password123'),
('jacqueline.martin@example.com', 'password123'),
('george.clark@example.com', 'password123'),
('melissa.wright@example.com', 'password123'),
('susan.garcia@example.com', 'password123'),
('daniel.kim@example.com', 'password123'),
('julia.martinez@example.com', 'password123'),
('aaron.williams@example.com', 'password123'),
('heather.perez@example.com', 'password123'),
('roberto.smith@example.com', 'password123'),
('laura.johnson@example.com', 'password123'),
('joseph.smith@example.com', 'password123'),
('anna.white@example.com', 'password123'),
('brian.miller@example.com', 'password123'),
('lisa.davis@example.com', 'password123'),
('john.morris@example.com', 'password123'),
('susan.kelly@example.com', 'password123'),
('stephen.smith@example.com', 'password123'),
('lisa.campbell@example.com', 'password123'),
('jackson.evans@example.com', 'password123'),
('olivia.jenkins@example.com', 'password123'),
('elizabeth.sullivan@example.com', 'password10'),
('benjamin.wood@example.com', 'password123'),
('rachel.lee@example.com', 'password123'),
('ethan.owen@example.com', 'password123'),
('zoe.adams@example.com', 'password123'),
('rachel.morris@example.com', 'password123'),
('louis.james@example.com', 'password123'),
('michelle.harris@example.com', 'password123'),
('nathan.wilson@example.com', 'password123'),
('laura.jameson@example.com', 'password123'),
('victor.carter@example.com', 'password123'),
('cheryl.roberts@example.com', 'password123'),
('olivia.morris@example.com', 'password123'),
('grace.hernandez@example.com', 'password123'),
('henry.king@example.com', 'password123'),
('elizabeth.carter@example.com', 'password123'),
('joseph.brown@example.com', 'password123');

select * from profile;
INSERT INTO profile (first_name, last_name, image, user_id) VALUES
('John', 'Doe', 'https://example.com/images/john_doe.jpg', 1);

INSERT INTO profile (first_name, last_name, image, user_id) VALUES
('John', 'Doe', 'https://example.com/images/john_doe.jpg', 1),
('Jane', 'Smith', 'https://example.com/images/jane_smith.jpg', 2),
('Michael', 'Brown', 'https://example.com/images/michael_brown.jpg', 3),
('Emily', 'Jones', 'https://example.com/images/emily_jones.jpg', 4),
('Chris', 'Davis', 'https://example.com/images/chris_davis.jpg', 5),
('Sarah', 'Miller', 'https://example.com/images/sarah_miller.jpg', 6),
('David', 'Moore', 'https://example.com/images/david_moore.jpg', 7),
('Mary', 'Taylor', 'https://example.com/images/mary_taylor.jpg', 8),
('Robert', 'Anderson', 'https://example.com/images/robert_anderson.jpg', 9),
('Lisa', 'Thomas', 'https://example.com/images/lisa_thomas.jpg', 10),
('James', 'Jackson', 'https://example.com/images/james_jackson.jpg', 11),
('Patricia', 'White', 'https://example.com/images/patricia_white.jpg', 12),
('Charles', 'Harris', 'https://example.com/images/charles_harris.jpg', 13),
('Joseph', 'Martin', 'https://example.com/images/joseph_martin.jpg', 14),
('Jennifer', 'Lee', 'https://example.com/images/jennifer_lee.jpg', 15),
('Mark', 'Young', 'https://example.com/images/mark_young.jpg', 16),
('Nancy', 'King', 'https://example.com/images/nancy_king.jpg', 17),
('Kevin', 'Green', 'https://example.com/images/kevin_green.jpg', 18),
('Daniel', 'Adams', 'https://example.com/images/daniel_adams.jpg', 19),
('Susan', 'Carter', 'https://example.com/images/susan_carter.jpg', 20),
('William', 'Hernandez', 'https://example.com/images/william_hernandez.jpg', 21),
('Barbara', 'Martinez', 'https://example.com/images/barbara_martinez.jpg', 22),
('David', 'Garcia', 'https://example.com/images/david_garcia.jpg', 23),
('Linda', 'Rodriguez', 'https://example.com/images/linda_rodriguez.jpg', 24),
('George', 'Wilson', 'https://example.com/images/george_wilson.jpg', 25),
('Elizabeth', 'Moore', 'https://example.com/images/elizabeth_moore.jpg', 26),
('Joseph', 'Taylor', 'https://example.com/images/joseph_taylor.jpg', 27),
('Martha', 'Jackson', 'https://example.com/images/martha_jackson.jpg', 28),
('Nicholas', 'Johnson', 'https://example.com/images/nicholas_johnson.jpg', 29),
('Christopher', 'Brown', 'https://example.com/images/christopher_brown.jpg', 30),
('Jacqueline', 'Martin', 'https://example.com/images/jacqueline_martin.jpg', 31),
('George', 'Clark', 'https://example.com/images/george_clark.jpg', 32),
('Melissa', 'Wright', 'https://example.com/images/melissa_wright.jpg', 33),
('Susan', 'Garcia', 'https://example.com/images/susan_garcia.jpg', 34),
('Daniel', 'Kim', 'https://example.com/images/daniel_kim.jpg', 35),
('Julia', 'Martinez', 'https://example.com/images/julia_martinez.jpg', 36),
('Aaron', 'Williams', 'https://example.com/images/aaron_williams.jpg', 37),
('Heather', 'Perez', 'https://example.com/images/heather_perez.jpg', 38),
('Roberto', 'Smith', 'https://example.com/images/roberto_smith.jpg', 39),
('Laura', 'Johnson', 'https://example.com/images/laura_johnson.jpg', 40),
('Joseph', 'Smith', 'https://example.com/images/joseph_smith.jpg', 41),
('Anna', 'White', 'https://example.com/images/anna_white.jpg', 42),
('Brian', 'Miller', 'https://example.com/images/brian_miller.jpg', 43),
('Lisa', 'Davis', 'https://example.com/images/lisa_davis.jpg', 44),
('John', 'Morris', 'https://example.com/images/john_morris.jpg', 45),
('Susan', 'Kelly', 'https://example.com/images/susan_kelly.jpg', 46),
('Stephen', 'Smith', 'https://example.com/images/stephen_smith.jpg', 47),
('Lisa', 'Campbell', 'https://example.com/images/lisa_campbell.jpg', 48),
('Jackson', 'Evans', 'https://example.com/images/jackson_evans.jpg', 49),
('Olivia', 'Jenkins', 'https://example.com/images/olivia_jenkins.jpg', 50),
('William', 'Reed', 'https://example.com/images/william_reed.jpg', 51),
('Charlotte', 'Baker', 'https://example.com/images/charlotte_baker.jpg', 52),
('George', 'King', 'https://example.com/images/george_king.jpg', 53),
('Isabella', 'Scott', 'https://example.com/images/isabella_scott.jpg', 54),
('Samuel', 'Adams', 'https://example.com/images/samuel_adams.jpg', 55),
('Lily', 'Martinez', 'https://example.com/images/lily_martinez.jpg', 56),
('Matthew', 'Lee', 'https://example.com/images/matthew_lee.jpg', 57),
('Sophia', 'Allen', 'https://example.com/images/sophia_allen.jpg', 58),
('Henry', 'Hall', 'https://example.com/images/henry_hall.jpg', 59),
('Grace', 'Bennett', 'https://example.com/images/grace_bennett.jpg', 60);

INSERT INTO movies (tittle, genre, images, synopsis, author, actors, release_date, duration, tag) VALUES
('Avatar', 'Action, Adventure, Science Fiction', 'avatar.jpg', 'Di dunia Pandora, manusia berusaha mengeksploitasi sumber daya alam, sementara kelompok Navi berjuang untuk melindungi rumah mereka.', 'James Cameron', 'Sam Worthington, Zoe Saldana', '2009-12-18', '02:42:00', 'epic'),
('The Dark Knight', 'Action, Crime, Drama', 'dark_knight.jpg', 'Batman berusaha menghentikan Joker yang berencana menciptakan kekacauan di Gotham City.', 'Christopher Nolan', 'Christian Bale, Heath Ledger', '2008-07-18', '02:32:00', 'thriller'),
('Inception', 'Action, Adventure, Science Fiction', 'inception.jpg', 'Seorang pencuri yang dapat memasuki mimpi orang lain diupah untuk menanamkan ide ke pikiran seseorang.', 'Christopher Nolan', 'Leonardo DiCaprio, Joseph Gordon-Levitt', '2010-07-16', '02:28:00', 'mind-bending'),
('Titanic', 'Drama, Romance', 'titanic.jpg', 'Kisah cinta antara Jack dan Rose yang terjebak dalam tragedi tenggelamnya kapal Titanic.', 'James Cameron', 'Leonardo DiCaprio, Kate Winslet', '1997-12-19', '03:14:00', 'classic'),
('The Matrix', 'Action, Science Fiction', 'matrix.jpg', 'Seorang programmer mengetahui bahwa dunia tempat dia tinggal adalah simulasi buatan yang diciptakan oleh mesin.', 'Lana Wachowski, Lilly Wachowski', 'Keanu Reeves, Laurence Fishburne', '1999-03-31', '02:16:00', 'cyberpunk'),
('Forrest Gump', 'Drama, Romance', 'forrest_gump.jpg', 'Seorang pria sederhana mengalami berbagai peristiwa besar dalam sejarah Amerika, dari Perang Vietnam hingga Watergate.', 'Eric Roth', 'Tom Hanks, Robin Wright', '1994-07-06', '02:22:00', 'inspirational'),
('Avengers: Endgame', 'Action, Adventure, Science Fiction', 'avengers_endgame.jpg', 'Para Avengers berusaha untuk mengalahkan Thanos setelah kehancuran yang dia bawa di Avengers: Infinity War.', 'Anthony Russo, Joe Russo', 'Robert Downey Jr., Chris Evans', '2019-04-26', '03:02:00', 'superhero'),
('The Shawshank Redemption', 'Drama', 'shawshank.jpg', 'Seorang pria yang dihukum untuk kejahatan yang tidak dia lakukan berteman dengan sesama tahanan dan merencanakan pelarian.', 'Frank Darabont', 'Tim Robbins, Morgan Freeman', '1994-09-22', '02:22:00', 'hope'),
('Jurassic Park', 'Action, Adventure, Science Fiction', 'jurassic_park.jpg', 'Seorang ilmuwan menciptakan taman safari yang dihuni oleh dinosaurus, tetapi segalanya menjadi kacau ketika dinosaurus lepas.', 'Michael Crichton', 'Sam Neill, Laura Dern', '1993-06-11', '02:06:00', 'thriller'),
('The Godfather', 'Crime, Drama', 'godfather.jpg', 'Kisah keluarga mafia Corleone dan perjuangan Michael Corleone untuk mempertahankan kekuatan dan kehormatan keluarga.', 'Mario Puzo', 'Marlon Brando, Al Pacino', '1972-03-24', '02:55:00', 'classic'),
('The Lion King', 'Animation, Adventure, Drama', 'lion_king.jpg', 'Seorang anak singa berjuang untuk merebut kembali tahtanya setelah ayahnya dibunuh.', 'Roger Allers, Rob Minkoff', 'Matthew Broderick, Jeremy Irons', '1994-06-24', '01:28:00', 'animated'),
('Gladiator', 'Action, Adventure, Drama', 'gladiator.jpg', 'Seorang prajurit yang dikhianati berjuang untuk membalas dendam di arena gladiator Roma.', 'Ridley Scott', 'Russell Crowe, Joaquin Phoenix', '2000-05-05', '02:35:00', 'epic'),
('Star Wars: Episode V - The Empire Strikes Back', 'Action, Adventure, Fantasy', 'empire_strikes_back.jpg', 'Rebel Alliance berjuang melawan Kekaisaran Galaksi, sementara Luke Skywalker melanjutkan latihannya dengan Yoda.', 'Irvin Kershner', 'Mark Hamill, Harrison Ford', '1980-05-21', '02:04:00', 'sci-fi'),
('Back to the Future', 'Adventure, Comedy, Science Fiction', 'back_to_the_future.jpg', 'Seorang remaja bepergian ke masa lalu menggunakan mesin waktu dan bertemu orang tuanya yang muda.', 'Robert Zemeckis', 'Michael J. Fox, Christopher Lloyd', '1985-07-03', '01:56:00', 'time-travel'),
('The Avengers', 'Action, Adventure, Science Fiction', 'avengers.jpg', 'Tim superhero berusaha untuk bekerja sama dan melawan ancaman yang mengancam dunia.', 'Joss Whedon', 'Robert Downey Jr., Chris Hemsworth', '2012-05-04', '02:23:00', 'teamwork'),
('The Lord of the Rings: The Return of the King', 'Action, Adventure, Drama', 'lotr_return_of_king.jpg', 'Frodo dan Sam berusaha untuk menghancurkan cincin kekuasaan di Gunung Doom, sementara pasukan Rohan dan Gondor bersatu melawan Sauron.', 'Peter Jackson', 'Elijah Wood, Ian McKellen', '2003-12-17', '03:21:00', 'epic'),
('Pulp Fiction', 'Crime, Drama', 'pulp_fiction.jpg', 'Kisah beberapa karakter yang terhubung melalui serangkaian kejadian kriminal di Los Angeles.', 'Quentin Tarantino', 'John Travolta, Uma Thurman', '1994-10-14', '02:34:00', 'cult'),
('Forrest Gump', 'Drama, Romance', 'forrest_gump.jpg', 'Seorang pria sederhana mengalami berbagai peristiwa besar dalam sejarah Amerika, dari Perang Vietnam hingga Watergate.', 'Eric Roth', 'Tom Hanks, Robin Wright', '1994-07-06', '02:22:00', 'inspirational'),
('Schindler List', 'Biography, Drama, History', 'schindlers_list.jpg', 'Seorang pengusaha Jerman menyelamatkan lebih dari seribu orang Yahudi selama Holocaust.', 'Steven Zaillian', 'Liam Neeson, Ben Kingsley', '1993-12-15', '03:15:00', 'historical'),
('The Departed', 'Crime, Drama, Thriller', 'departed.jpg', 'Seorang polisi menyamar sebagai gangster, sementara seorang gangster menyamar sebagai polisi dalam perjuangan untuk melindungi kota Boston.', 'William Monahan', 'Leonardo DiCaprio, Matt Damon', '2006-10-06', '02:31:00', 'thriller'),
('The Silence of the Lambs', 'Crime, Drama, Thriller', 'silence_of_the_lambs.jpg', 'Seorang agen FBI bekerja sama dengan seorang pembunuh berantai untuk menangkap pembunuh lainnya.', 'Thomas Harris', 'Jodie Foster, Anthony Hopkins', '1991-02-14', '01:58:00', 'psychological'),
('The Prestige', 'Drama, Mystery, Science Fiction', 'prestige.jpg', 'Dua pesulap terlibat dalam persaingan sengit untuk menciptakan trik sulap yang lebih hebat.', 'Jonathan Nolan', 'Christian Bale, Hugh Jackman', '2006-10-20', '02:10:00', 'mystery'),
('The Dark Knight Rises', 'Action, Drama', 'dark_knight_rises.jpg', 'Batman kembali untuk melawan Bane yang berusaha menghancurkan Gotham City.', 'Christopher Nolan', 'Christian Bale, Tom Hardy', '2012-07-20', '02:44:00', 'superhero'),
('The Green Mile', 'Crime, Drama, Fantasy', 'green_mile.jpg', 'Seorang sipir penjara menjalin hubungan dengan seorang terpidana mati yang memiliki kekuatan penyembuhan luar biasa.', 'Frank Darabont', 'Tom Hanks, Michael Clarke Duncan', '1999-12-10', '03:09:00', 'emotional'),
('Interstellar', 'Adventure, Drama, Science Fiction', 'interstellar.jpg', 'Seorang ilmuwan dan timnya melakukan perjalanan antar bintang untuk mencari tempat tinggal baru bagi umat manusia.', 'Jonathan Nolan', 'Matthew McConaughey, Anne Hathaway', '2014-11-07', '02:49:00', 'space'),
('The Wizard of Oz', 'Adventure, Family, Fantasy', 'wizard_of_oz.jpg', 'Seorang gadis muda bertualang ke dunia fantasi untuk menemukan jalan pulang, bertemu teman-teman baru sepanjang perjalanan.', 'L. Frank Baum', 'Judy Garland, Frank Morgan', '1939-08-15', '01:42:00', 'classic'),
('Casablanca', 'Drama, Romance, War', 'casablanca.jpg', 'Seorang pemilik kafe di Casablanca menjadi terlibat dalam perjuangan cinta dan perjuangan melawan Nazi selama Perang Dunia II.', 'Julius J. Epstein', 'Humphrey Bogart, Ingrid Bergman', '1942-11-26', '01:42:00', 'classic'),
('The Shining', 'Horror, Mystery, Thriller', 'shining.jpg', 'Seorang penulis yang sedang mengalami masalah kejiwaan diisolasi bersama keluarganya di hotel terpencil yang penuh dengan kekuatan jahat.', 'Stephen King', 'Jack Nicholson, Shelley Duvall', '1980-05-23', '02:26:00', 'psychological'),
('Fight Club', 'Drama', 'fight_club.jpg', 'Seorang pria yang bosan dengan hidupnya memulai sebuah klub perkelahian bawah tanah yang mengarah pada peristiwa-peristiwa berbahaya.', 'Chuck Palahniuk', 'Brad Pitt, Edward Norton', '1999-10-15', '02:19:00', 'cult'),
('The Godfather: Part II', 'Crime, Drama', 'godfather_part2.jpg', 'Melanjutkan kisah keluarga Corleone, film ini memperlihatkan awal mula perjalanan Michael Corleone menjadi seorang pemimpin mafia.', 'Francis Ford Coppola', 'Al Pacino, Robert De Niro', '1974-12-20', '03:22:00', 'epic'),
('The Princess Bride', 'Adventure, Family, Fantasy', 'princess_bride.jpg', 'Seorang pemuda menceritakan kisah petualangan menakjubkan dengan aksi heroik, cinta, dan persahabatan kepada kakeknya.', 'William Goldman', 'Cary Elwes, Robin Wright', '1987-09-25', '01:38:00', 'classic'),
('The Godfather: Part III', 'Crime, Drama', 'godfather_part3.jpg', 'Michael Corleone berusaha untuk melepaskan diri dari dunia kejahatan dan menghadapi masa lalu keluarganya.', 'Mario Puzo', 'Al Pacino, Diane Keaton', '1990-12-25', '02:42:00', 'mafia'),
('The Revenant', 'Action, Drama, Adventure', 'revenant.jpg', 'Seorang pemburu berjuang untuk bertahan hidup setelah dikhianati dan diserang oleh beruang.', 'Alejandro González Iñárritu', 'Leonardo DiCaprio, Tom Hardy', '2015-12-25', '02:36:00', 'survival'),
('The Social Network', 'Biography, Drama', 'social_network.jpg', 'Kisah pendirian Facebook dan perjuangan Mark Zuckerberg untuk mempertahankan kontrol terhadap perusahaannya.', 'Aaron Sorkin', 'Jesse Eisenberg, Andrew Garfield', '2010-10-01', '02:00:00', 'biography'),
('Mad Max: Fury Road', 'Action, Adventure, Science Fiction', 'mad_max.jpg', 'Seorang wanita berusaha melarikan diri dari tiran di dunia pasca-apokaliptik, dengan bantuan seorang pembawa jalan.', 'George Miller', 'Tom Hardy, Charlize Theron', '2015-05-15', '02:00:00', 'post-apocalyptic'),
('The Usual Suspects', 'Crime, Drama, Mystery', 'usual_suspects.jpg', 'Seorang detektif mencoba mengungkap siapa yang bertanggung jawab atas kejahatan yang melibatkan lima penjahat yang telah berkumpul bersama.', 'Christopher McQuarrie', 'Kevin Spacey, Gabriel Byrne', '1995-08-16', '01:46:00', 'thriller'),
('The Terminator', 'Action, Science Fiction, Thriller', 'terminator.jpg', 'Seorang pembunuh cyborg dikirim kembali ke masa lalu untuk membunuh ibu pemimpin perlawanan manusia.', 'James Cameron', 'Arnold Schwarzenegger, Linda Hamilton', '1984-10-26', '01:47:00', 'action'),
('The Great Gatsby', 'Drama, Romance', 'great_gatsby.jpg', 'Kisah seorang pria misterius yang berusaha meraih kembali cinta lamanya dalam dunia penuh kemewahan dan dekadensi.', 'Baz Luhrmann', 'Leonardo DiCaprio, Carey Mulligan', '2013-05-10', '02:23:00', 'romantic'),
('The Grand Budapest Hotel', 'Comedy, Drama', 'grand_budapest_hotel.jpg', 'Seorang penjaga hotel yang terhormat berjuang untuk membuktikan bahwa dia tidak bersalah atas tuduhan pencurian sebuah lukisan langka.', 'Wes Anderson', 'Ralph Fiennes, F. Murray Abraham', '2014-03-28', '01:39:00', 'quirky'),
('The Martian', 'Adventure, Drama, Science Fiction', 'martian.jpg', 'Seorang astronaut yang terdampar di Mars berjuang untuk bertahan hidup dan mencoba untuk kembali ke Bumi.', 'Drew Goddard', 'Matt Damon, Jessica Chastain', '2015-10-02', '02:24:00', 'survival'),
('Guardians of the Galaxy', 'Action, Adventure, Comedy', 'guardians_of_the_galaxy.jpg', 'Sekelompok individu yang tidak biasa menjadi pahlawan galaksi dengan melawan ancaman besar terhadap alam semesta.', 'James Gunn', 'Chris Pratt, Zoe Saldana', '2014-08-01', '02:01:00', 'superhero');

INSERT INTO cinema (name, movies_id)
VALUES
    ('Cinema A', 1),
    ('Cinema B', 2),
    ('Cinema C', 3),
    ('Cinema D', 4),
    ('Cinema E', 5),
    ('Cinema F', 6),
    ('Cinema G', 7),
    ('Cinema H', 8),
    ('Cinema I', 9),
    ('Cinema J', 10),
    ('Cinema K', 11),
    ('Cinema L', 12),
    ('Cinema M', 13),
    ('Cinema N', 14),
    ('Cinema O', 15),
    ('Cinema P', 16),
    ('Cinema Q', 17),
    ('Cinema R', 18),
    ('Cinema S', 19),
    ('Cinema T', 20),
    ('Cinema U', 21),
    ('Cinema V', 22),
    ('Cinema W', 23),
    ('Cinema X', 24),
    ('Cinema Y', 25),
    ('Cinema Z', 26),
    ('Cinema AA', 27),
    ('Cinema AB', 28),
    ('Cinema AC', 29),
    ('Cinema AD', 30),
    ('Cinema AE', 31),
    ('Cinema AF', 32),
    ('Cinema AG', 33),
    ('Cinema AH', 34),
    ('Cinema AI', 35),
    ('Cinema AJ', 36),
    ('Cinema AK', 37),
    ('Cinema AL', 38),
    ('Cinema AM', 39),
    ('Cinema AN', 40),
    ('Cinema AO', 1),
    ('Cinema AP', 2),
    ('Cinema AQ', 3),
    ('Cinema AR', 4),
    ('Cinema AS', 5),
    ('Cinema AT', 6),
    ('Cinema AU', 7),
    ('Cinema AV', 8),
    ('Cinema AW', 9),
    ('Cinema AX', 10),
    ('Cinema AY', 11),
    ('Cinema AZ', 12),
    ('Cinema BA', 13),
    ('Cinema BB', 14),
    ('Cinema BC', 15),
    ('Cinema BD', 16),
    ('Cinema BE', 17),
    ('Cinema BF', 18),
    ('Cinema BG', 19),
    ('Cinema BH', 20),
    ('Cinema BI', 21),
    ('Cinema BJ', 22),
    ('Cinema BK', 23),
    ('Cinema BL', 24),
    ('Cinema BM', 25),
    ('Cinema BN', 26),
    ('Cinema BO', 27),
    ('Cinema BP', 28),
    ('Cinema BQ', 29),
    ('Cinema BR', 30),
    ('Cinema BS', 31),
    ('Cinema BT', 32),
    ('Cinema BU', 33),
    ('Cinema BV', 34),
    ('Cinema BW', 35),
    ('Cinema BX', 36),
    ('Cinema BY', 37),
    ('Cinema BZ', 38),
    ('Cinema CA', 39),
    ('Cinema CB', 40),
    ('Cinema CC', 1),
    ('Cinema CD', 2),
    ('Cinema CE', 3),
    ('Cinema CF', 4),
    ('Cinema CG', 5),
    ('Cinema CH', 6),
    ('Cinema CI', 7),
    ('Cinema CJ', 8),
    ('Cinema CK', 9),
    ('Cinema CL', 10),
    ('Cinema CM', 11),
    ('Cinema CN', 12),
    ('Cinema CO', 13),
    ('Cinema CP', 14),
    ('Cinema CQ', 15),
    ('Cinema CR', 16),
    ('Cinema CS', 17),
    ('Cinema CT', 18),
    ('Cinema CU', 19),
    ('Cinema CV', 20),
    ('Cinema CW', 21),
    ('Cinema CX', 22),
    ('Cinema CY', 23),
    ('Cinema CZ', 24),
    ('Cinema DA', 25),
    ('Cinema DB', 26),
    ('Cinema DC', 27),
    ('Cinema DD', 28),
    ('Cinema DE', 29),
    ('Cinema DF', 30),
    ('Cinema DG', 31),
    ('Cinema DH', 32),
    ('Cinema DI', 33),
    ('Cinema DJ', 34),
    ('Cinema DK', 35),
    ('Cinema DL', 36),
    ('Cinema DM', 37),
    ('Cinema DN', 38),
    ('Cinema DO', 39),
    ('Cinema DP', 40);

INSERT INTO cinema_date (name_date, cinema_id)
VALUES
('2024-12-22', 1),
('2024-12-22', 2),
('2024-12-22', 3),
('2024-12-22', 4),
('2024-12-22', 5),
('2024-12-22', 6),
('2024-12-23', 7),
('2024-12-23', 8),
('2024-12-23', 9),
('2024-12-23', 10),
('2024-12-24', 11),
('2024-12-24', 12),
('2024-12-24', 13),
('2024-12-24', 14),
('2024-12-24', 15),
('2024-12-25', 16),
('2024-12-25', 17),
('2024-12-25', 18),
('2024-12-25', 19),
('2024-12-25', 20),
('2024-12-26', 21),
('2024-12-26', 22),
('2024-12-26', 23),
('2024-12-26', 24),
('2024-12-27', 25),
('2024-12-27', 26),
('2024-12-27', 27),
('2024-12-27', 28),
('2024-12-28', 29),
('2024-12-28', 30),
('2024-12-28', 31),
('2024-12-28', 32),
('2024-12-29', 33),
('2024-12-29', 34),
('2024-12-29', 35),
('2024-12-29', 36),
('2024-12-29', 37),
('2024-12-30', 38),
('2024-12-30', 39),
('2024-12-30', 40),
('2024-12-22', 1),
('2024-12-23', 2),
('2024-12-24', 3),
('2024-12-25', 4),
('2024-12-26', 5),
('2024-12-27', 6),
('2024-12-28', 7),
('2024-12-29', 8),
('2024-12-30', 9),
('2024-12-22', 41),
('2024-12-22', 42),
('2024-12-22', 43),
('2024-12-22', 44),
('2024-12-22', 45),
('2024-12-22', 46),
('2024-12-23', 47),
('2024-12-23', 48),
('2024-12-23', 49),
('2024-12-23', 50),
('2024-12-24', 51),
('2024-12-24', 52),
('2024-12-24', 53),
('2024-12-24', 54),
('2024-12-24', 55),
('2024-12-25', 56),
('2024-12-25', 57),
('2024-12-25', 58),
('2024-12-25', 59),
('2024-12-25', 60),
('2024-12-26', 61),
('2024-12-26', 62),
('2024-12-26', 63),
('2024-12-26', 64),
('2024-12-27', 65),
('2024-12-27', 66),
('2024-12-27', 67),
('2024-12-27', 68),
('2024-12-28', 69),
('2024-12-28', 70),
('2024-12-28', 71),
('2024-12-28', 72),
('2024-12-29', 73),
('2024-12-29', 74),
('2024-12-29', 75),
('2024-12-29', 76),
('2024-12-29', 77),
('2024-12-30', 78),
('2024-12-30', 79),
('2024-12-30', 80),
('2024-12-22', 81),
('2024-12-23', 82),
('2024-12-24', 83),
('2024-12-25', 84),
('2024-12-26', 85),
('2024-12-27', 86),
('2024-12-28', 87),
('2024-12-29', 88),
('2024-12-30', 89),
('2024-12-22', 10),
('2024-12-22', 11),
('2024-12-22', 12),
('2024-12-22', 13),
('2024-12-22', 14),
('2024-12-22', 15),
('2024-12-23', 16),
('2024-12-23', 17),
('2024-12-23', 18),
('2024-12-23', 90),
('2024-12-24', 91),
('2024-12-24', 92),
('2024-12-24', 93),
('2024-12-24', 94),
('2024-12-24', 95),
('2024-12-25', 96),
('2024-12-25', 97),
('2024-12-25', 98),
('2024-12-25', 99),
('2024-12-25', 100),
('2024-12-26', 100),
('2024-12-26', 99),
('2024-12-26', 98),
('2024-12-26', 97),
('2024-12-27', 96),
('2024-12-27', 95),
('2024-12-27', 94),
('2024-12-27', 93),
('2024-12-28', 92),
('2024-12-28', 80),
('2024-12-28', 81),
('2024-12-28', 82),
('2024-12-29', 83),
('2024-12-29', 84),
('2024-12-29', 85),
('2024-12-29', 86),
('2024-12-29', 87),
('2024-12-30', 88),
('2024-12-30', 89),
('2024-12-30', 80),
('2024-12-22', 1),
('2024-12-23', 2),
('2024-12-24', 3),
('2024-12-25', 4),
('2024-12-26', 5),
('2024-12-27', 6),
('2024-12-28', 7),
('2024-12-29', 8),
('2024-12-30', 9),
('2024-12-22', 71),
('2024-12-22', 72),
('2024-12-22', 73),
('2024-12-22', 74),
('2024-12-22', 75),
('2024-12-22', 76),
('2024-12-23', 77),
('2024-12-23', 78),
('2024-12-23', 79),
('2024-12-23', 70),
('2024-12-24', 71),
('2024-12-24', 62),
('2024-12-24', 63),
('2024-12-24', 64),
('2024-12-24', 65),
('2024-12-25', 66),
('2024-12-25', 67),
('2024-12-25', 68),
('2024-12-25', 69),
('2024-12-25', 60),
('2024-12-26', 61),
('2024-12-26', 62),
('2024-12-26', 63),
('2024-12-26', 64),
('2024-12-27', 65),
('2024-12-27', 66),
('2024-12-27', 67),
('2024-12-27', 68),
('2024-12-28', 69),
('2024-12-28', 70),
('2024-12-28', 71),
('2024-12-28', 72),
('2024-12-29', 73),
('2024-12-29', 74),
('2024-12-29', 75),
('2024-12-29', 76),
('2024-12-29', 77),
('2024-12-30', 78),
('2024-12-30', 79),
('2024-12-30', 80),
('2024-12-22', 81),
('2024-12-23', 82),
('2024-12-24', 83),
('2024-12-25', 84),
('2024-12-26', 85),
('2024-12-27', 86),
('2024-12-28', 87),
('2024-12-29', 88),
('2024-12-30', 89);
INSERT INTO cinema_time (name_time, cinema_id)
VALUES
('10:00:00', 1),
('10:00:00', 2),
('10:00:00', 3),
('10:00:00', 4),
('10:00:00', 5),
('10:00:00', 6),
('10:00:00', 7),
('10:00:00', 8),
('10:00:00', 9),
('10:00:00', 10),
('11:00:00', 11),
('11:00:00', 12),
('11:00:00', 13),
('11:00:00', 14),
('11:00:00', 15),
('11:00:00', 16),
('11:00:00', 17),
('11:00:00', 18),
('11:00:00', 19),
('11:00:00', 20),
('12:00:00', 21),
('12:00:00', 22),
('12:00:00', 23),
('12:00:00', 24),
('12:00:00', 25),
('12:00:00', 26),
('12:00:00', 27),
('12:00:00', 28),
('12:00:00', 29),
('12:00:00', 30),
('13:00:00', 31),
('13:00:00', 32),
('13:00:00', 33),
('13:00:00', 34),
('13:00:00', 35),
('13:00:00', 36),
('13:00:00', 37),
('13:00:00', 38),
('13:00:00', 39),
('13:00:00', 40),
('14:00:00', 41),
('14:00:00', 42),
('14:00:00', 43),
('14:00:00', 44),
('14:00:00', 45),
('14:00:00', 46),
('14:00:00', 47),
('14:00:00', 48),
('14:00:00', 49),
('14:00:00', 50),
('15:00:00', 51),
('15:00:00', 52),
('15:00:00', 53),
('15:00:00', 54),
('15:00:00', 55),
('15:00:00', 56),
('15:00:00', 57),
('15:00:00', 58),
('15:00:00', 59),
('15:00:00', 60),
('16:00:00', 61),
('16:00:00', 62),
('16:00:00', 63),
('16:00:00', 64),
('16:00:00', 65),
('16:00:00', 66),
('16:00:00', 67),
('16:00:00', 68),
('16:00:00', 69),
('16:00:00', 70),
('17:00:00', 71),
('17:00:00', 72),
('17:00:00', 73),
('17:00:00', 74),
('17:00:00', 75),
('17:00:00', 76),
('17:00:00', 77),
('17:00:00', 78),
('17:00:00', 79),
('17:00:00', 80),
('18:00:00', 81),
('18:00:00', 82),
('18:00:00', 83),
('18:00:00', 84),
('18:00:00', 85),
('18:00:00', 86),
('18:00:00', 87),
('18:00:00', 88),
('18:00:00', 89),
('18:00:00', 90),
('19:00:00', 91),
('19:00:00', 92),
('19:00:00', 93),
('19:00:00', 94),
('19:00:00', 95),
('19:00:00', 96),
('19:00:00', 97),
('19:00:00', 98),
('19:00:00', 99),
('19:00:00', 100);

SELECT movies.id, movies.tittle, movies.genre,
	movies.synopsis, movies.author, movies.actors,
	movies.release_date, movies.duration, cinema.name, 
	cinema_date.name_date, cinema_time.name_time,
	cinema_location.name_location
	FROM movies 
    JOIN cinema ON cinema.movies_id = movies.id
    JOIN cinema_date ON cinema_date.cinema_id = cinema.id
    JOIN cinema_time ON cinema_time.cinema_id = cinema.id
    JOIN cinema_location ON cinema_location.cinema_id = cinema.id
    WHERE movies.id = 1;

select * from orders;

select 

UPDATE users
SET email=$1, password=$2
FROM profile
WHERE users.id=$3
AND users.id = profile.user_id
RETURNING users.email, users.password;

SELECT * from users;