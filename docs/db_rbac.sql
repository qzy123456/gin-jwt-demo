/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50647
 Source Host           : localhost:3306
 Source Schema         : db_rbac

 Target Server Type    : MySQL
 Target Server Version : 50647
 File Encoding         : 65001

 Date: 02/08/2021 13:58:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for persistent_logins
-- ----------------------------
DROP TABLE IF EXISTS `persistent_logins`;
CREATE TABLE `persistent_logins` (
  `username` varchar(64) NOT NULL,
  `series` varchar(64) NOT NULL,
  `token` varchar(64) NOT NULL,
  `last_used` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`series`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tbl_menu
-- ----------------------------
DROP TABLE IF EXISTS `tbl_menu`;
CREATE TABLE `tbl_menu` (
  `menu_id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `menuname` varchar(50) DEFAULT NULL COMMENT '权限名',
  `menu_url` varchar(50) DEFAULT NULL COMMENT '菜单url',
  `parent_id` bigint(11) DEFAULT NULL COMMENT '父级ID',
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='权限表';

-- ----------------------------
-- Records of tbl_menu
-- ----------------------------
BEGIN;
INSERT INTO `tbl_menu` VALUES (1, '管理员管理', '', 0);
INSERT INTO `tbl_menu` VALUES (2, '商品管理', '', 0);
INSERT INTO `tbl_menu` VALUES (3, '订单管理', '', 0);
INSERT INTO `tbl_menu` VALUES (4, '系统设置', '', 0);
INSERT INTO `tbl_menu` VALUES (5, '统计分析', '/settings/echart.html', 4);
INSERT INTO `tbl_menu` VALUES (6, '日历安排', '/settings/calendar.html', 4);
INSERT INTO `tbl_menu` VALUES (7, '日志列表', '/settings/logs.html', 4);
INSERT INTO `tbl_menu` VALUES (8, '管理员列表', '/system/list.html', 1);
INSERT INTO `tbl_menu` VALUES (9, '角色管理', '/role/list.html', 1);
INSERT INTO `tbl_menu` VALUES (10, '权限管理', '/permission/list.html', 1);
INSERT INTO `tbl_menu` VALUES (11, '库存管理', '', 0);
INSERT INTO `tbl_menu` VALUES (12, '其他管理', NULL, 0);
INSERT INTO `tbl_menu` VALUES (13, '会员管理', NULL, 0);
COMMIT;

-- ----------------------------
-- Table structure for tbl_role
-- ----------------------------
DROP TABLE IF EXISTS `tbl_role`;
CREATE TABLE `tbl_role` (
  `role_id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(50) DEFAULT NULL COMMENT '角色名称',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of tbl_role
-- ----------------------------
BEGIN;
INSERT INTO `tbl_role` VALUES (1, '管理员');
INSERT INTO `tbl_role` VALUES (2, '普通用户');
COMMIT;

-- ----------------------------
-- Table structure for tbl_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `tbl_role_menu`;
CREATE TABLE `tbl_role_menu` (
  `role_id` bigint(11) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(11) DEFAULT NULL COMMENT '权限ID',
  KEY `FK_FK_PERMISSION` (`menu_id`),
  KEY `FK_FK_ROLE` (`role_id`),
  CONSTRAINT `FK_FK_PERMISSION` FOREIGN KEY (`menu_id`) REFERENCES `tbl_menu` (`menu_id`),
  CONSTRAINT `FK_FK_ROLE` FOREIGN KEY (`role_id`) REFERENCES `tbl_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色权限表';

-- ----------------------------
-- Records of tbl_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `tbl_role_menu` VALUES (1, 1);
INSERT INTO `tbl_role_menu` VALUES (1, 2);
INSERT INTO `tbl_role_menu` VALUES (1, 3);
INSERT INTO `tbl_role_menu` VALUES (1, 4);
INSERT INTO `tbl_role_menu` VALUES (1, 5);
INSERT INTO `tbl_role_menu` VALUES (1, 6);
INSERT INTO `tbl_role_menu` VALUES (1, 7);
INSERT INTO `tbl_role_menu` VALUES (1, 8);
INSERT INTO `tbl_role_menu` VALUES (1, 9);
INSERT INTO `tbl_role_menu` VALUES (1, 10);
INSERT INTO `tbl_role_menu` VALUES (1, 11);
INSERT INTO `tbl_role_menu` VALUES (1, 12);
INSERT INTO `tbl_role_menu` VALUES (1, 13);
INSERT INTO `tbl_role_menu` VALUES (2, 2);
INSERT INTO `tbl_role_menu` VALUES (2, 11);
INSERT INTO `tbl_role_menu` VALUES (2, 4);
INSERT INTO `tbl_role_menu` VALUES (2, 6);
INSERT INTO `tbl_role_menu` VALUES (2, 7);
INSERT INTO `tbl_role_menu` VALUES (2, 12);
INSERT INTO `tbl_role_menu` VALUES (2, 13);
INSERT INTO `tbl_role_menu` VALUES (2, 5);
INSERT INTO `tbl_role_menu` VALUES (2, 3);
COMMIT;

-- ----------------------------
-- Table structure for tbl_user
-- ----------------------------
DROP TABLE IF EXISTS `tbl_user`;
CREATE TABLE `tbl_user` (
  `user_id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(200) DEFAULT NULL COMMENT '密码',
  `enabled` int(2) DEFAULT NULL COMMENT '是否激活',
  `create_time` date DEFAULT NULL COMMENT '创建时间',
  `last_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of tbl_user
-- ----------------------------
BEGIN;
INSERT INTO `tbl_user` VALUES (1, 'admin', '123456', 1, '2020-02-18', '2020-02-21 04:01:15');
INSERT INTO `tbl_user` VALUES (2, 'test', '123456', 1, '2020-02-21', '2020-02-21 17:55:52');
INSERT INTO `tbl_user` VALUES (3, 'test01', '123456', 0, '2020-02-21', '2020-02-21 17:56:00');
COMMIT;

-- ----------------------------
-- Table structure for tbl_user_role
-- ----------------------------
DROP TABLE IF EXISTS `tbl_user_role`;
CREATE TABLE `tbl_user_role` (
  `user_id` bigint(11) DEFAULT NULL COMMENT '用户ID',
  `role_id` bigint(11) NOT NULL COMMENT '角色ID',
  KEY `FK_FK_USER` (`user_id`),
  CONSTRAINT `FK_FK_USER` FOREIGN KEY (`user_id`) REFERENCES `tbl_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色表';

-- ----------------------------
-- Records of tbl_user_role
-- ----------------------------
BEGIN;
INSERT INTO `tbl_user_role` VALUES (1, 1);
INSERT INTO `tbl_user_role` VALUES (2, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
