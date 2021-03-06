-- MySQL Script generated by MySQL Workbench
-- Sat Dec 30 10:40:00 2017
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema travel_senpai
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema travel_senpai
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `travel_senpai` DEFAULT CHARACTER SET utf8 ;
USE `travel_senpai` ;

-- -----------------------------------------------------
-- Table `travel_senpai`.`destination`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `travel_senpai`.`destination` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `budget` INT NULL,
  `country` VARCHAR(45) NULL,
  `city` VARCHAR(45) NULL,
  `currency` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `travel_senpai`.`transport_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `travel_senpai`.`transport_details` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `departure_date` TIMESTAMP(6) NULL,
  `arrival_date` TIMESTAMP(6) NULL,
  `layover` INT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `travel_senpai`.`transport`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `travel_senpai`.`transport` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(100) NULL,
  `cost` INT NULL DEFAULT 0,
  `provider` VARCHAR(100) NULL,
  `destination_id` INT NOT NULL,
  `outbound_id` INT NOT NULL,
  `return_id` INT NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_transport_destination1_idx` (`destination_id` ASC),
  INDEX `fk_transport_transport_details1_idx` (`outbound_id` ASC),
  INDEX `fk_transport_transport_details2_idx` (`return_id` ASC),
  CONSTRAINT `fk_transport_destination1`
    FOREIGN KEY (`destination_id`)
    REFERENCES `travel_senpai`.`destination` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_transport_transport_details1`
    FOREIGN KEY (`outbound_id`)
    REFERENCES `travel_senpai`.`transport_details` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_transport_transport_details2`
    FOREIGN KEY (`return_id`)
    REFERENCES `travel_senpai`.`transport_details` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `travel_senpai`.`accommodation`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `travel_senpai`.`accommodation` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL,
  `area` VARCHAR(100) NULL,
  `cost` INT NULL DEFAULT 0,
  `checkin_date` TIMESTAMP(6) NULL,
  `checkout_date` TIMESTAMP(6) NULL,
  `destination_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_accommodation_destination1_idx` (`destination_id` ASC),
  CONSTRAINT `fk_accommodation_destination1`
    FOREIGN KEY (`destination_id`)
    REFERENCES `travel_senpai`.`destination` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `travel_senpai`.`activity`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `travel_senpai`.`activity` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL,
  `cost` INT NULL DEFAULT 0,
  `provider` VARCHAR(100) NULL,
  `activity_date` TIMESTAMP(6) NULL,
  `estimated_time` INT NULL,
  `destination_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_activity_destination1_idx` (`destination_id` ASC),
  CONSTRAINT `fk_activity_destination1`
    FOREIGN KEY (`destination_id`)
    REFERENCES `travel_senpai`.`destination` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
