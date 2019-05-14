# Dump of table chapter_unlock
# ------------------------------------------------------------
# Length of unique index openid is long. But it is below 767.

DROP TABLE IF EXISTS `chapter_unlock`;

CREATE TABLE `chapter_unlock`(
    id     BIGINT UNSIGNED AUTO_INCREMENT
        PRIMARY KEY,
    mobile VARCHAR(13) NOT NULL COMMENT '电话',
    openid VARCHAR(63) DEFAULT '' NOT NULL,
    status TINYINT(1) UNSIGNED DEFAULT 2  NULL COMMENT ' 1已绑定 2未绑定 默认 2',
    CONSTRAINT chapter_unlock_openid_uindex
        UNIQUE (openid)
)
 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;