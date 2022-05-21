CREATE DATABASE auto3mad;

USE auto3mad;

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

CREATE TABLE `daily_time_use` (
	`start_time` INT unsigned NOT NULL,
	`end_time` INT unsigned NOT NULL,
	`routine_id` INT NOT NULL,
	`date` DATE NOT NULL,
	`month` CHAR(7) NOT NULL,
	PRIMARY KEY (`start_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `day_memorial` (
	`id` int NOT NULL AUTO_INCREMENT,
	`date` date NOT NULL,
	`desc` varchar(128) NOT NULL,
	`remind_type` int NOT NULL DEFAULT '1',
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `url_group` (
	`id` int NOT NULL AUTO_INCREMENT,
	`desc` varchar(64) NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `url_item` (
	`id` int NOT NULL AUTO_INCREMENT,
	`icon` varchar(256) NOT NULL,
	`url` varchar(256) NOT NULL,
	`title` varchar(64) NOT NULL,
	`group_id` int DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;