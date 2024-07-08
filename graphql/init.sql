CREATE TABLE IF NOT EXISTS categories
(
    id          VARCHAR(36)  NOT NULL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description MEDIUMTEXT,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products
(
    id          VARCHAR(36)   NOT NULL PRIMARY KEY,
    name        VARCHAR(255)  NOT NULL,
    description MEDIUMTEXT,
    price       DECIMAL(6, 2) NOT NULL,
    created_at  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_categories
(
    id          INTEGER PRIMARY KEY AUTO_INCREMENT,
    product_id  VARCHAR(36),
    category_id VARCHAR(36),
    CONSTRAINT UNIQUE (product_id, category_id),
    CONSTRAINT FOREIGN KEY (category_id) REFERENCES categories (id),
    CONSTRAINT FOREIGN KEY (product_id) REFERENCES products (id)
);