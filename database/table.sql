CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    fullname VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at INT,
    is_deleted BOOLEAN
);

CREATE TABLE terminal (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at INT,
    is_deleted BOOLEAN
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