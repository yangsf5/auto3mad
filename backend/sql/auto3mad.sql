-- MySQL dump 10.13  Distrib 8.0.23, for osx10.16 (x86_64)

-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */

/*!50503 SET NAMES utf8mb4 */

/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */

/*!40103 SET TIME_ZONE='+00:00' */

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */

/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */

/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */

--

-- Table structure for table `daily_time_routine`

--

/*!40101 SET @saved_cs_client     = @@character_set_client */

/*!50503 SET character_set_client = utf8mb4 */

CREATE TABLE `daily_time_routine` (
	`id` int NOT NULL AUTO_INCREMENT,
	`short_name` varchar(32) NOT NULL DEFAULT '',
	`event_scope` varchar(256) NOT NULL,
	`will_spend` int NOT NULL,
	`history_spend` double DEFAULT '0',
	`icon` varchar(64) DEFAULT '',
	`sort` int DEFAULT NULL,
	`object` int DEFAULT '0',
	`object_unit` char(10) DEFAULT '',
	`progress` int DEFAULT '0',
	`start_date` date NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

/*!40101 SET character_set_client = @saved_cs_client */

--

-- Table structure for table `daily_time_use`

--

/*!40101 SET @saved_cs_client     = @@character_set_client */

/*!50503 SET character_set_client = utf8mb4 */

CREATE TABLE `daily_time_use` (
	`start_time` int unsigned NOT NULL,
	`end_time` int unsigned NOT NULL,
	`routine_id` int NOT NULL,
	PRIMARY KEY (`start_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

/*!40101 SET character_set_client = @saved_cs_client */

--

-- Table structure for table `day_memorial`

--

/*!40101 SET @saved_cs_client     = @@character_set_client */

/*!50503 SET character_set_client = utf8mb4 */

CREATE TABLE `day_memorial` (
	`id` int NOT NULL AUTO_INCREMENT,
	`date` date NOT NULL,
	`desc` varchar(128) NOT NULL,
	`remind_type` int NOT NULL DEFAULT '1',
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

/*!40101 SET character_set_client = @saved_cs_client */

--

-- Table structure for table `url_group`

--

/*!40101 SET @saved_cs_client     = @@character_set_client */

/*!50503 SET character_set_client = utf8mb4 */

CREATE TABLE `url_group` (
	`id` int NOT NULL AUTO_INCREMENT,
	`desc` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

/*!40101 SET character_set_client = @saved_cs_client */

--

-- Table structure for table `url_item`

--

/*!40101 SET @saved_cs_client     = @@character_set_client */

/*!50503 SET character_set_client = utf8mb4 */

CREATE TABLE `url_item` (
	`id` int NOT NULL AUTO_INCREMENT,
	`icon` varchar(256) NOT NULL,
	`url` varchar(256) NOT NULL,
	`title` varchar(64) NOT NULL,
	`group_id` int DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

/*!40101 SET character_set_client = @saved_cs_client */

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */

/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */

/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */

/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */

/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */