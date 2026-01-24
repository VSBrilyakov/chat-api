-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE chat (
    id SERIAL PRIMARY KEY,
    title varchar(200) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_chat_title ON chat(title);

INSERT INTO chat (title, created_at)
VALUES
    ('first chat', '01-01-2021'),
    ('second chat', '05-08-2022'),
    ('third chat','10-21-2025');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP INDEX IF EXISTS idx_chat_title;

DROP TABLE chat;