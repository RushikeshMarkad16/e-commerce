CREATE TABLE IF NOT EXISTS `order1`(
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_value` INT NOT NULL,
    `dispatch_date` DATE NOT NULL,
    `order_status` ENUM('Placed','Dispatched','Completed','Returned','Cancelled') NOT NULL
);