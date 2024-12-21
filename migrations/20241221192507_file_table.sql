-- +goose Up
-- +goose StatementBegin
CREATE TABLE files
(
    id            SERIAL PRIMARY KEY,
    date_create   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    file_name     VARCHAR(255) NOT NULL,
    original_name VARCHAR(255),
    file_size     INTEGER   DEFAULT 0,
    sub_dir        VARCHAR(12),
    content_type  VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files
-- +goose StatementEnd
