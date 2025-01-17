
CREATE TABLE deals (
    id INT PRIMARY KEY AUTO_INCREMENT,
    offer_id INT NOT NULL,
    need_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (offer_id) REFERENCES offers(id),
    FOREIGN KEY (need_id) REFERENCES needs(id)
);