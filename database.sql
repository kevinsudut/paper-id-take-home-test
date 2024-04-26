CREATE DATABASE paperid;

CREATE TABLE wallets (
    wallet_id VARCHAR PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    amount NUMERIC NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

CREATE TABLE histories (
    transaction_id VARCHAR PRIMARY KEY,
    wallet_id VARCHAR NOT NULL,
    unique_id VARCHAR UNIQUE NOT NULL,
    type INT2 NOT NULL,
    amount NUMERIC NOT NULL,
    notes VARCHAR NOT NULL,
    metadata JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

INSERT INTO wallets(wallet_id, user_id, amount, created_at) VALUES
('f7a2da93-975e-4ee5-abc8-12512a03243a', '4830fd33-9624-470d-b023-f09ea81506a2', 583612, NOW()),
('990c6daa-c727-4fd7-b00c-c76d222777bf', 'f0783eab-2358-45c2-88b5-2fdc434f83e0', 844802, NOW()),
('71949238-d44a-4bed-92b5-ccb569d92913', '8157ed01-3ddf-4071-9704-46be399e1c4e', 593828, NOW()),
('e3890191-49fe-4e1c-8cf0-f8f7fea1501f', '4830fd33-9624-470d-b023-f09ea81506a2', 632092, NOW());
