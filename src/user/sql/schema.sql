CREATE DATABASE IF NOT EXISTS `user` CHARACTER SET = 'utf8mb4';
CREATE USER 'user'@'%' IDENTIFIED BY 'user123456';
GRANT ALL PRIVILEGES ON user.* TO 'user'@'%';

CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键。',
  `name` varchar(20) NOT NULL COMMENT '用户名。用户在注册时指定，不能更改。',
  `nickname` varchar(32) NOT NULL COMMENT '昵称。',
  `avatar` varchar(255) NOT NULL COMMENT '头像url。',
  `password` varchar(200) NOT NULL COMMENT '密码。通常是原始密码的hash。',
  `bio` varchar(500) NOT NULL COMMENT '个人简介。',
  `version` int NOT NULL DEFAULT '0' COMMENT '版本号。',
  `create_time` datetime NOT NULL COMMENT '创建时间。',
  `update_time` datetime NOT NULL COMMENT '更新时间。',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息。';

CREATE TABLE `role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键。',
  `name` varchar(45) NOT NULL COMMENT '角色名。',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色。';

CREATE TABLE `permission` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键。',
  `name` varchar(45) NOT NULL COMMENT '权限名。通常是某项操作的名称。',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限。';

CREATE TABLE `user_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键。',
  `user_id` bigint NOT NULL COMMENT '指向user表的外键。',
  `role_id` bigint NOT NULL COMMENT '指向role表的外键。',
  PRIMARY KEY (`id`),
  KEY `fk_user_role_user_idx` (`user_id`),
  KEY `fk_user_role_role_idx` (`role_id`),
  CONSTRAINT `fk_user_role_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
  CONSTRAINT `fk_user_role_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关系。';

CREATE TABLE `role_permission` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键。',
  `role_id` bigint NOT NULL COMMENT '指向role表的外键。',
  `permission_id` bigint NOT NULL COMMENT '指向permission表的外键。',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限关系。';
