/*
 * table.sql
 * Copyright (C) 2022 veypi <i@veypi.com>
 *
 * Distributed under terms of the Apache license.
 */

CREATE TABLE IF NOT EXISTS `user`
(
    `id`            varchar(32)  NOT NULL DEFAULT '' COMMENT 'User UUID',
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_flag`   int(1) NOT NULL,

    `username`      varchar(255) NOT NULL UNIQUE,
    `nickname`      varchar(255),
    `email`         varchar(255) UNIQUE,
    `phone`         varchar(255) UNIQUE,
    `icon`          varchar(255),

    `status`        int NOT NULL COMMENT '状态（0：ok，1：disabled）',
    `used`          int,
    `space`         int DEFAULT 300,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `app`
(
    `id`            varchar(32)  NOT NULL,
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_flag`   int(1) NOT NULL,

    `name`          varchar(255) NOT NULL,
    `icon`          varchar(255),
    `des`           varchar(255),
    `user_count`    int NOT NULL,
    `hide`          int(1) NOT NULL,
    `regist`        int(1) NOT NULL,

    `role_id`       varchar(32),
    `redirect`      varchar(255),
    `status`        int NOT NULL COMMENT '状态（0：ok，1：disabled）',

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `app_user`
(
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `app_id`        varchar(32) NOT NULL,
    `user_id`       varchar(32) NOT NULL,
    `status`        varchar(32) NOT NULL,

    PRIMARY KEY (`user_id`,`app_id`) USING BTREE,
    FOREIGN KEY (`app_id`) REFERENCES `app`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8  COMMENT='app_to_user';



CREATE TABLE IF NOT EXISTS `role`
(
    `id`            varchar(32)  NOT NULL,
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_flag`   int(1) NOT NULL,
    `app_id`        varchar(32) NOT NULL,

    `name`          varchar(255) NOT NULL,
    `des`           varchar(255),
    `user_count`    int NOT NULL,

    PRIMARY KEY (`id`) USING BTREE,
    FOREIGN KEY (`app_id`) REFERENCES `app`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `user_role`
(
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `user_id`       varchar(32) NOT NULL,
    `role_id`        varchar(32) NOT NULL,
    `status`        varchar(32) NOT NULL,

    PRIMARY KEY (`user_id`,`role_id`) USING BTREE,
    FOREIGN KEY (`role_id`) REFERENCES `role`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `resource`
(
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_flag`   int(1) NOT NULL,

    `app_id`        varchar(32) NOT NULL,
    `name`          varchar(32) NOT NULL,
    `des`           varchar(255),


    PRIMARY KEY (`app_id`,`name`) USING BTREE,
    FOREIGN KEY (`app_id`) REFERENCES `app`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `access`
(
    `created`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_flag`   int(1) NOT NULL,

    `app_id`        varchar(32) NOT NULL,
    `name`          varchar(32) NOT NULL,

    `role_id`       varchar(32) NULL DEFAULT NULL,
    `user_id`       varchar(32) NULL DEFAULT NULL,
    `rid`           varchar(32) DEFAULT NULL COMMENT '资源子id',
    `level`         int DEFAULT 0,

    -- PRIMARY KEY (`app_id`,`name`, `role_id`, `user_id`) USING BTREE,
    FOREIGN KEY (`role_id`) REFERENCES `role`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`app_id`,`name`) REFERENCES `resource`(`app_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `app`
ADD FOREIGN KEY (`role_id`) REFERENCES `role`(`id`);

INSERT INTO `user` (`id`, `username`)
VALUES ('user0', 'name');

INSERT INTO `app` (`id`, `name`)
VALUES ('app0', 'a0');

INSERT INTO `resource` (`app_id`, `name`)
VALUES ('app0', 'sound');

INSERT INTO `role` (`id`, `app_id`)
VALUES ('role0', 'app0');

INSERT INTO `access` (`app_id`, `name`, `role_id`, `user_id`)
VALUES ('app0', 'sound', NULL, 'user0');
