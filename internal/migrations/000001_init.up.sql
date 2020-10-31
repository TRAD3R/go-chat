CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    name varchar(255),
    username varchar(255) unique not null,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE IF NOT EXISTS items
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean default false
);

CREATE TABLE IF NOT EXISTS user_list
(
    id serial not null unique,
    user_id integer not null,
    list_id integer not null
);

CREATE TABLE IF NOT EXISTS list_item
(
    id serial not null unique,
    list_id integer not null,
    item_id integer not null
);
