
CREATE TABLE deals {
    id INT PRIMARY KEY AUTO_INCREMENT,
    offer_id INT NOT NULL,
    need_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (real_estate_object_id) REFERENCES real_estate_objects(id),
};