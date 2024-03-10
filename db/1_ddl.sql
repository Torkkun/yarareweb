CREATE TABLE users (
    id TEXT NOT NULL,
    name TEXT NOT NULL ,
    password TEXT NOT NULL DEFAULT '',
    PRIMARY KEY(id)
);

CREATE TABLE purchase (
    product_id TEXT NOT NULL,
    product_name TEXT NOT NULL,
    user_id TEXT NOT NULL,
    purchase_date TEXT NOT NULL,
    PRIMARY KEY(product_id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);