-- SQL dump generated using DBML (dbml-lang.org)
-- Database: MySQL
-- Generated at: 2023-03-06T04:07:49.897Z

CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` varchar(20) UNIQUE NOT NULL,
  `password` varchar(200) NOT NULL,
  `role` varchar(20) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
