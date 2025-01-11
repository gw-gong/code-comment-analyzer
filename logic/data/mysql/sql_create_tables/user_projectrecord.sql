/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : code_comment_analyzer

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 11/01/2025 22:26:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_projectrecord
-- ----------------------------
DROP TABLE IF EXISTS `user_projectrecord`;
CREATE TABLE `user_projectrecord`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `project_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operating_record_id` bigint(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `operating_record_id`(`operating_record_id`) USING BTREE,
  CONSTRAINT `user_projectrecord_operating_record_id_434718a7_fk_user_oper` FOREIGN KEY (`operating_record_id`) REFERENCES `user_operatingrecord` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
