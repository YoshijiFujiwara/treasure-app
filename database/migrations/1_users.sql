-- +goose Up
CREATE TABLE users
(
    id                int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    username          VARCHAR(255),
    email             VARCHAR(255),
    password          VARCHAR(255),
    gender            ENUM ('man', 'woman', 'other'),
    birthday          DATE                      DEFAULT NULL,
    created_at        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (email)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `users` (`username`, `email`, `password`, `gender`, `birthday`) VALUES
('ユーザーネーム2', 'user1@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'man', '2019-07-22'),
('hogehoge', 'update3@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22'),
('hogehoge', 'update4@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'man', '2019-07-22'),
('hogehoge', 'update5@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22'),
('hogehoge', 'update6@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'man', '2019-07-22'),
('hogehoge', 'update7@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'man', '2019-07-22'),
('hogehoge', 'update8@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'man', '2019-07-22'),
('hogehoge', 'update9@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22'),
('hogehoge', 'update10@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22'),
('hogehoge', 'update11@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22'),
('hogehoge', 'update12@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22'),
('hogehoge', 'update13@hoge.com', '$2a$10$aVbSd2rQ78sL/RUpqVnOL.yrJJXNs935N2Ev9TR4v1I3SYCRmsp4u', 'woman', '2019-07-22');


-- +goose Down
DROP TABLE users;
