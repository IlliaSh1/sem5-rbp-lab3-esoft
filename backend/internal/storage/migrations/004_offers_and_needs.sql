

CREATE TABLE offers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    real_estate_object_id INT NOT NULL,
    client_id INT NOT NULL,
    realtor_id INT NOT NULL,
    price INT UNSIGNED NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (realtor_id) REFERENCES realtors(id)
);

CREATE TABLE needs (
    id INT PRIMARY KEY AUTO_INCREMENT,
    real_estate_object_type_id INT NOT NULL,
    client_id INT NOT NULL,
    realtor_id INT NOT NULL,
    min_price INT UNSIGNED DEFAULT NULL,
    max_price INT UNSIGNED DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (real_estate_object_type_id) REFERENCES real_estate_object_types(id),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (realtor_id) REFERENCES realtors(id)
);

CREATE TABLE apartment_needs (
    need_id INT PRIMARY KEY,
    min_floor_number INT DEFAULT NULL,
    max_floor_number INT DEFAULT NULL,
    min_rooms_count INT UNSIGNED DEFAULT NULL,
    max_rooms_count INT UNSIGNED DEFAULT NULL,
    min_area INT UNSIGNED DEFAULT NULL,
    max_area INT UNSIGNED DEFAULT NULL,
    FOREIGN KEY (need_id) REFERENCES needs(id)
);

CREATE TABLE house_needs (
    need_id INT PRIMARY KEY,
    min_floors_count INT DEFAULT NULL,
    max_floors_count INT DEFAULT NULL,
    min_rooms_count INT UNSIGNED DEFAULT NULL,
    max_rooms_count INT UNSIGNED DEFAULT NULL,
    min_area INT UNSIGNED DEFAULT NULL,
    max_area INT UNSIGNED DEFAULT NULL,
    FOREIGN KEY (need_id) REFERENCES needs(id)
);

CREATE TABLE land_needs (
    need_id INT PRIMARY KEY,
    min_area INT UNSIGNED DEFAULT NULL,
    max_area INT UNSIGNED DEFAULT NULL,
    FOREIGN KEY (need_id) REFERENCES needs(id)
);