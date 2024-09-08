-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE categories;
-- +migrate StatementEnd