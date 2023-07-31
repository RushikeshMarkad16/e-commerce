CREATE TABLE IF NOT EXISTS `product`(
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL UNIQUE,
    `availability` INT NOT NULL,
    `price` INT NOT NULL,
    `category` ENUM('Premium','Regular','Budget') NOT NULL
);