-- +goose Up
CREATE TABLE chats(
    id SERIAL PRIMARY KEY,
    chatId UUID UNIQUE NOT NULL
);


-- +goose Down
DROP TABLE chats;

