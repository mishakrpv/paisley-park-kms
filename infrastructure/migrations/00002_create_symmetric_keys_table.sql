CREATE TABLE IF NOT EXISTS `symmetric_keys` (
    key_id VARCHAR(255) NOT NULL,
    account_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    region ENUM('us-east-1', 'us-east-2', 'us-west-1', 'ru-west-1', 'eu-south-1', 'eu-west-1') NOT NULL,
    algorithm ENUM('AES-256 HSM', 'AES-256', 'AES-192', 'AES-128') NOT NULL,
    rotation_period INT,
    date_created TIMESTAMP NOT NULL,
    ciphertext_blob TEXT NOT NULL,
    PRIMARY KEY (key_id)
);