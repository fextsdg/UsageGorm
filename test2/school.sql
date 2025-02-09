/*
 Navicat Premium Dump SQL

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 50717 (5.7.17-log)
 Source Host           : localhost:3306
 Source Schema         : school

 Target Server Type    : MySQL
 Target Server Version : 50717 (5.7.17-log)
 File Encoding         : 65001

 Date: 09/02/2025 15:48:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Credit` float NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for major
-- ----------------------------
DROP TABLE IF EXISTS `major`;
CREATE TABLE `major`  (
  `MajorNum` bigint(20) NOT NULL,
  `MajorName` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`MajorNum`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stu_course
-- ----------------------------
DROP TABLE IF EXISTS `stu_course`;
CREATE TABLE `stu_course`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `StudentNum` bigint(20) NOT NULL,
  `TCourseNum` bigint(20) NOT NULL,
  `Grade` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '暂未上传',
  `Time` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `StudentNum`(`StudentNum`) USING BTREE,
  INDEX `stu_course_ibfk_2`(`TCourseNum`) USING BTREE,
  CONSTRAINT `stu_course_ibfk_1` FOREIGN KEY (`StudentNum`) REFERENCES `student` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `stu_course_ibfk_2` FOREIGN KEY (`TCourseNum`) REFERENCES `t_course` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `MajorNum` bigint(20) NULL DEFAULT NULL,
  `Identity` tinyint(4) NULL DEFAULT 1,
  `Credit` float NULL DEFAULT 0,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `MajorNum`(`MajorNum`) USING BTREE,
  CONSTRAINT `student_ibfk_1` FOREIGN KEY (`MajorNum`) REFERENCES `major` (`MajorNum`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_course
-- ----------------------------
DROP TABLE IF EXISTS `t_course`;
CREATE TABLE `t_course`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CourseNum` bigint(20) NOT NULL,
  `TeacherNum` bigint(20) NOT NULL,
  `Time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Num` int(11) NULL DEFAULT 0,
  `Total` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `CourseNum`(`CourseNum`) USING BTREE,
  INDEX `TeacherNum`(`TeacherNum`) USING BTREE,
  CONSTRAINT `t_course_ibfk_1` FOREIGN KEY (`CourseNum`) REFERENCES `course` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `t_course_ibfk_2` FOREIGN KEY (`TeacherNum`) REFERENCES `teacher` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Identity` tinyint(4) NULL DEFAULT 2,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
