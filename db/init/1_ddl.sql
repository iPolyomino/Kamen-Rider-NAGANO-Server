SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `imageboard` DEFAULT CHARACTER SET utf8mb4;
USE `imageboard`;

CREATE TABLE IF NOT EXISTS `imageboard`.`comment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `sender` VARCHAR(255) CHARACTER SET 'utf8mb4' NOT NULL,
  `text` VARCHAR(255) CHARACTER SET 'utf8mb4' NOT NULL,
  `image_url` VARCHAR(512) CHARACTER SET 'utf8mb4' NULL,
  `room_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_comment_room_idx` (`room_id` ASC) VISIBLE,
  CONSTRAINT `fk_comment_room`
    FOREIGN KEY (`room_id`)
    REFERENCES `imageboard`.`room` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `imageboard`.`room` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) CHARACTER SET 'utf8mb4' NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;