CREATE TABLE IF NOT EXISTS `symmetric_keys` (
    key_id VARCHAR(255) PRIMARY KEY,
    account_id VARCHAR(255) NOT NULL,
    region ENUM('us-east-1', 'us-east-2', 'us-west-1', 'ru-west-1', 'eu-south-1', 'eu-west-1') NOT NULL,
    description VARCHAR(500),
    key_spec ENUM('AES-256 HSM', 'AES-256', 'AES-192', 'AES-128') NOT NULL,
    ciphertext TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL
);