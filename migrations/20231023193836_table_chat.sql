-- +goose Up
CREATE TABLE chat_server(
    id serial primary KEY,
    username text not null,
    message text not null,
    sendtime timestamp not null default now()
);


-- +goose Down
DROP TABLE chat_server;
