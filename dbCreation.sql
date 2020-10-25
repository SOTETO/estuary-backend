-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema estuary
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema estuary
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `estuary` DEFAULT CHARACTER SET utf8 ;
USE `estuary` ;

-- -----------------------------------------------------
-- Table `estuary`.`workshop`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`workshop` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `date` BIGINT(8) NOT NULL,
  `teaser` TEXT NULL,
  `locationName` VARCHAR(255) NOT NULL,
  `locationOnMap` VARCHAR(2048) NULL,
  PRIMARY KEY (`UUID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`tag`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`tag` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`UUID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`workshopTags`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`workshopTags` (
  `ID` INT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `workshop_UUID` VARCHAR(36) NOT NULL,
  `tags_UUID` VARCHAR(36) NOT NULL,
  INDEX `fk_workshopTags_workshop_idx` (`workshop_UUID` ASC) VISIBLE,
  INDEX `fk_workshopTags_tags1_idx` (`tags_UUID` ASC) VISIBLE,
  PRIMARY KEY (`ID`),
  CONSTRAINT `fk_workshopTags_workshop`
    FOREIGN KEY (`workshop_UUID`)
    REFERENCES `estuary`.`workshop` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_workshopTags_tags1`
    FOREIGN KEY (`tags_UUID`)
    REFERENCES `estuary`.`tag` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`Content`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`Content` (
  `UUID` VARCHAR(36) NOT NULL,
  `title` VARCHAR(255) NULL,
  `workshop_UUID` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`UUID`),
  INDEX `fk_Content_workshop1_idx` (`workshop_UUID` ASC) VISIBLE,
  CONSTRAINT `fk_Content_workshop1`
    FOREIGN KEY (`workshop_UUID`)
    REFERENCES `estuary`.`workshop` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`likes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`likes` (
  `ID` INT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_UUID` VARCHAR(36) NOT NULL,
  `Conent_ID` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`ID`),
  INDEX `fk_likes_ProblemStatement1_idx` (`Conent_ID` ASC) VISIBLE,
  CONSTRAINT `fk_likes_ProblemStatement1`
    FOREIGN KEY (`Conent_ID`)
    REFERENCES `estuary`.`Content` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`authors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`authors` (
  `ID` INT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `workshop_UUID` VARCHAR(36) NOT NULL,
  `user_UUID` VARCHAR(36) NOT NULL,
  `visible` TINYINT NOT NULL,
  INDEX `fk_authors_workshop1_idx` (`workshop_UUID` ASC) VISIBLE,
  PRIMARY KEY (`ID`),
  CONSTRAINT `fk_authors_workshop1`
    FOREIGN KEY (`workshop_UUID`)
    REFERENCES `estuary`.`workshop` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`linkTag`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`linkTag` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`UUID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`ContenttLink`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`ContenttLink` (
  `ID` INT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `linkTag_UUID` VARCHAR(36) NOT NULL,
  `Content1_UUID` VARCHAR(36) NOT NULL,
  `Content2_UUID` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`ID`),
  INDEX `fk_ProblemStatementLink_linkTag1_idx` (`linkTag_UUID` ASC) VISIBLE,
  INDEX `fk_ContenttLink_Content1_idx` (`Content1_UUID` ASC) VISIBLE,
  INDEX `fk_ContenttLink_Content2_idx` (`Content2_UUID` ASC) VISIBLE,
  CONSTRAINT `fk_ProblemStatementLink_linkTag1`
    FOREIGN KEY (`linkTag_UUID`)
    REFERENCES `estuary`.`linkTag` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ContenttLink_Content1`
    FOREIGN KEY (`Content1_UUID`)
    REFERENCES `estuary`.`Content` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ContenttLink_Content2`
    FOREIGN KEY (`Content2_UUID`)
    REFERENCES `estuary`.`Content` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`ProblemStatement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`ProblemStatement` (
  `Content_UUID` VARCHAR(36) NOT NULL,
  `iAM` TEXT NULL,
  `iWant` TEXT NULL,
  `but` TEXT NULL,
  `because` TEXT NULL,
  `feel` TEXT NULL,
  INDEX `fk_ProblemStatement_Content1_idx` (`Content_UUID` ASC) VISIBLE,
  PRIMARY KEY (`Content_UUID`),
  CONSTRAINT `fk_ProblemStatement_Content1`
    FOREIGN KEY (`Content_UUID`)
    REFERENCES `estuary`.`Content` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
