CREATE TABLE IF NOT EXISTS
    user_invitations (
        token bytea PRIMARY KEY,
        user_id bigint NOT NULL,
        expiry TIMESTAMP(0)
        WITH
            TIME ZONE NOT NULL
    );

ALTER TABLE users
ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT FALSE;