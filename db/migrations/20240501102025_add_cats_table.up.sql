CREATE TABLE cats
(
    id UUID PRIMARY KEY,
    name VARCHAR(30) NOT NULL CHECK (LENGTH(name) >= 1 AND LENGTH(name) <= 30),
    race VARCHAR(30) NOT NULL CHECK (race IN ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman')),
    sex VARCHAR(6) NOT NULL CHECK (sex IN ('male', 'female')),
    ageInMonth INT NOT NULL CHECK (ageInMonth >= 1 AND ageInMonth <= 120082),
    description VARCHAR(200) NOT NULL CHECK (LENGTH(description) >= 1 AND LENGTH(description) <= 200),
    imageUrls VARCHAR
    [] NOT NULL,
    created_at timestamptz default now
    ()
)
