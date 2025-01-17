
CREATE TABLE real_estate_object_types (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO real_estate_object_types (id, name) 
VALUES 
    (1, 'apartment'), 
    (2, 'house'), 
    (3, 'land');

DROP TABLE IF EXISTS real_estate_objects;
CREATE TABLE real_estate_objects (
    id INT PRIMARY KEY AUTO_INCREMENT,
    real_estate_object_type_id INT NOT NULL,
    latitude FLOAT DEFAULT NULL CHECK (latitude >= -90 AND latitude <= 90),
    longitude FLOAT DEFAULT NULL CHECK (longitude >= -180 AND longitude <= 180),
    city VARCHAR(255) DEFAULT NULL,
    street VARCHAR(255) DEFAULT NULL,
    house_number VARCHAR(255) DEFAULT NULL,
    apartment_number VARCHAR(255) DEFAULT NULL,

    FOREIGN KEY (real_estate_object_type_id) REFERENCES real_estate_object_types(id)
);
