-- migrations/002_create_requests_table.sql

CREATE TABLE IF NOT EXISTS requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    full_name TEXT NOT NULL,
    date_of_birth TEXT NOT NULL,
    illness TEXT NOT NULL,
    location TEXT NOT NULL,
    phone_number TEXT NOT NULL
);
