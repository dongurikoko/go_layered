-- -----------------------------------------------------
-- Schema go_lesson1_api
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `go_lesson1_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `go_lesson1_api` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `go_lesson1_api`.`todo`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_lesson1_api`.`todo` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'TODO_ID',
  `title` VARCHAR(255) NOT NULL COMMENT 'タイトル名',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'TODO情報';
