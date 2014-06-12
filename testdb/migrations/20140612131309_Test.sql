
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE dameleon (
    id INTEGER NOT NULL AUTO_INCREMENT,
    word VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE dameleon;
