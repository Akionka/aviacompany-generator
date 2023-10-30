CREATE SCHEMA IF NOT EXISTS `aircompany` DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
USE `aircompany`;


CREATE TABLE IF NOT EXISTS `airport` (
  `iata_code` CHAR(3) NOT NULL,
  `city` VARCHAR(64) NOT NULL,
  `timezone` VARCHAR(64) NOT NULL,
  PRIMARY KEY (`iata_code`));

  CREATE TABLE IF NOT EXISTS `attendant` (
  `id` INT(11) NOT NULL,
  `given_name` VARCHAR(64) NOT NULL,
  `last_name` VARCHAR(64) NOT NULL,
  PRIMARY KEY (`id`));

  CREATE TABLE IF NOT EXISTS `booking_office` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(256) NOT NULL,
  `phone_number` CHAR(11) NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE IF NOT EXISTS `cashier` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `login` VARCHAR(32) NOT NULL,
  `last_name` VARCHAR(64) NOT NULL,
  `first_name` VARCHAR(64) NOT NULL,
  `middle_name` VARCHAR(64) NULL,
  `password` BINARY(60) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `login` (`login` ASC) VISIBLE);

CREATE TABLE IF NOT EXISTS `line` (
  `line_code` VARCHAR(6) NOT NULL,
  `dep_time` TIME NOT NULL,
  `flight_time` TIME NOT NULL,
  `base_price` DECIMAL(14,4) NOT NULL,
  `dep_airport` CHAR(3) NOT NULL,
  `arr_airport` CHAR(3) NOT NULL,
  PRIMARY KEY (`line_code`),
  INDEX `dep_airport` (`dep_airport` ASC) VISIBLE,
  INDEX `arr_airport` (`arr_airport` ASC) VISIBLE,
  CONSTRAINT `line_ibfk_1` FOREIGN KEY (`dep_airport`) REFERENCES `airport` (`iata_code`),
  CONSTRAINT `line_ibfk_2` FOREIGN KEY (`arr_airport`) REFERENCES `airport` (`iata_code`));

    CREATE TABLE IF NOT EXISTS `plane_model` (
  `icao_type_des` CHAR(4) NOT NULL,
  `manufacturer` VARCHAR(32) NOT NULL,
  `model` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`icao_type_des`));

  CREATE TABLE IF NOT EXISTS `plane` (
  `iata_code` VARCHAR(7) NOT NULL,
  `model_code` CHAR(4) NOT NULL,
  PRIMARY KEY (`iata_code`),
  INDEX `model_code` (`model_code` ASC) VISIBLE,
  CONSTRAINT `liner_ibfk_1` FOREIGN KEY (`model_code`) REFERENCES `plane_model` (`icao_type_des`));

    CREATE TABLE IF NOT EXISTS `pilot` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `given_name` VARCHAR(128) NOT NULL,
  `last_name` VARCHAR(64) NOT NULL,
  PRIMARY KEY (`id`));

  CREATE TABLE IF NOT EXISTS `flight` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `dep_date` DATE NOT NULL,
  `act_dep_time` DATETIME NULL,
  `act_arr_time` DATETIME NULL,
  `line_code` VARCHAR(6) NOT NULL,
  `plane_code` VARCHAR(7) NOT NULL,
  `pilot_id` INT(11) NOT NULL,
  `copilot_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `dep_date` (`dep_date` ASC, `line_code` ASC) VISIBLE,
  INDEX `line_code` (`line_code` ASC) VISIBLE,
  INDEX `plane_code` (`plane_code` ASC) VISIBLE,
  INDEX `fk_flight_pilot_idx` (`pilot_id` ASC) VISIBLE,
  INDEX `fk_flight_copilot_idx` (`copilot_id` ASC) INVISIBLE,
  CONSTRAINT `flight_ibfk_1` FOREIGN KEY (`line_code`) REFERENCES `line` (`line_code`),
  CONSTRAINT `flight_ibfk_2` FOREIGN KEY (`plane_code`) REFERENCES `plane` (`iata_code`),
  CONSTRAINT `fk_flight_pilot1` FOREIGN KEY (`pilot_id`) REFERENCES `pilot` (`id`),
  CONSTRAINT `fk_flight_pilot2` FOREIGN KEY (`copilot_id`) REFERENCES `pilot` (`id`));

CREATE TABLE IF NOT EXISTS `seat` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `number` VARCHAR(3) NOT NULL,
  `class` ENUM('J', 'W', 'Y') NOT NULL,
  `model_code` CHAR(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `number` (`number` ASC, `model_code` ASC) VISIBLE,
  INDEX `model_code` (`model_code` ASC) VISIBLE,
  CONSTRAINT `seat_ibfk_1` FOREIGN KEY (`model_code`) REFERENCES `plane_model` (`icao_type_des`));

CREATE TABLE IF NOT EXISTS `tariff` (
  `id` INT NOT NULL,
  `name` VARCHAR(32) NOT NULL,
  `is_avaliable` TINYINT NOT NULL,
  `description` TEXT NULL,
  `multiplier` DOUBLE NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`));

CREATE TABLE IF NOT EXISTS `purchase` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `date` TIMESTAMP NOT NULL,
  `booking_office_id` INT(11) NOT NULL,
  `total_price` DECIMAL(14,4) NOT NULL,
  `contact_phone` CHAR(11) NOT NULL,
  `contact_email` VARCHAR(64) NOT NULL,
  `cashier_id` INT(11) NOT NULL,
  `tariff_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `booking_office_id` (`booking_office_id` ASC) VISIBLE,
  INDEX `cashier_id` (`cashier_id` ASC) VISIBLE,
  INDEX `fk_purchase_tariff1_idx` (`tariff_id` ASC) VISIBLE,
  CONSTRAINT `purchase_ibfk_1` FOREIGN KEY (`booking_office_id`) REFERENCES `booking_office` (`id`),
  CONSTRAINT `purchase_ibfk_2` FOREIGN KEY (`cashier_id`) REFERENCES `cashier` (`id`),
  CONSTRAINT `fk_purchase_tariff1` FOREIGN KEY (`tariff_id`) REFERENCES `tariff` (`id`));

CREATE TABLE IF NOT EXISTS `passenger` (
  `id` INT NOT NULL,
  `last_name` VARCHAR(64) NOT NULL,
  `given_name` VARCHAR(128) NOT NULL,
  `sex` TINYINT(1) NOT NULL,
  `passport_number` CHAR(10) NOT NULL,
  `birth_date` DATE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `last_name` (`last_name` ASC) VISIBLE);

CREATE TABLE IF NOT EXISTS `ticket` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `purchase_id` INT(11) NOT NULL,
  `passenger_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `purchase_id` (`purchase_id` ASC) VISIBLE,
  INDEX `fk_ticket_passenger1_idx` (`passenger_id` ASC) VISIBLE,
  CONSTRAINT `ticket_ibfk_1` FOREIGN KEY (`purchase_id`) REFERENCES `purchase` (`id`),
  CONSTRAINT `fk_ticket_passenger1` FOREIGN KEY (`passenger_id`) REFERENCES `passenger` (`id`));

CREATE TABLE IF NOT EXISTS `flight_in_ticket` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `flight_id` INT(11) NOT NULL,
  `seat_id` INT(11) NOT NULL,
  `ticket_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `flight_id` (`flight_id` ASC, `seat_id` ASC) VISIBLE,
  INDEX `seat_id` (`seat_id` ASC) VISIBLE,
  INDEX `ticket_id` (`ticket_id` ASC) VISIBLE,
  CONSTRAINT `flight_in_ticket_ibfk_1` FOREIGN KEY (`flight_id`) REFERENCES `flight` (`id`),
  CONSTRAINT `flight_in_ticket_ibfk_2` FOREIGN KEY (`seat_id`) REFERENCES `seat` (`id`),
  CONSTRAINT `flight_in_ticket_ibfk_3` FOREIGN KEY (`ticket_id`) REFERENCES `ticket` (`id`));

CREATE TABLE IF NOT EXISTS `pilot_flies_plane_model` (
  `pilot_id` INT(11) NOT NULL,
  `model_code` CHAR(4) NOT NULL,
  PRIMARY KEY (`model_code`, `pilot_id`),
  INDEX `fk_pilot_flies_plane_model_plane_model_idx` (`model_code` ASC) VISIBLE,
  INDEX `fk_pilot_flies_plane_model_pilot_idx` (`pilot_id` ASC) VISIBLE,
  CONSTRAINT `fk_pilot_has_plane_model_pilot1` FOREIGN KEY (`pilot_id`) REFERENCES `pilot` (`id`),
  CONSTRAINT `fk_pilot_has_plane_model_plane_model1` FOREIGN KEY (`model_code`) REFERENCES `plane_model` (`icao_type_des`));

CREATE TABLE IF NOT EXISTS `flight_attendant` (
  `attendant_id` INT(11) NOT NULL,
  `flight_id` INT(11) NOT NULL,
  PRIMARY KEY (`attendant_id`, `flight_id`),
  INDEX `fk_attendant_has_flight_flight1_idx` (`flight_id` ASC) VISIBLE,
  INDEX `fk_attendant_has_flight_attendant1_idx` (`attendant_id` ASC) VISIBLE,
  CONSTRAINT `fk_attendant_has_flight_attendant1` FOREIGN KEY (`attendant_id`) REFERENCES `attendant` (`id`),
  CONSTRAINT `fk_attendant_has_flight_flight1` FOREIGN KEY (`flight_id`) REFERENCES `flight` (`id`));

