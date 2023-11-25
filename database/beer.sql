CREATE TABLE `beer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `beer_name` varchar(250) NOT NULL,
  `beer_type_id` int(10) unsigned NOT NULL,
  `image` varchar(250) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `beer_type_id_fk` (`beer_type_id`),
  CONSTRAINT `beer_type_id_fk` FOREIGN KEY (`beer_type_id`) REFERENCES `beer_type` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;