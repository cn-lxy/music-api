CREATE DATABASE  `music`;

USE `music`;

# create user table
CREATE TABLE `users` (
    `id` INT UNSIGNED AUTO_INCREMENT,
    `nick_name` VARCHAR(64) NOT NULL,
    `email` VARCHAR(32) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`id`)
);

# playlist table(用户创建的)
CREATE TABLE `playlist` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `create_user_id` INT UNSIGNED,
    `create_time` TIMESTAMP DEFAULT NOW(),
    `update_time` TIMESTAMP DEFAULT NOW(),
    `play_count` INT UNSIGNED DEFAULT 0,
    PRIMARY KEY (`id`),
    CONSTRAINT FOREIGN KEY (`create_user_id`) REFERENCES `users`(`id`)
);