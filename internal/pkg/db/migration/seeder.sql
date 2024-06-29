INSERT INTO users(email, name, password)
VALUES
('alice@email.com', 'Alice', '1'),
('bob@email.com', 'Bob', '2'),
('charlie@email.com', 'Charlie', '3'),
('david@email.com', 'David', '4'),
('eve@email.com', 'Eve', '5');

INSERT INTO wallets(user_id)
VALUES
(1),
(2),
(3),
(4),
(5);

INSERT INTO transactions(sender_wallet_id, recipient_wallet_id, amount, source_of_funds, description)
VALUES
('7770000000001','7770000000001',50000,'Bank Transfer','Top Up From Bank Transfer'),
('7770000000001','7770000000002',50000,'Transfer','Bayar uang makan'),
('7770000000002','7770000000002',200000,'Credit Card','Top Up From Credit Card'),
('7770000000001','7770000000001',100000,'Bank Transfer','Top Up From Bank Transfer'),
('7770000000002','7770000000001',1000,'Transfer','Balikin Uang'),
('7770000000003','7770000000003',50000,'Bank Transfer','Top Up From Bank Transfer'),
('7770000000003','7770000000004',50000,'Transfer','Bayar uang makan'),
('7770000000004','7770000000004',200000,'Credit Card','Top Up From Credit Card'),
('7770000000003','7770000000003',300000,'Bank Transfer','Top Up From Bank Transfer'),
('7770000000004','7770000000001',1000,'Transfer','Balikin Uang'),
('7770000000005','7770000000005',60000,'Credit Card','Top Up From Credit Card'),
('7770000000003','7770000000003',50000,'Reward','Reward From Gacha'),
('7770000000003','7770000000004',50000,'Transfer','Bayar uang makan'),
('7770000000004','7770000000004',200000,'Credit Card','Top Up From Credit Card'),
('7770000000003','7770000000003',300000,'Bank Transfer','Top Up From Bank Transfer'),
('7770000000004','7770000000001',1000,'Transfer','Balikin Uang'),
('7770000000005','7770000000005',60000,'Reward','Reward from Gacha'),
('7770000000004','7770000000004',200000,'Credit Card','Top Up From Credit Card'),
('7770000000003','7770000000003',300000,'Bank Transfer','Top Up From Bank Transfer'),
('7770000000004','7770000000001',1000,'Transfer','Balikin Uang');