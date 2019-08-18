-- +goose Up
CREATE TABLE house_event_questionnaire_dates
(
    id                 int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    questionnaire_id   int(10) UNSIGNED NOT NULL,
    candidate_datetime TIMESTAMP        NOT NULL,
    created_at         TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT house_event_questionnaire_dates_fk_questionnaire FOREIGN KEY (questionnaire_id) REFERENCES house_event_questionnaires (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `house_event_questionnaire_dates` (`id`, `questionnaire_id`, `candidate_datetime`) VALUES
(1, 1, '2019-08-07 00:00:00'),
(2, 1, '2019-08-14 00:00:00'),
(3, 1, '2019-08-14 00:00:00'),
(4, 2, '2019-08-21 00:00:00'),
(5, 2, '2019-08-17 00:00:00'),
(6, 2, '2019-08-21 00:00:00'),
(7, 3, '2019-08-20 00:00:00'),
(8, 3, '2019-08-20 00:00:00'),
(9, 4, '2019-08-23 00:00:00'),
(10, 4, '2019-08-20 00:00:00'),
(11, 5, '2019-08-21 00:00:00'),
(12, 6, '2019-08-30 00:00:00'),
(13, 7, '2019-08-21 00:00:00'),
(14, 8, '2019-08-23 00:00:00'),
(15, 9, '2019-08-19 00:00:00'),
(16, 5, '2019-08-25 00:00:00'),
(17, 6, '2019-08-19 00:00:00'),
(18, 6, '2019-08-13 00:00:00');

-- +goose Down
DROP TABLE house_event_questionnaire_dates;
