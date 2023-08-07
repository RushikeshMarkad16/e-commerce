CREATE TABLE IF NOT EXISTS `order1`(
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `amount` INT NOT NULL,
    `discount_perc` INT NOT NULL,
    `final_amount` INT NOT NULL,
    `dispatch_date` DATE ,
    `order_status` ENUM('Placed','Dispatched','Completed','Returned','Cancelled') NOT NULL
);