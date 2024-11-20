INSERT INTO users (name, email)
VALUES 
('Alice Johnson', 'alice.johnson@example.com'),
('Bob Smith', 'bob.smith@example.com'),
('Charlie Brown', 'charlie.brown@example.com'),
('Diana Prince', 'diana.prince@example.com'),
('Edward Norton', 'edward.norton@example.com'),
('Fiona Carter', 'fiona.carter@example.com');

INSERT INTO product_categories (name)
VALUES 
('Steam Wallet'),
('Game'),
('Diamond'),
('Crystal'),
('UC');

INSERT INTO products (name, price, quantity, categoryID)
VALUES 
('Steam Wallet 60.000', 61899, 3, 1),
('Steam Wallet 120.000', 123599, 1, 1),
('Black Myth Wukong', 890000, 1, 2),
('370 Diamond', 101430, 2, 3),
('89 Diamond', 25000, 3, 3),
('60 Crystal', 10000, 20, 4),
('30 UC', 6000, 10, 5),
('750 UC', 151000, 5, 5),
('Steam Wallet 60.000', 61899, 2, 1),
('God of War Ragnarok', 750000, 1, 2);

INSERT INTO transactions (userID, productID, purchasedAt)
VALUES 
(1, 1, '2024-11-01'),
(2, 3, '2024-11-02'),
(3, 4, '2024-11-03'),
(1, 2, '2024-11-04'),
(2, 5, '2024-11-05'),
(4, 6, '2024-11-06'),
(5, 7, '2024-11-07'),
(6, 8, '2024-11-08'),
(3, 9, '2024-11-09'),
(4, 10, '2024-11-10');
