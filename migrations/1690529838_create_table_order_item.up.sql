CREATE TABLE IF NOT EXISTS `order_item`(
    `id` VARCHAR(100) NOT NULL,
    `order_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `quantity` INT NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY(`product_id`) REFERENCES `product`(`id`) ON DELETE CASCADE,
    FOREIGN KEY(`order_id`) REFERENCES `order1`(`id`) ON DELETE CASCADE
);