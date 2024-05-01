CREATE TABLE users
(
    id UUID NOT NULL PRIMARY KEY,
    created_at timestamptz default now(),
    email VARCHAR(255) NOT NULL UNIQUE,
    name TEXT NOT NULL check ( char_length(name) >= 5 AND char_length(name) <= 50),
    password_hash TEXT NOT NULL check ( char_length(name) >= 5 AND char_length(name) <= 15)
)
