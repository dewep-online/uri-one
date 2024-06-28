CREATE TABLE `urls_named` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `urls_id` bigint(20) unsigned NOT NULL,
  `name` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `urls_id` (`urls_id`),
  CONSTRAINT `urls_named_ibfk_1` FOREIGN KEY (`urls_id`) REFERENCES `urls` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;