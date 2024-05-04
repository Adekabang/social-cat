CREATE TABLE matches
(
    id UUID NOT NULL PRIMARY KEY,
    created_at timestamptz default now(),
    issuedBy UUID REFERENCES users(id),
    issuerCatId UUID REFERENCES cats(id),
    receiverCatId UUID REFERENCES cats(id),
    message VARCHAR(120) NOT NULL CHECK (LENGTH(message) >= 5 AND LENGTH(message) <= 120),
    status VARCHAR(10) NOT NULL
);
