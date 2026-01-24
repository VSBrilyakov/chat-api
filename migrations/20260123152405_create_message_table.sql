-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    chat_id integer,
    text varchar(5000) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (chat_id) REFERENCES chat(id) ON DELETE CASCADE
);

INSERT INTO message (chat_id, text, created_at)
VALUES
    (1,'Hello bro', '01-01-2021 10:12:37'),
    (1,'Hi, how are you?', '01-02-2021 17:55:27'),
    (3,'What''s up?','12-21-2025 16:12:10'),
    (1,'I''m looking for a Go developer job','01-02-2021 21:49:01'),
    (1,'That''s cool, good luck!', '03-09-2022 19:41:38'),
    (3,'I''m ok, ','12-25-2025 15:15:59'),
    (2,'Wake up, Neo...', '10-09-2022 04:20:00'),
    (3,'Would you like to go to the gym together?','08-01-2026 09:02:12');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE message;