-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
                       id BIGSERIAL PRIMARY KEY,
                       original_url TEXT NOT NULL,
                       short_url TEXT NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd
