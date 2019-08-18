-- +goose Up
CREATE TABLE house_events
(
    id                 int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    house_id           int(10) UNSIGNED NOT NULL,
    name               VARCHAR(255),
    candidate_datetime DATETIME                  DEFAULT NULL,
    shop_id            VARCHAR(255)              DEFAULT NULL,
    created_at         TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT events_fk_house FOREIGN KEY (house_id) REFERENCES houses (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `house_events` (`id`, `house_id`, `name`, `candidate_datetime`, `shop_id`) VALUES
(1, 1, 'イベント１', '2019-08-07 00:00:00', 'hogehogeshop'),
(2, 1, 'イベント2', '2019-08-22 00:00:00', 'hogehgoeshop2'),
(3, 2, 'イベント３', '2019-08-17 00:00:00', 'hogehogeshop'),
(4, 2, 'イベント４', '2019-08-29 00:00:00', 'hogehogeshop'),
(5, 4, 'イベント5', '2019-08-29 00:00:00', 'hogehgoe_super_shop');

-- +goose Down
DROP TABLE house_events;
