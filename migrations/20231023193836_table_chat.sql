-- +goose Up
CREATE TABLE chat_server(
    id serial primary KEY,
    username TEXT not null,
    message text not null
);


-- +goose Down
DROP TABLE chat_server;
