CREATE TABLE `menu_children` (
  `id` int(11) NOT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `route` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `menu_children_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `menus` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `menu_children` (`id`, `parent_id`, `name`, `route`) VALUES (1, 1, 'สินทรัพย์', '/dashboards/asset');
INSERT INTO `menu_children` (`id`, `parent_id`, `name`, `route`) VALUES (2, 1, 'สภาพทาง', '/dashboards/condition');