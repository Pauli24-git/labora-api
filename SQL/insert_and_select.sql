INSERT INTO public.items(
	customer_name, order_date, product, quantity, price)

VALUES ('Pauli', '03-05-2023', 'Zapatillas', 2, 15000),
       ('Lolis', '03-05-2023', 'Borcegos', 1, 75000),
       ('Misita', '03-05-2023', 'Sandalias', 3, 90000);
       ('Oli', '03-05-2023', 'Botines', 2, 10000)
       ('Dua', '03-05-2023', 'Botas', 4, 90000);
       ('Lipa', '03-05-2023', 'Sandalias', 4, 20000);       


SELECT id, customer_name, order_date, product, quantity, price
FROM items
WHERE quantity > 2 AND price > 50;
