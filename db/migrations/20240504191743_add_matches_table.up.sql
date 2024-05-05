CREATE TABLE matches
(
    id UUID NOT NULL PRIMARY KEY,
    created_at timestamptz default now(),
    issuedBy UUID REFERENCES users(id),
    issuerCatId UUID NOT NULL,
    receiverCatId UUID NOT NULL,
    message VARCHAR(120) NOT NULL CHECK (LENGTH(message) >= 5 AND LENGTH(message) <= 120),
    status VARCHAR(10) NOT NULL
);
