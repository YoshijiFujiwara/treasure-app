-- +goose Up
CREATE TABLE house_users
(
    user_id    int(10) UNSIGNED NOT NULL,
    house_id   int(10) UNSIGNED NOT NULL,
    is_admin   BOOLEAN                   DEFAULT FALSE,
    created_at TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, house_id),
    CONSTRAINT house_users_fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ,
    CONSTRAINT house_users_fk_house FOREIGN KEY (house_id) REFERENCES houses (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `house_users` (`user_id`, `house_id`, `is_admin`) VALUES
(1, 1, 1),
(1, 3, 1),
(6, 1, 0),
(6, 4, 1),
(7, 2, 1),
(7, 4, 0),
(8, 1, 1),
(8, 3, 0),
(9, 3, 1),
(9, 4, 0);

-- +goose Down
DROP TABLE house_users;