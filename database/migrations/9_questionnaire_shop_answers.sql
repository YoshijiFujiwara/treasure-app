-- +goose Up
CREATE TABLE house_event_questionnaire_shop_answers
(
    id                    int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id               int(10) UNSIGNED NOT NULL,
    questionnaire_shop_id int(10) UNSIGNED NOT NULL,
    ok                    BOOLEAN                   DEFAULT FALSE,
    created_at            TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, questionnaire_shop_id),
    CONSTRAINT house_event_questionnaire_shop_answers_fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT house_event_questionnaire_shop_answers_fk_shop FOREIGN KEY (questionnaire_shop_id) REFERENCES house_event_questionnaire_shops (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `house_event_questionnaire_shop_answers` (`user_id`, `questionnaire_shop_id`, `ok`)
VALUES (1, 1, 1),
       (1, 2, 0),
       (2, 1, 1),
       (2, 2, 0),
       (1, 3, 1);

-- +goose Down
DROP TABLE house_event_questionnaire_shop_answers;
