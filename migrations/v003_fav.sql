CREATE TABLE `fav` (
    `id`         INT         NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email`      VARCHAR(32) NOT NULL,
    `surah`      SMALLINT    NOT NULL,
    `verse`      SMALLINT    NOT NULL,
    `created_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP   NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT UNIQUE INDEX (`email`, `surah`, `verse`)
);
