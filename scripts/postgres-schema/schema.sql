create table todos
(
    id serial not null
        constraint todos_id
            primary key,
    title text not null,
    user_id integer
);

alter table todos owner to postgres;

