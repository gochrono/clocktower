CREATE TABLE tokens (
    id serial PRIMARY KEY,
    value VARCHAR(80) UNIQUE NOT NULL,
    revoked BOOL NOT NULL
)
