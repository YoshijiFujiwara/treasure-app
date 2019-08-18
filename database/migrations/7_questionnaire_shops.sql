-- +goose Up
CREATE TABLE house_event_questionnaire_shops
(
    id               int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    questionnaire_id int(10) UNSIGNED NOT NULL,
    shop_id          VARCHAR(255)              DEFAULT NULL,
    created_at       TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT house_event_questionnaire_shops_fk_questionnaire FOREIGN KEY (questionnaire_id) REFERENCES house_event_questionnaires (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `house_event_questionnaire_shops` (`id`, `questionnaire_id`, `shop_id`) VALUES
(1, 1, 'fsafdsafdsa'),
(2, 1, 'fdsafdsasa'),
(3, 1, 'fdsafsa'),
(4, 2, 'fdsafdsa'),
(5, 1, 'fdsafdsa'),
(6, 5, 'fdsafdsa'),
(7, 5, 'fdsafdsa'),
(8, 3, 'sfdafdsafa'),
(9, 4, 'fdsafdsa'),
(10, 8, 'fdsafdsafsad'),
(11, 8, 'ssssssssssss'),
(12, 7, 'fdsafdsa'),
(13, 7, 'fsafdsafdsa'),
(14, 7, 'sss'),
(15, 8, 'fdsafds'),
(16, 9, 'fdsafdsfdsfdsa'),
(17, 10, 'fdsfdsfdsdfsfdss'),
(18, 10, 'fdsaaaaaaaaaaaaa'),
(19, 1, 'sfdsfdsfdsfdsfdsfsfsa'),
(20, 1, 'fdsafdsafdsafasdfdsafasd');

-- +goose Down
DROP TABLE house_event_questionnaire_shops;
