-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS chat (
    id SERIAL PRIMARY KEY,
    title varchar(200) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_chat_title ON chat(title);

INSERT INTO chat (title, created_at)
VALUES
    ('first chat', '01-02-2021 04:12:27'),
    ('second chat', '10-31-2022 21:49:39'),
    ('third chat','10-21-2025 16:06:14');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP INDEX IF EXISTS idx_chat_title;

DROP TABLE chat;