DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) DEFAULT NULL,
    surname VARCHAR(255) DEFAULT NULL,
    patronymic VARCHAR(255) DEFAULT NULL,
    phone VARCHAR(255) DEFAULT NULL UNIQUE,
    email VARCHAR(255) DEFAULT NULL UNIQUE
);

ALTER TABLE clients
ADD CONSTRAINT chk_clients_phone_or_email 
CHECK (phone IS NOT NULL OR email IS NOT NULL);

DROP TABLE IF EXISTS realtors;
CREATE TABLE realtors (
    id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    percentage_of_commission INT UNSIGNED NOT NULL CHECK (percentage_of_commission >= 0 AND percentage_of_commission <= 100)
);
