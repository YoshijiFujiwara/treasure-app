-- +goose Up
CREATE TABLE house_event_questionnaires
(
    id         int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    event_id   int(10) UNSIGNED NOT NULL,
    created_at TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT house_event_questionnaires_fk_event FOREIGN KEY (event_id) REFERENCES house_events (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `house_event_questionnaires` (`id`, `event_id`) VALUES
(4, 1),
(7, 1),
(1, 1),
(6, 1),
(8, 1),
(10, 3),
(3, 4),
(5, 4),
(2, 5),
(9, 5);

-- +goose Down
DROP TABLE house_event_questionnaires;
