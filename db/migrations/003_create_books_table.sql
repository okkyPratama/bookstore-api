-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    release_year INTEGER CHECK (release_year BETWEEN 1980 AND 2024),
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR(20),
    category_id INTEGER REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE books;
-- +migrate StatementEnd