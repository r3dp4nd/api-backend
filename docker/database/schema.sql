DROP DATABASE IF EXISTS mydatabase;
CREATE DATABASE IF NOT EXISTS mydatabase;

USE mydatabase;

DROP TABLE IF EXISTS cities;
CREATE TABLE cities
(
    id   INT(250) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(250)

) ENGINE = InnoDB;

DROP TABLE IF EXISTS customers;
CREATE TABLE customers
(
    id        INT(250) AUTO_INCREMENT PRIMARY KEY NOT NULL,
    dni       VARCHAR(8) UNIQUE NOT NULL,
    name      VARCHAR(250)      NOT NULL,
    last_name VARCHAR(250)      NOT NULL,
    telephone VARCHAR(9)        NOT NULL,
    email     VARCHAR(250)      NOT NULL,
    birthdate DATE              NOT NULL,
    city_id   INT(250) NOT NULL,
    enabled   BIT(1)            NOT NULL DEFAULT 1,

    CONSTRAINT `fk_customer_cities` FOREIGN KEY (city_id) REFERENCES cities (id) ON DELETE CASCADE

) ENGINE = InnoDB;

INSERT INTO cities (name)
VALUES ('Amazonas'),
       ('Áncash'),
       ('Apurímac'),
       ('Arequipa'),
       ('Ayacucho'),
       ('Cajamarca'),
       ('Callao'),
       ('Cusco'),
       ('Huancavelica'),
       ('Huánuco'),
       ('Ica'),
       ('Junín'),
       ('La Libertad'),
       ('Lambayeque'),
       ('Lima'),
       ('Loreto'),
       ('Madre de Dios'),
       ('Moquegua'),
       ('Pasco'),
       ('Piura'),
       ('Puno'),
       ('San Martin'),
       ('Tacna'),
       ('Tumbes'),
       ('Ucayali');

INSERT INTO customers (dni, name, last_name, telephone, email, birthdate, city_id)
VALUES ('28282801', 'pedro1', 'lazo', '999782882', 'pedro@mail.com', '1994-06-20', 1),
       ('28282802', 'pedro1', 'lazo', '999782882', 'pedro@mail.com', '1994-06-20', 1);