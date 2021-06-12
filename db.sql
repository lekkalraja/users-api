CREATE DATABASE users_db;

USE users_db;

CREATE TABLE users_db.users
(
  id              INT unsigned NOT NULL AUTO_INCREMENT,
  first_name      VARCHAR(150) NOT NULL,
  last_name       VARCHAR(150) NOT NULL,
  email_id        VARCHAR(150) NOT NULL,
  date_created    VARCHAR(150) NOT NULL,
  PRIMARY KEY     (id),
  UNIQUE KEY unique_email (email_id)
);