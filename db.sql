/*
 Navicat Premium Data Transfer

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 29/11/2022 18:23:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rustdesk_peers
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_peers`;
CREATE TABLE `rustdesk_peers` (
  `deviceid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uid` int(10) unsigned NOT NULL COMMENT 'User ID',
  `id` char(16) NOT NULL DEFAULT '' COMMENT 'Device ID',
  `username` varchar(128) DEFAULT NULL COMMENT 'Operating system username',
  `hostname` varchar(128) DEFAULT NULL COMMENT 'Operating system name',
  `alias` char(20) DEFAULT NULL COMMENT 'Alias',
  `platform` char(20) DEFAULT NULL COMMENT 'Platform',
  `tags` varchar(256) DEFAULT NULL COMMENT 'Tags',
  PRIMARY KEY (`deviceid`),
  UNIQUE KEY `uuid` (`id`,`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=61 DEFAULT CHARSET=utf8 COMMENT='Remote device table';

-- ----------------------------
-- Records of rustdesk_peers
-- ----------------------------

-- ----------------------------
-- Table structure for rustdesk_tags
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_tags`;
CREATE TABLE `rustdesk_tags` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Tag ID',
  `uid` int(10) unsigned NOT NULL COMMENT 'User ID',
  `tag` char(20) NOT NULL DEFAULT '' COMMENT 'Tag name',
  `color` char(10) NULL COMMENT 'Tag color',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tag` (`tag`,`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=136 DEFAULT CHARSET=utf8 COMMENT='Tags table';

-- ----------------------------
-- Records of rustdesk_tags
-- ----------------------------

-- ----------------------------
-- Table structure for rustdesk_token
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_token`;
CREATE TABLE `rustdesk_token` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` char(16) NOT NULL COMMENT 'Username',
  `uid` int(10) unsigned NOT NULL COMMENT 'User ID',
  `client_id` char(16) NOT NULL COMMENT 'Device code',
  `uuid` char(64) NOT NULL COMMENT 'Device ID',
  `access_token` varchar(128) NOT NULL DEFAULT '' COMMENT 'Login token',
  `login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Login time',
  `expire_time` int(11) DEFAULT NULL COMMENT 'Expiration time',
  `active_time` int(10) DEFAULT NULL COMMENT 'Last active time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `login_token` (`uid`,`client_id`,`uuid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='Login Token table';

-- ----------------------------
-- Records of rustdesk_token
-- ----------------------------

-- ----------------------------
-- Table structure for rustdesk_users
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_users`;
CREATE TABLE `rustdesk_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `username` varchar(16) NOT NULL COMMENT 'Username',
  `password` char(32) NOT NULL COMMENT 'Password',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT 'User status',
  `last_login_ip` varchar(255) NOT NULL COMMENT 'Last login IP',
  `last_login_time` int(10) NOT NULL DEFAULT '0' COMMENT 'Last login time',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Creation time',
  `update_time` int(10) NOT NULL DEFAULT '0' COMMENT 'Update time',
  `delete_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Deletion time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='Users table';

-- ----------------------------
-- Records of rustdesk_users
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;