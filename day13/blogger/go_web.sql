/*
 Navicat Premium Data Transfer

 Source Server         : localhost MySQL
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : go_web

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 07/07/2019 13:46:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) unsigned NOT NULL COMMENT '分类id',
  `content` longtext NOT NULL COMMENT '文章内容',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `view_count` int(255) unsigned NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) unsigned NOT NULL COMMENT '评论次数',
  `username` varchar(128) NOT NULL COMMENT '作者',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_view_count` (`view_count`) USING BTREE COMMENT '阅读次数索引',
  KEY `idx_comment_count` (`comment_count`) USING BTREE COMMENT '评论数索引',
  KEY `idx_category_id` (`category_id`) USING BTREE COMMENT '分类id索引'
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` VALUES (1, 1, 'Vue真简单', 'Vue真简单', 10, 0, '尤雨溪', 1, 'Vue真简单', '2019-07-05 20:54:32', '2019-07-05 20:54:32');
INSERT INTO `article` VALUES (2, 5, 'Go语言真简洁', 'Go语言真简洁', 100, 0, '若波派克', 1, 'Go语言真简洁', '2019-07-05 20:55:31', '2019-07-05 20:55:31');
INSERT INTO `article` VALUES (3, 5, 'Gin框架真简单', 'Gin framework', 130, 0, 'gin-gonic', 1, '学好Gin框架走遍天下都不怕', '2019-07-06 21:51:11', '2019-07-06 21:51:11');
INSERT INTO `article` VALUES (4, 2, 'Java\'s verbose', 'Java\'s verbose', 200, 0, 'java', 1, 'Java\'s verbose', '2019-07-06 21:52:06', '2019-07-06 21:52:06');
INSERT INTO `article` VALUES (5, 3, 'C++真难啊', 'C++真难啊', 20, 0, 'cplusplus', 1, 'C++真难啊', '2019-07-06 21:52:40', '2019-07-06 21:52:40');
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int(10) unsigned NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, '前端开发', 1, '2019-06-29 22:55:45', '2019-06-12 22:59:00');
INSERT INTO `category` VALUES (2, 'Java开发', 2, '2019-06-29 22:56:16', '2019-06-12 22:59:05');
INSERT INTO `category` VALUES (3, 'C++开发', 3, '2019-06-29 22:56:24', '2019-06-12 22:59:08');
INSERT INTO `category` VALUES (4, '架构剖析', 4, '2019-06-29 22:56:36', '2019-06-12 22:59:10');
INSERT INTO `category` VALUES (5, 'Golang开发', 5, '2019-06-29 22:56:45', '2019-06-12 22:59:14');
COMMIT;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text NOT NULL COMMENT '评论内容',
  `username` varchar(64) NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) unsigned NOT NULL COMMENT '评论状态: 0, 删除；1， 正常',
  `article_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) NOT NULL,
  `nickname` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `gender` tinyint(4) NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
