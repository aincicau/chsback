BEGIN;
create extension if not exists "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    CONSTRAINT pk_user_id PRIMARY KEY (id)
);

CREATE TABLE histories (
    id uuid DEFAULT uuid_generate_v4(),
    file_name varchar(255) NOT NULL,
    date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    result varchar(255) NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT pk_history_id PRIMARY KEY (id),
    CONSTRAINT fk_history_user FOREIGN KEY (user_id) REFERENCES users(id)
);

COMMIT;