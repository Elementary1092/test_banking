CREATE TABLE IF NOT EXISTS "accounts" (
    "number"      VARCHAR(50) PRIMARY KEY,
    "customer_id" uuid NOT NULL,
    "currency"    VARCHAR(5) NOT NULL,
    "balance"     NUMERIC(20, 2) NOT NULL,
    "created_at"  TIMESTAMP NOT NULL,

    CONSTRAINT fk_account_owner_id
        FOREIGN KEY("customer_id") REFERENCES "customers"("uuid")
            ON DELETE CASCADE
);