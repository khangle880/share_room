-- +goose Up
CREATE TYPE userrole AS ENUM ('ADMIN', 'ROOMMATE', 'CAPTAIN');
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    username TEXT UNIQUE NOT NULL,
    lastname VARCHAR(255),
    firstname VARCHAR(255),
    bio TEXT,
    avatar VARCHAR(255),
    hashed_password VARCHAR(255) NOT NULL,
    role UserRole NOT NULL,
    phone VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    UNIQUE(username, email)
);
CREATE TABLE icons (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    type VARCHAR(255)
);
CREATE TYPE categorytype AS ENUM ('INCOME', 'OUTCOME');
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    type CategoryType NOT NULL,
    icon_id UUID NOT NULL REFERENCES icons(id),
    parent_id UUID REFERENCES categories(id) ON DELETE CASCADE
);
CREATE TABLE events (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    icon_id UUID NOT NULL REFERENCES icons(id),
    background VARCHAR(255) NOT NULL
);
CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    captain UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    avatar VARCHAR(255),
    background VARCHAR(255)
);
CREATE TABLE budgets (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    balance BIGINT NOT NULL DEFAULT 0,
    icon_id UUID NOT NULL REFERENCES icons(id) ON DELETE CASCADE,
    member_ids UUID [],
    room_id UUID REFERENCES rooms(id) ON DELETE CASCADE
);
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    exc_time TIMESTAMP NOT NULL,
    description TEXT,
    budget_id UUID NOT NULL REFERENCES budgets(id) ON DELETE CASCADE,
    creators_ids UUID [],
    partner_ids UUID [],
    events_id UUID,
    images VARCHAR (255) []
);
-- +goose Down
DROP TABLE categories;
DROP TABLE transactions;
DROP TABLE budgets;
DROP TABLE rooms;
DROP TABLE icons;
DROP TABLE users;
DROP TYPE UserRole;
DROP TYPE CategoryType;