CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    card_number VARCHAR UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    balance INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE terminal (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE gate (
    id SERIAL PRIMARY KEY,
    terminal_id INTEGER REFERENCES terminal (id),
    gate_number VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),
    gate_id INTEGER REFERENCES gate (id),
    terminal_id INTEGER REFERENCES terminal (id),
    timestamp TIMESTAMP NOT NULL,
    action VARCHAR CHECK (
        action IN ('checkin', 'checkout')
    ) NOT NULL,
    synced BOOLEAN DEFAULT FALSE,
    fare INTEGER NOT NULL,
    entry_terminal_id INTEGER REFERENCES terminal (id),
    exit_terminal_id INTEGER REFERENCES terminal (id)
);

CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    fullname VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE sync_queue (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER REFERENCES transaction (id),
    retry_count INTEGER DEFAULT 0,
    last_attempt TIMESTAMP,
    status VARCHAR CHECK (
        status IN (
            'pending',
            'success',
            'failed'
        )
    ) NOT NULL DEFAULT 'pending'
);

INSERT INTO
    admin (
        username,
        fullname,
        password,
        created_at,
        updated_at,
        deleted_at,
        is_deleted
    )
VALUES (
        'admin1',
        'Dzaru Rizky Fathan Fortuna',
        '$2a$12$6FVo9T42ASBtapVIXzD3M.3kA7nIQJ/vRikLhCX.ICTCpsEaiy.4u',
        now(),
        now(),
        0,
        FALSE
    );