CREATE TABLE `beer_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(250) NOT NULL,
  `type_description` text NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `beer_type` (`id`, `type_name`, `type_description`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'Lager Beer (เบียร์ลาเกอร์)', 'Lager Beer (เบียร์ลาเกอร์) เป็นเบียร์ที่ผ่านการหมักแบบสต์จมอยู่ที่กันภาชนะ เมื่อหมักเสร็จเรียบร้อย ยีสต์จะจมอยู่ที่ด้านล่างของเบียร์ หรือที่เรียกว่า Bottom-Fermenting Yeast ครับ โดยจะใช้อุณหภูมิในการหมักอยู่ที่ 7-12 องศาเซลเซียส ซึ่งเป็นอุณหภูมิที่ค่อนข้างต่ำ และถ้าจะดื่มก็ต้องรอนานกันหน่อยนะ เพราะเบียร์ลาเกอร์จะใช้เวลาหมักนานถึง 2 - 3 เดือนเลยทีเดียว', NULL, NULL, NULL);
INSERT INTO `beer_type` (`id`, `type_name`, `type_description`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'Ale Beer (เบียร์เอล)', 'Ale Beer (เบียร์เอล) เป็นเบียร์ที่ใช้การหมักแบบยีสต์ลอยที่ผิวน้ำ หรือ Top-Fermenting Yeast ยีสต์มันจะทำงานที่ด้านบนของเบียร์ หรือพูดให้เห็นภาพง่าย ๆ มันจะลอยตัวอยู่บนผิวน้ำเมื่อหมักเสร็จแล้วนั่นแหละครับ โดยการทำเบียร์เอลจะใช้อุณหภูมิ 18- 24 องศาเซลเซียส และใช้เวลาในการหมักเบียร์ประมาณ 3 - 4 สัปดาห์เท่านั้น', NULL, NULL, NULL);