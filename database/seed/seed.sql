# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.32)
# Database: todo
# Generation Time: 2021-03-06 17:43:50 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table friends
# ------------------------------------------------------------

DROP TABLE IF EXISTS `friends`;

CREATE TABLE `friends` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `user_friends` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `friends` WRITE;
/*!40000 ALTER TABLE `friends` DISABLE KEYS */;

INSERT INTO `friends` (`id`, `user_id`, `user_friends`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,1,'[2]','2021-03-07 01:02:29',NULL,NULL),
	(2,2,'[1]','2021-03-07 01:02:36',NULL,NULL),
	(3,3,'[]','2021-03-07 01:28:53',NULL,NULL);

/*!40000 ALTER TABLE `friends` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table todos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `todos`;

CREATE TABLE `todos` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `status` varchar(64) NOT NULL,
  `priority` int(11) NOT NULL,
  `is_private` tinyint(1) DEFAULT '1',
  `due_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(128) NOT NULL DEFAULT '',
  `updated_by` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `todos` WRITE;
/*!40000 ALTER TABLE `todos` DISABLE KEYS */;

INSERT INTO `todos` (`id`, `name`, `description`, `is_deleted`, `status`, `priority`, `is_private`, `due_date`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`)
VALUES
	(1,'Testing','testing',0,'in-progress',1,0,'2021-03-12 15:22:41','2021-03-06 17:02:55','2021-03-06 17:02:55',NULL,'',''),
	(2,'Testing2','testing2',0,'in-progress',1,0,'2021-03-12 15:22:41','2021-03-06 17:03:02','2021-03-06 17:03:02',NULL,'',''),
	(3,'Testing2','testing2',0,'in-progress',1,1,'2021-03-12 15:22:41','2021-03-06 17:03:12','2021-03-06 17:03:12',NULL,'',''),
	(4,'Testing3','testing3',0,'in-progress',2,0,'2021-03-07 01:27:14','2021-03-06 17:26:15','2021-03-07 01:27:14',NULL,'','');

/*!40000 ALTER TABLE `todos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_todos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_todos`;

CREATE TABLE `user_todos` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `todo_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user_todos` WRITE;
/*!40000 ALTER TABLE `user_todos` DISABLE KEYS */;

INSERT INTO `user_todos` (`id`, `user_id`, `todo_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,1,1,'2021-03-06 17:02:55','2021-03-06 17:02:55',NULL),
	(2,1,2,'2021-03-06 17:03:02','2021-03-06 17:03:02',NULL),
	(3,2,3,'2021-03-06 17:03:12','2021-03-06 17:03:12',NULL),
	(4,2,4,'2021-03-06 17:26:15','2021-03-06 17:26:15',NULL);

/*!40000 ALTER TABLE `user_todos` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'user1','2021-03-07 01:02:14',NULL,NULL),
	(2,'user2','2021-03-07 01:02:18',NULL,NULL),
	(3,'user3','2021-03-07 01:28:46',NULL,NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
