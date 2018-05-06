CREATE TABLE IF NOT EXISTS `clients`(
    id INT NOT NULL AUTO_INCREMENT,
    name text,
    email text,
    phone text,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `accounts`(
    id INT NOT NULL AUTO_INCREMENT,
    client_id integer,
    balance integer,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `transactions`(
    id INT NOT NULL AUTO_INCREMENT,
    account_id integer,
    amount integer,
    subject text,
    timestamp integer,
    PRIMARY KEY (id)
);