CREATE TABLE IF NOT EXISTS task
(
    id          uuid     PRIMARY KEY NOT NULL,
    title       varchar(255)     NOT NULL,
    description varchar(255)     NOT NULL,
    category varchar(255) NOT NULL,
    status   boolean NOT NULL,
    date       TIMESTAMP NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
)