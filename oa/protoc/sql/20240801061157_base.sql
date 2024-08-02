CREATE TABLE IF NOT EXISTS `user`
(
    `id`            varchar(32)  NOT NULL DEFAULT '' COMMENT 'User UUID',
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,

    `username`      varchar(255) NOT NULL UNIQUE,
    `nickname`      varchar(255) NOT NULL,
    `email`         varchar(255) NOT NULL,
    `phone`         varchar(255) NOT NULL,
    `icon`          varchar(255) NOT NULL DEFAULT "",
    `_real_code`    varchar(32) NOT NULL,
    `_check_code`   varchar(64) NOT NULL,

    `status`        int NOT NULL COMMENT '状态（0：ok，1：disabled）' DEFAULT 0,
    `used`          int NOT NULL DEFAULT 0,
    `space`         int NOT NULL DEFAULT 300,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `app`
(
    `id`            varchar(32)  NOT NULL,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,

    `_key`           varchar(32) NOT NULL,
    `name`          varchar(255) NOT NULL,
    `icon`          varchar(255) NOT NULL DEFAULT "",
    `des`           TEXT,
    `user_count`    int NOT NULL DEFAULT 0,
    `hide`          tinyint(1) NOT NULL DEFAULT 0,
    `join_method`   int NOT NULL DEFAULT 0,

    `role_id`       varchar(32) NOT NULL,
    `host`          varchar(255) NOT NULL DEFAULT '',
    `redirect`      varchar(255) NOT NULL DEFAULT '',
    `status`        int NOT NULL COMMENT '状态（0：ok，1：disabled）' DEFAULT 0,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `app_user`
(
     id             int AUTO_INCREMENT,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,

    `app_id`        varchar(32) NOT NULL,
    `user_id`       varchar(32) NOT NULL,
    `status`        int NOT NULL DEFAULT 0 COMMENT '0: ok,1:disabled,2:applying,3:deny',
    PRIMARY KEY (`id`),
    unique index (`user_id`,`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8  COMMENT='app_to_user';

CREATE TABLE IF NOT EXISTS `role`
(
    `id`            varchar(32)  NOT NULL,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `app_id`        varchar(32) NOT NULL,

    `name`          varchar(255) NOT NULL,
    `des`           varchar(255) NOT NULL DEFAULT '',
    `user_count`    int NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `user_role`
(
     id             int AUTO_INCREMENT,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,

    `user_id`       varchar(32) NOT NULL,
    `role_id`       varchar(32) NOT NULL,
    `status`        varchar(32) NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`),
    UNIQUE INDEX (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `resource`
(
     id             int AUTO_INCREMENT,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,

    `app_id`        varchar(32) NOT NULL,
    `name`          varchar(32) NOT NULL,
    `des`           varchar(255) NOT NULL DEFAULT '',


    PRIMARY KEY (`id`),
    UNIQUE INDEX (`app_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `access`
(
    `id`            int NOT NULL AUTO_INCREMENT,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,

    `app_id`        varchar(32) NOT NULL,
    `access_id`     int NOT NULL DEFAULT 0,
    `name`          varchar(32) NOT NULL,

    `role_id`       varchar(32) NULL DEFAULT NULL,
    `user_id`       varchar(32) NULL DEFAULT NULL,
    `rid`           varchar(32) DEFAULT NULL COMMENT '资源子id',
    `level`         int NOT NULL DEFAULT 0,

    -- PRIMARY KEY (`app_id`,`name`, `role_id`, `user_id`) USING BTREE,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `token`
(
    `code`          varchar(32) NOT NULL,
    `created`       datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated`       datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `expired`       datetime NOT NULL,
    `client_id`     varchar(32) NOT NULL,
    `app_id`        varchar(32) NOT NULL,
    `user_id`       varchar(32) NOT NULL,
    `meta`          json,

    PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



INSERT INTO `app` (`id`, `name`, `_key`, `role_id` )
VALUES ('FR9P5t8debxc11aFF', 'oa', 'AMpjwQHwVjGsb1WC4WG6', '1lytMwQL4uiNd0vsc');

INSERT INTO `resource` (`app_id`, `name`)
VALUES ('FR9P5t8debxc11aFF', 'app'),
('FR9P5t8debxc11aFF', 'user');

INSERT INTO `role` (`id`, `app_id`, `name`)
VALUES ('1lytMwQL4uiNd0vsc', 'FR9P5t8debxc11aFF', 'admin');

INSERT INTO `access` (`app_id`, `name`, `role_id`, `user_id`,`level`)
VALUES ('FR9P5t8debxc11aFF', 'app', '1lytMwQL4uiNd0vsc', NULL,5),
('FR9P5t8debxc11aFF', 'user', '1lytMwQL4uiNd0vsc', NULL,5);

