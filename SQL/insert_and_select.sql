INSERT INTO public.items(
	id, customer_name, order_date, product, quantity, price)

VALUES (1, 'Pauli', '03-05-2023', 'Zapatillas', 2, 15000),
       (2,'Lolis', '03-05-2023', 'Borcegos', 1, 75000),
       (3, 'Misita', '03-05-2023', 'Sandalias', 3, 90000);
       (4, 'Oli', '03-05-2023', 'Botines', 2, 10000)
       (5, 'Dua', '03-05-2023', 'Botas', 4, 90000);
       (6, 'Lipa', '03-05-2023', 'Sandalias', 4, 20000);       


SELECT id, customer_name, order_date, product, quantity, price
FROM items
WHERE quantity > 2 AND price > 50;