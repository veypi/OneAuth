/*
 * table.sql
 * Copyright (C) 2022 veypi <i@veypi.com>
 *
 * Distributed under terms of the Apache license.
 */

CREATE TABLE IF NOT EXISTS `user`
(
    `id`            varchar(32)  NOT NULL DEFAULT '' COMMENT 'User UUID',
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `username`      varchar(255) NOT NULL UNIQUE,
    `nickname`      varchar(255),
    `email`         varchar(255) UNIQUE,
    `phone`         varchar(255) UNIQUE,
    `icon`          varchar(255),
    `_real_code`     varchar(32),
    `_check_code`    binary(48),

    `status`        int NOT NULL COMMENT '状态（0：ok，1：disabled）' DEFAULT 0,
    `used`          int NOT NULL DEFAULT 0,
    `space`         int NOT NULL DEFAULT 300,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `app`
(
    `id`            varchar(32)  NOT NULL,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `_key`           varchar(32) NOT NULL,
    `name`          varchar(255) NOT NULL,
    `icon`          varchar(255),
    `des`           varchar(255),
    `user_count`    int NOT NULL DEFAULT 0,
    `hide`          tinyint(1) NOT NULL DEFAULT 0,
    `join_method`   int NOT NULL DEFAULT 0,

    `role_id`       varchar(32),
    `redirect`      varchar(255),
    `status`        int NOT NULL COMMENT '状态（0：ok，1：disabled）' DEFAULT 0,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `app_user`
(
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `app_id`        varchar(32) NOT NULL,
    `user_id`       varchar(32) NOT NULL,
    `status`        int NOT NULL DEFAULT 0 COMMENT '0: ok,1:disabled,2:applying,3:deny',

    PRIMARY KEY (`user_id`,`app_id`) USING BTREE,
    FOREIGN KEY (`app_id`) REFERENCES `app`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8  COMMENT='app_to_user';



CREATE TABLE IF NOT EXISTS `role`
(
    `id`            varchar(32)  NOT NULL,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `app_id`        varchar(32) NOT NULL,

    `name`          varchar(255) NOT NULL,
    `des`           varchar(255),
    `user_count`    int NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`) USING BTREE,
    FOREIGN KEY (`app_id`) REFERENCES `app`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `user_role`
(
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `user_id`       varchar(32) NOT NULL,
    `role_id`        varchar(32) NOT NULL,
    `status`        varchar(32) NOT NULL DEFAULT 0,

    PRIMARY KEY (`user_id`,`role_id`) USING BTREE,
    FOREIGN KEY (`role_id`) REFERENCES `role`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `resource`
(
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `app_id`        varchar(32) NOT NULL,
    `name`          varchar(32) NOT NULL,
    `des`           varchar(255),


    PRIMARY KEY (`app_id`,`name`) USING BTREE,
    FOREIGN KEY (`app_id`) REFERENCES `app`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `access`
(
    `id`            int NOT NULL AUTO_INCREMENT,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `app_id`        varchar(32) NOT NULL,
    `name`          varchar(32) NOT NULL,

    `role_id`       varchar(32) NULL DEFAULT NULL,
    `user_id`       varchar(32) NULL DEFAULT NULL,
    `rid`           varchar(32) DEFAULT NULL COMMENT '资源子id',
    `level`         int NOT NULL DEFAULT 0,

    -- PRIMARY KEY (`app_id`,`name`, `role_id`, `user_id`) USING BTREE,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`role_id`) REFERENCES `role`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`app_id`,`name`) REFERENCES `resource`(`app_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



INSERT INTO `app` (`id`, `name`, `_key`, `role_id`)
VALUES ('FR9P5t8debxc11aFF', 'oa', 'AMpjwQHwVjGsb1WC4WG6', '1lytMwQL4uiNd0vsc');

INSERT INTO `resource` (`app_id`, `name`)
VALUES ('FR9P5t8debxc11aFF', 'app'),
('FR9P5t8debxc11aFF', 'user');

INSERT INTO `role` (`id`, `app_id`, `name`)
VALUES ('1lytMwQL4uiNd0vsc', 'FR9P5t8debxc11aFF', 'admin');

INSERT INTO `access` (`app_id`, `name`, `role_id`, `user_id`,`level`)
VALUES ('FR9P5t8debxc11aFF', 'app', '1lytMwQL4uiNd0vsc', NULL,5),
('FR9P5t8debxc11aFF', 'user', '1lytMwQL4uiNd0vsc', NULL,5);



ALTER TABLE `app`
ADD FOREIGN KEY (`role_id`) REFERENCES `role`(`id`);
