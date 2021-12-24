BEGIN;
create extension if not exists "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    CONSTRAINT pk_user_id PRIMARY KEY (id),
);

COMMIT;