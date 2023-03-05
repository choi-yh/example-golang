CREATE TABLE `user`
(
    id         INT       NOT NULL PRIMARY KEY,
    email      VARCHAR(255),
    name       VARCHAR(32),
    phone      VARCHAR(32),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = innoDB
;

CREATE TABLE user_password
(
    id       INT NOT NULL PRIMARY KEY,
    user_id  INT,
    password VARCHAR(255),

    FOREIGN KEY (user_id) REFERENCES `user` (id)
) ENGINE = innoDB
;
