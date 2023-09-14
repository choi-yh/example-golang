CREATE TABLE user
(
    id         INT       NOT NULL PRIMARY KEY,
    email      VARCHAR(255),
    password   VARCHAR(255),
    name       VARCHAR(32),
    phone      VARCHAR(32),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `email` (`email`)
) ENGINE = innoDB
;