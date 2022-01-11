CREATE TABLE `transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext,
  `last_name` longtext,
  `debit_amt` bigint DEFAULT NULL,
  `credit_amt` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_transactions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `transactions`
(`id`,
`created_at`,
`updated_at`,
`deleted_at`,
`first_name`,
`last_name`,
`debit_amt`,
`credit_amt`)
VALUES
(1,
'2022-01-02 11:38:30.360',
'2022-01-02 11:50:41.769',
NULL,
'Ruroni',
'Kenshin',
5000,
1000);
