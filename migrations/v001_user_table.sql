CREATE TABLE `user` (
    `id`         INT          NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email`      VARCHAR(32)  NOT NULL,
    `name`       VARCHAR(32)  NOT NULL,
    `token`      VARCHAR(512) NOT NULL,
    `picture`    VARCHAR(128) NOT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT UNIQUE INDEX (`email`),
    INDEX (`token`)
);
