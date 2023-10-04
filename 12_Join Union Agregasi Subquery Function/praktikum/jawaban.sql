SELECT name FROM db_alta_join.users WHERE gender = 'M';

SELECT * FROM products WHERE id = 3;

SELECT *
FROM users
WHERE created_at >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
  AND name LIKE '%a%';


SELECT COUNT(*) FROM users WHERE gender = 'F';

SELECT * FROM users ORDER BY name ASC;


SELECT * FROM products LIMIT 5;


UPDATE products SET name = 'product dummy' WHERE id = 1;

UPDATE transaction_detail  SET qty = 3 WHERE product_id = 1;

DELETE FROM products WHERE product_id = 1;

DELETE FROM products WHERE product_type_id = 1;


SELECT * FROM transactions WHERE user_id IN (1, 2);

SELECT user_id, SUM(total_price) AS total_harga FROM transactions WHERE user_id = 1 GROUP BY user_id;

SELECT SUM(total_price) AS total_transaksi FROM transactions WHERE product_type_id = 2;

SELECT products.*, product_types.name AS product_type_name FROM products JOIN product_types ON products.product_type_id = product_types.product_type_id;

SELECT transactions.*, products.name AS product_name, users.name AS user_name
FROM transactions
JOIN products ON transactions.product_id = products.product_id
JOIN users ON transactions.user_id = users.user_id;

DELIMITER //

CREATE FUNCTION delete_transaction(IN transaction_id INT)
RETURNS BOOLEAN
BEGIN
  DELETE FROM transaction_detail WHERE transaction_id = transaction_id;
  DELETE FROM transactions WHERE transaction_id = transaction_id;
  RETURN TRUE;
END //

DELIMITER ;


DELIMITER //

CREATE FUNCTION update_total_qty(IN transaction_detail_id INT)
RETURNS BOOLEAN
BEGIN
  DECLARE transaction_id INT;
  DECLARE qty_deleted INT;
  SELECT transaction_id, qty INTO transaction_id, qty_deleted FROM transaction_detail WHERE transaction_detail_id = transaction_detail_id;
  UPDATE transactions SET total_qty = total_qty - qty_deleted WHERE transaction_id = transaction_id;
  RETURN TRUE;
END //

DELIMITER ;

SELECT *
FROM products
WHERE product_id NOT IN (SELECT DISTINCT product_id FROM transaction_detail);

