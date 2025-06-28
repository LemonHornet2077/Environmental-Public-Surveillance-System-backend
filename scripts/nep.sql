/*
MySQL Data Transfer
Source Host: localhost
Source Database: nep
Target Host: localhost
Target Database: nep
*/

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `admin_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '系统管理员编号',
  `admin_code` varchar(20) NOT NULL COMMENT '系统管理员登录编码',
  `password` varchar(20) NOT NULL COMMENT '系统管理员登录密码',
  `remarks` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`admin_id`),
  UNIQUE KEY `dis_code` (`admin_code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for aqi
-- ----------------------------
DROP TABLE IF EXISTS `aqi`;
CREATE TABLE `aqi` (
  `aqi_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '空气质量指数级别(共六级)',
  `chinese_explain` varchar(10) NOT NULL COMMENT '空气质量指数级别汉字表述',
  `aqi_explain` varchar(20) NOT NULL COMMENT '空气质量指数级别描述',
  `color` varchar(7) NOT NULL COMMENT '空气质量指数级别表示颜色',
  `health_impact` varchar(500) NOT NULL COMMENT '对健康影响情况',
  `take_steps` varchar(500) NOT NULL COMMENT '建议采取的措施',
  `so2_min` int(11) NOT NULL COMMENT '本级别二氧化硫浓度最小限值',
  `so2_max` int(11) NOT NULL COMMENT '本级别二氧化硫浓度最大限值',
  `co_min` int(11) NOT NULL COMMENT '本级别一氧化碳浓度最小限值',
  `co_max` int(11) NOT NULL COMMENT '本级别一氧化碳浓度最大限值',
  `spm_min` int(11) NOT NULL COMMENT '本级别悬浮颗粒物浓度最小限值',
  `spm_max` int(11) NOT NULL COMMENT '本级别悬浮颗粒物浓度最大限值',
  `remarks` varchar(100) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`aqi_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for aqi_feedback
-- ----------------------------
DROP TABLE IF EXISTS `aqi_feedback`;
CREATE TABLE `aqi_feedback` (
  `af_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '空气质量公众监督反馈信息编号',
  `tel_id` varchar(20) NOT NULL COMMENT '所属公众监督员编号（即手机号码）',
  `province_id` int(11) NOT NULL COMMENT '反馈信息所在省区域编号',
  `city_id` int(11) NOT NULL COMMENT '反馈信息所在市区域编号',
  `address` varchar(200) NOT NULL COMMENT '反馈信息所在区域详细地址',
  `information` varchar(400) NOT NULL COMMENT '反馈信息描述',
  `estimated_grade` int(11) NOT NULL COMMENT '反馈者对空气质量指数级别的预估等级',
  `af_date` varchar(20) NOT NULL COMMENT '反馈日期',
  `af_time` varchar(20) NOT NULL COMMENT '反馈时间',
  `gm_id` int(11) NOT NULL DEFAULT '0' COMMENT '指派网格员编号',
  `assign_date` varchar(20) DEFAULT NULL COMMENT '指派日期',
  `assign_time` varchar(20) DEFAULT NULL COMMENT '指派时间',
  `state` int(11) NOT NULL COMMENT '信息状态: 0:未指派; 1:已指派; 2:已确认',
  `remarks` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`af_id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for grid_city
-- ----------------------------
DROP TABLE IF EXISTS `grid_city`;
CREATE TABLE `grid_city` (
  `city_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '系统网格覆盖市区域编号',
  `city_name` varchar(20) NOT NULL COMMENT '系统网格覆盖市区域名称',
  `province_id` int(11) NOT NULL COMMENT '所属省区域编号',
  `remarks` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`city_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for grid_member
-- ----------------------------
DROP TABLE IF EXISTS `grid_member`;
CREATE TABLE `grid_member` (
  `gm_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '网格员编号',
  `gm_name` varchar(20) NOT NULL COMMENT '网格员名称',
  `gm_code` varchar(20) NOT NULL COMMENT '网格员登录编码',
  `password` varchar(20) NOT NULL COMMENT '格网员登录密码',
  `province_id` int(11) NOT NULL COMMENT '网格区域：省编号',
  `city_id` int(11) NOT NULL COMMENT '网格区域：市编号',
  `tel` varchar(20) NOT NULL COMMENT '联系电话',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '网格员状态（0:工作状态; 1:非工作状态（由考勤系统管理）; 2:其它）',
  `remarks` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`gm_id`),
  UNIQUE KEY `gm_code` (`gm_code`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for grid_province
-- ----------------------------
DROP TABLE IF EXISTS `grid_province`;
CREATE TABLE `grid_province` (
  `province_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '系统网格覆盖省区域编号',
  `province_name` varchar(20) NOT NULL COMMENT '系统网格覆盖省区域名称',
  `province_abbr` varchar(2) NOT NULL COMMENT '系统网格覆盖省区域简称',
  `remarks` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`province_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for statistics
-- ----------------------------
DROP TABLE IF EXISTS `statistics`;
CREATE TABLE `statistics` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '统计信息编号',
  `province_id` int(11) NOT NULL COMMENT '所属省区域编号',
  `city_id` int(11) NOT NULL COMMENT '所属市区域编号',
  `address` varchar(200) NOT NULL COMMENT '反馈信息所在区域详细地址',
  `so2_value` int(11) NOT NULL COMMENT '空气二氧化硫浓度值（单位：μg/m3）',
  `so2_level` int(11) NOT NULL COMMENT '空气二氧化硫指数级别',
  `co_value` int(11) NOT NULL COMMENT '空气一氧化碳浓度值（单位：μg/m3）',
  `co_level` int(11) NOT NULL COMMENT '空气一氧化碳指数级别',
  `spm_value` int(11) NOT NULL COMMENT '空气悬浮颗粒物浓度值（单位：μg/m3）',
  `spm_level` int(11) NOT NULL COMMENT '空气PM2.5指数级别',
  `aqi_id` int(11) NOT NULL COMMENT '空气质量指数级别',
  `confirm_date` varchar(20) NOT NULL COMMENT '确认日期',
  `confirm_time` varchar(20) NOT NULL COMMENT '确认时间',
  `gm_id` int(11) NOT NULL COMMENT '所属网格员编号',
  `fd_id` varchar(20) NOT NULL COMMENT '反馈者编号（公众监督员电话号码）',
  `information` varchar(400) NOT NULL COMMENT '反馈信息描述',
  `remarks` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for supervisor
-- ----------------------------
DROP TABLE IF EXISTS `supervisor`;
CREATE TABLE `supervisor` (
  `tel_id` varchar(11) NOT NULL COMMENT '公众监督员编号（即手机号码）',
  `password` varchar(20) NOT NULL COMMENT '公众监督员登录密码',
  `real_name` varchar(20) NOT NULL COMMENT '公众监督员真实姓名',
  `birthday` varchar(20) NOT NULL COMMENT '公众监督员出生日期',
  `sex` int(11) NOT NULL DEFAULT '1' COMMENT '公众监督员性别（1：男；0：女）',
  `remarks` varchar(100) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`tel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;