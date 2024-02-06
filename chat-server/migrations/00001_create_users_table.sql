-- +goose Up
create table messages (
    id serial primary key,
    from text not null,
    text text not null,
    timestamp timestamp
);

-- +goose Down
drop table messages;