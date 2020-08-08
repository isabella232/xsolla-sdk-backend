CREATE TABLE `sdk_db`.`users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(50) NOT NULL,
    `access_token` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
)