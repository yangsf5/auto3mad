CREATE DATABASE auto3mad;

USE auto3mad;

CREATE TABLE `daily_time_routine` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `short_name` VARCHAR(32) NOT NULL DEFAULT '',
    `event_scope` VARCHAR(256) NOT NULL,
    `will_spend` INT NOT NULL,
    `history_spend` DOUBLE DEFAULT '0',
    `icon` VARCHAR(64) DEFAULT '',
    `sort` INT DEFAULT NULL,
    `object` INT DEFAULT '0',
    `object_unit` CHAR(10) DEFAULT '',
    `progress` INT DEFAULT '0',
    `start_date` DATE NOT NULL,
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
    `id` INT NOT NULL AUTO_INCREMENT,
    `date` DATE NOT NULL,
    `desc` VARCHAR(128) NOT NULL,
    `remind_type` INT NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `url_group` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `desc` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `url_item` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `icon` VARCHAR(256) NOT NULL,
    `url` VARCHAR(256) NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `group_id` INT DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;