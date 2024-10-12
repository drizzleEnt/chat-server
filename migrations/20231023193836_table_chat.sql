-- +goose Up
CREATE TABLE chats(
    id SERIAL PRIMARY KEY,
    chat_id TEXT NOT NULL
);
CREATE TABLE messages(
    id serial primary KEY,
    user_id INTEGER not null,
    message text not null,
    chat_id INTEGER REFERENCES chats (id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE chats;
DROP TABLE messages;
