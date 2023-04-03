CREATE DATABASE IF NOT EXISTS `recommend` CHARACTER SET = 'utf8mb4';
CREATE USER 'recommend'@'%' IDENTIFIED BY 'recommend123456';
GRANT ALL PRIVILEGES ON recommend.* TO 'recommend'@'%';

USE `recommend`;

CREATE TABLE `article` (
  `id` bigint NOT NULL COMMENT '主键。',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='文章。';
