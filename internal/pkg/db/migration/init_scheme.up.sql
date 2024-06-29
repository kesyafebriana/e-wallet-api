CREATE DATABASE ewallet_db;
-- DROP DATABASE ewallet_db WITH (FORCE);

SET TIMEZONE TO 'Asia/Jakarta';

CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE gacha_boxes(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    amount DECIMAL NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE wallets(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    wallet_number CHAR(13) UNIQUE DEFAULT NULL,
    balance DECIMAL NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE OR REPLACE FUNCTION wallet_number_insert() RETURNS TRIGGER AS $$
BEGIN
    NEW.wallet_number := CONCAT('777', LPAD(CAST(NEW.id AS VARCHAR), 10, '0'));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE TRIGGER wallet_number_insert
BEFORE INSERT ON wallets
FOR EACH ROW
EXECUTE FUNCTION wallet_number_insert();

CREATE TABLE password_tokens(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    token VARCHAR NOT NULL UNIQUE,
    expired_at TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '10 minutes',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE user_gachas(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    wallet_id BIGINT NOT NULL,
    gacha_box_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id),
    FOREIGN KEY (gacha_box_id) REFERENCES gacha_boxes(id)
);

CREATE TABLE gacha_attempts(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transactions(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    sender_wallet_id CHAR(13) NOT NULL,
    recipient_wallet_id CHAR(13) NOT NULL,
    amount DECIMAL NOT NULL,
    source_of_funds VARCHAR(20) NOT NULL,
    description VARCHAR(35) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (sender_wallet_id) REFERENCES wallets(wallet_number),
    FOREIGN KEY (recipient_wallet_id) REFERENCES wallets(wallet_number)
);