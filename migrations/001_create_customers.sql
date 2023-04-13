CREATE TABLE IF NOT EXISTS "customers" (
    "uuid"       uuid PRIMARY KEY,
    "email"      VARCHAR NOT NULL UNIQUE,
    "password"   VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);