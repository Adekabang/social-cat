create table users
(
    id            uuid primary key,
    created_at    timestamptz default now(),
    username      text not null check ( char_length(username) >= 1 AND char_length(username) <= 32),
    password_hash text not null,
    unique (username)
);
