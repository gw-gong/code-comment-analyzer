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

 Date: 11/01/2025 22:26:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_operatingrecord
-- ----------------------------
DROP TABLE IF EXISTS `user_operatingrecord`;
CREATE TABLE `user_operatingrecord`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `operation_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_operatingrecord_user_id_e78978af_fk_user_user_uid`(`user_id`) USING BTREE,
  CONSTRAINT `user_operatingrecord_user_id_e78978af_fk_user_user_uid` FOREIGN KEY (`user_id`) REFERENCES `user_user` (`uid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 154 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
