-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_token (
      id SERIAL PRIMARY KEY,
      uid INT NOT NULL,
      token VARCHAR(255) NOT NULL,
      refresh_token VARCHAR(255) NOT NULL,
      created_date TIMESTAMP NOT NULL,
      expiration_date TIMESTAMP NOT NULL,
      refresh_expiration_date TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_token
-- +goose StatementEnd
