CREATE TABLE `shorten_url`
(
    `id`            bigint unsigned                                                NOT NULL AUTO_INCREMENT,
    `source`        varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `domain`        varchar(253) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL,
    `hash`          varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci   NOT NULL,
    `lock`          tinyint unsigned                                               NOT NULL,
    `request_count` bigint unsigned                                                NOT NULL,
    `created_at`    timestamp                                                      NOT NULL,
    `updated_at`    timestamp                                                      NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash` (`hash`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;