-- Database: laboraProject1

-- DROP DATABASE IF EXISTS "laboraProject1";

CREATE DATABASE "laboraProject1"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Spanish_Spain.1252'
    LC_CTYPE = 'Spanish_Spain.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    order_date DATE NOT NULL,
    product VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL
);

INSERT INTO items (customer_name, order_date, product, quantity, price)

VALUES ('Pauli', '03-05-2023', 'Zapatillas', 2, 15000),
       ('Lolis', '03-05-2023', 'Borcegos', 1, 75000),
       ('Misita', '03-05-2023', 'Sandalias', 3, 90000);