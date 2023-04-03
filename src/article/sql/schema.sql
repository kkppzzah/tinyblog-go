CREATE DATABASE IF NOT EXISTS `article` CHARACTER SET = 'utf8mb4';
CREATE USER 'article'@'%' IDENTIFIED BY 'article123456';
GRANT ALL PRIVILEGES ON article.* TO 'article'@'%';

USE `article`;

CREATE TABLE `article_meta` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键。',
  `user_id` bigint NOT NULL COMMENT '指向user表的外键。',
  `title` varchar(120) NOT NULL COMMENT '文章标题。',
  `tags` varchar(40) NOT NULL COMMENT '文章标签。',
  `summary` varchar(200) NOT NULL COMMENT '文章概要。',
  `version` int NOT NULL COMMENT '版本号。',
  `publish_time` datetime NOT NULL COMMENT '发布时间。',
  `update_time` datetime NOT NULL COMMENT '更新时间。',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb3 COMMENT='用来存放文章的元数据。';

CREATE TABLE `article_content` (
  `id` bigint NOT NULL COMMENT '文章ID。指向article_meta表的外键。',
  `content` varchar(10000) NOT NULL COMMENT '文章内容。',
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_article_content_meta` FOREIGN KEY (`id`) REFERENCES `article_meta` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='文章内容。';

CREATE TABLE `user` (
  `id` bigint NOT NULL COMMENT '主键。',
  `name` varchar(20) NOT NULL COMMENT '用户名。',
  `nickname` varchar(32) NOT NULL COMMENT '昵称。',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户信息。';
