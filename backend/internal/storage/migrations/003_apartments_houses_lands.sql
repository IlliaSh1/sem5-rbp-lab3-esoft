CREATE TABLE apartments (
    real_estate_object_id INT PRIMARY KEY,
    floor_number INT DEFAULT NULL,
    rooms_count INT UNSIGNED DEFAULT NULL,
    area INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id)
);

CREATE TABLE houses (
    real_estate_object_id INT PRIMARY KEY,
    floors_count INT DEFAULT NULL,
    rooms_count INT UNSIGNED DEFAULT NULL,
    area INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id)
);

CREATE TABLE lands (
    real_estate_object_id INT PRIMARY KEY,
    area INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id)
);