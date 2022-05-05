-- 短链表 lucky_market_url.sql

CREATE TABLE `lucky_market_url`
(
    `id`           bigint(20) unsigned NOT NULL COMMENT '哈希ID',
    `hash_key`     varchar(512) NOT NULL DEFAULT '' COMMENT '哈希后缀',
    `short_url`    varchar(512) NOT NULL DEFAULT '' COMMENT '短链接',
    `original_url` varchar(512) NOT NULL DEFAULT '' COMMENT '原链接',
    `salt`         varchar(512) NOT NULL DEFAULT '' COMMENT '加密盐',
    `status`       tinyint(4) NOT NULL DEFAULT '0' COMMENT '链接状态',
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_hashkey`(`hash_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='链接映射表';
