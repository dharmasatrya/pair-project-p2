CREATE TRIGGER update_product_quantity_after_transaction
AFTER INSERT ON transactions
FOR EACH ROW
BEGIN
    UPDATE products
    SET quantity = quantity - 1
    WHERE productID = NEW.productID;
END;