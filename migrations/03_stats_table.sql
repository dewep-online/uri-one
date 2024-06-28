CREATE TABLE `stats` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `urls_id` bigint(20) unsigned NOT NULL,
  `month` int(11) NOT NULL,
  `count` bigint(20) NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `month` (`month`),
  KEY `urls_id` (`urls_id`),
  CONSTRAINT `stats_ibfk_1` FOREIGN KEY (`urls_id`) REFERENCES `urls` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;