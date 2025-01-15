CREATE TABLE apartments {
    id INT PRIMARY KEY AUTO_INCREMENT,
    real_estate_object_id INT NOT NULL UNIQUE,
    floor_number INT DEFAULT NULL,
    rooms_count INT UNSIGNED DEFAULT NULL,
    area INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id)
};

CREATE TABLE houses {
    id INT PRIMARY KEY AUTO_INCREMENT,
    real_estate_object_id INT NOT NULL UNIQUE,
    floors_count INT DEFAULT NULL,
    rooms_count INT UNSIGNED DEFAULT NULL,
    area INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id)
};

CREATE TABLE lands {
    id INT PRIMARY KEY AUTO_INCREMENT,
    real_estate_object_id INT NOT NULL UNIQUE,
    area INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id)
};