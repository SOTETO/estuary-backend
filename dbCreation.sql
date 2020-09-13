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
  `UUID` VARCHAR(255) NOT NULL,
  `username` VARCHAR(16) NOT NULL,
  PRIMARY KEY (`UUID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`workshop`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`workshop` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `type` INT NOT NULL,
  `status` INT NOT NULL,
  `date` INT NOT NULL,
  `teaser` VARCHAR(1023) NULL,
  `locationName` VARCHAR(255) NOT NULL,
  `locationOnMap` VARCHAR(255) NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`tag`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`tag` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`workshopTags`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`workshopTags` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `workshop_ID` INT NOT NULL,
  `tags_ID` INT NOT NULL,
  INDEX `fk_workshopTags_workshop_idx` (`workshop_ID` ASC) VISIBLE,
  INDEX `fk_workshopTags_tags1_idx` (`tags_ID` ASC) VISIBLE,
  PRIMARY KEY (`ID`),
  CONSTRAINT `fk_workshopTags_workshop`
    FOREIGN KEY (`workshop_ID`)
    REFERENCES `estuary`.`workshop` (`ID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_workshopTags_tags1`
    FOREIGN KEY (`tags_ID`)
    REFERENCES `estuary`.`tag` (`ID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`ProblemStatement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`ProblemStatement` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `iAM` VARCHAR(255) NULL,
  `iWant` VARCHAR(1023) NULL,
  `but` VARCHAR(1023) NULL,
  `because` VARCHAR(1023) NULL,
  `feel` VARCHAR(1023) NULL,
  `workshop_ID` INT NOT NULL,
  PRIMARY KEY (`ID`),
  INDEX `fk_ProblemStatement_workshop1_idx` (`workshop_ID` ASC) VISIBLE,
  CONSTRAINT `fk_ProblemStatement_workshop1`
    FOREIGN KEY (`workshop_ID`)
    REFERENCES `estuary`.`workshop` (`ID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`likes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`likes` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `user_UUID` VARCHAR(255) NOT NULL,
  `ProblemStatement_ID` INT NOT NULL,
  INDEX `fk_likes_user1_idx` (`user_UUID` ASC) VISIBLE,
  PRIMARY KEY (`ID`),
  INDEX `fk_likes_ProblemStatement1_idx` (`ProblemStatement_ID` ASC) VISIBLE,
  CONSTRAINT `fk_likes_user1`
    FOREIGN KEY (`user_UUID`)
    REFERENCES `estuary`.`user` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_likes_ProblemStatement1`
    FOREIGN KEY (`ProblemStatement_ID`)
    REFERENCES `estuary`.`ProblemStatement` (`ID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`authors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`authors` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `workshop_ID` INT NOT NULL,
  `user_UUID` VARCHAR(255) NOT NULL,
  `visible` TINYINT NOT NULL,
  INDEX `fk_authors_workshop1_idx` (`workshop_ID` ASC) VISIBLE,
  INDEX `fk_authors_user1_idx` (`user_UUID` ASC) VISIBLE,
  PRIMARY KEY (`ID`),
  CONSTRAINT `fk_authors_workshop1`
    FOREIGN KEY (`workshop_ID`)
    REFERENCES `estuary`.`workshop` (`ID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_authors_user1`
    FOREIGN KEY (`user_UUID`)
    REFERENCES `estuary`.`user` (`UUID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `estuary`.`ProblemStatementLink`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `estuary`.`ProblemStatementLink` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `tag` VARCHAR(45) NULL,
  `ProblemStatement1_ID` INT NOT NULL,
  `ProblemStatement2_ID` INT NOT NULL,
  PRIMARY KEY (`ID`),
  INDEX `fk_ProblemStatementLink_ProblemStatement1_idx` (`ProblemStatement1_ID` ASC) VISIBLE,
  INDEX `fk_ProblemStatementLink_ProblemStatement2_idx` (`ProblemStatement2_ID` ASC) VISIBLE,
  CONSTRAINT `fk_ProblemStatementLink_ProblemStatement1`
    FOREIGN KEY (`ProblemStatement1_ID`)
    REFERENCES `estuary`.`ProblemStatement` (`ID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ProblemStatementLink_ProblemStatement2`
    FOREIGN KEY (`ProblemStatement2_ID`)
    REFERENCES `estuary`.`ProblemStatement` (`ID`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
