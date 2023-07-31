CREATE TABLE IF NOT EXISTS `order_item`(
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `quantity` INT NOT NULL,
    `item_value` INT NOT NULL,
    FOREIGN KEY(`product_id`) REFERENCES `product`(`id`) ON DELETE CASCADE,
    FOREIGN KEY(`order_id`) REFERENCES `order1`(`id`) ON DELETE CASCADE
);