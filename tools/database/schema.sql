CREATE TABLE IF NOT EXISTS "task"
(
    id          uuid     PRIMARY KEY NOT NULL,
    user_id     uuid     NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    title       varchar(255)     NOT NULL,
    description varchar(255)     NOT NULL,
    category varchar(255) NOT NULL,
    status   boolean NOT NULL,
    date       TIMESTAMP NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "user"
(
    id          uuid     PRIMARY KEY NOT NULL,
    username       varchar(255)     NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password   varchar(255) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
)


