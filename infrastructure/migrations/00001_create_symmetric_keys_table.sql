CREATE DATABASE IF NOT EXISTS kms;

USE kms; 

CREATE TABLE IF NOT EXISTS symmetric_keys (
  account_id VARCHAR(36) NOT NULL,
  key_id VARCHAR(255) PRIMARY KEY,
  algorithm ENUM('AES-256', 'AES-192', 'AES-128') NOT NULL,
  date_created TIMESTAMP NOT NULL,
  ciphertext_blob TEXT NOT NULL
);