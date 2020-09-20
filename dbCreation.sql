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
-- Table `estuary`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`user` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'external',
  PRIMARY KEY (`UUID`))
ENGINE = InnoDB;


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
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `workshop_ID` INT NOT NULL,
  `tags_ID` INT NOT NULL,
  INDEX `fk_workshopTags_workshop_idx` (`workshop_ID` ASC) VISIBLE,
  INDEX `fk_workshopTags_tags1_idx` (`tags_ID` ASC) VISIBLE,
  PRIMARY KEY (`UUID`),
  CONSTRAINT `fk_workshopTags_workshop`
    FOREIGN KEY (`workshop_ID`)
    REFERENCES `estuary`.`workshop` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_workshopTags_tags1`
    FOREIGN KEY (`tags_ID`)
    REFERENCES `estuary`.`tag` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`ProblemStatement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`ProblemStatement` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `iAM` TEXT NULL,
  `iWant` TEXT NULL,
  `but` TEXT NULL,
  `because` TEXT NULL,
  `feel` TEXT NULL,
  `workshop_ID` INT NOT NULL,
  PRIMARY KEY (`UUID`),
  INDEX `fk_ProblemStatement_workshop1_idx` (`workshop_ID` ASC) VISIBLE,
  CONSTRAINT `fk_ProblemStatement_workshop1`
    FOREIGN KEY (`workshop_ID`)
    REFERENCES `estuary`.`workshop` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`likes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`likes` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `user_UUID` VARCHAR(255) NOT NULL,
  `ProblemStatement_ID` INT NOT NULL,
  INDEX `fk_likes_user1_idx` (`user_UUID` ASC) VISIBLE,
  PRIMARY KEY (`UUID`),
  INDEX `fk_likes_ProblemStatement1_idx` (`ProblemStatement_ID` ASC) VISIBLE,
  CONSTRAINT `fk_likes_user1`
    FOREIGN KEY (`user_UUID`)
    REFERENCES `estuary`.`user` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_likes_ProblemStatement1`
    FOREIGN KEY (`ProblemStatement_ID`)
    REFERENCES `estuary`.`ProblemStatement` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`authors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`authors` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `workshop_ID` INT NOT NULL,
  `user_UUID` VARCHAR(255) NOT NULL,
  `visible` TINYINT NOT NULL,
  INDEX `fk_authors_workshop1_idx` (`workshop_ID` ASC) VISIBLE,
  INDEX `fk_authors_user1_idx` (`user_UUID` ASC) VISIBLE,
  PRIMARY KEY (`UUID`),
  CONSTRAINT `fk_authors_workshop1`
    FOREIGN KEY (`workshop_ID`)
    REFERENCES `estuary`.`workshop` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_authors_user1`
    FOREIGN KEY (`user_UUID`)
    REFERENCES `estuary`.`user` (`UUID`)
    ON DELETE NO ACTION
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
-- Table `estuary`.`ProblemStatementLink`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`ProblemStatementLink` (
  `UUID` VARCHAR(36) NOT NULL COMMENT 'UUID',
  `linkTag_ID` INT NOT NULL,
  `ProblemStatement1_ID` INT NOT NULL,
  `ProblemStatement2_ID` INT NOT NULL,
  PRIMARY KEY (`UUID`),
  INDEX `fk_ProblemStatementLink_ProblemStatement1_idx` (`ProblemStatement1_ID` ASC) VISIBLE,
  INDEX `fk_ProblemStatementLink_ProblemStatement2_idx` (`ProblemStatement2_ID` ASC) VISIBLE,
  INDEX `fk_ProblemStatementLink_linkTag1_idx` (`linkTag_ID` ASC) VISIBLE,
  CONSTRAINT `fk_ProblemStatementLink_ProblemStatement1`
    FOREIGN KEY (`ProblemStatement1_ID`)
    REFERENCES `estuary`.`ProblemStatement` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ProblemStatementLink_ProblemStatement2`
    FOREIGN KEY (`ProblemStatement2_ID`)
    REFERENCES `estuary`.`ProblemStatement` (`UUID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ProblemStatementLink_linkTag1`
    FOREIGN KEY (`linkTag_ID`)
    REFERENCES `estuary`.`linkTag` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
