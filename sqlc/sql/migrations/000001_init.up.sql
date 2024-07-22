CREATE TABLE products
(
    id          VARCHAR(36) NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2)
);