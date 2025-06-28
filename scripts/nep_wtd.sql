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

-- ----------------------------
-- Records 
-- ----------------------------
INSERT INTO `admins` VALUES ('1', 'admin', '123', null);
INSERT INTO `aqi` VALUES ('1', '一', '优', '#02E300', '空气质量令人满意，基本无空气污染', '各类人群可正常活动', '0', '50', '0', '5', '0', '35', null);
INSERT INTO `aqi` VALUES ('2', '二', '良', '#FFFF00', '空气质量可接受，但某些污染物口能对极少数异常敏感人群健康有较弱影响', '极少数异常敏感人群应减少户外活动', '51', '150', '6', '10', '36', '75', null);
INSERT INTO `aqi` VALUES ('3', '三', '轻度污染', '#FF7E00', '易感人群症状有轻度加剧，健康人群出现刺激症状', '儿童、老年人及心脏病、呼吸系统疾病患者应减少长时间、高强度的户外锻炼', '151', '475', '11', '35', '76', '115', null);
INSERT INTO `aqi` VALUES ('4', '四', '中度污染', '#FE0000', '进一步加剧易感人群症状，可能对健康人群心脏、呼吸系统有影响', '儿童、老年人及心脏病、呼吸系统疾病患者避免长时间、高强度的户外锻练，一般人群适量减少户外运动', '476', '800', '36', '60', '116', '150', null);
INSERT INTO `aqi` VALUES ('5', '五', '重度污染', '#98004B', '心脏病和肺病患者症状显著加剧，运动耐受力降低，健康人群普遍出现症状', '儿童、老年人和心脏病、肺病患者应停留在室内，停止户外运动，-般人群减少户外运动', '801', '1600', '61', '90', '151', '250', null);
INSERT INTO `aqi` VALUES ('6', '六', '严重污染', '#7E0123', '健康人群运动耐受力降低，有明显强烈症状，提前出现某些疾病', '儿童、老年人和病人应当留在室内，避免体力消耗，一般人群应避免户外活动', '1601', '2620', '91', '150', '251', '500', null);
INSERT INTO `aqi_feedback` VALUES ('1', '13147859658', '1', '1', '朝阳区建国路123号', '空气能见度不足，稍有异味。', '3', '2022-01-26', '09:28:04', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('2', '13245871254', '2', '2', '塘沽区延庆街乐亭理', '空气中似乎有粉尘，呼吸不畅，刺激。', '5', '2022-02-26', '09:32:16', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('3', '13369852458', '3', '3', '昌平区临西路45-69号', '月朦胧，鸟朦胧，空气雾霾浓。', '5', '2022-03-26', '09:36:12', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('4', '13512345678', '4', '4', '阳泉区天镇街平顺胡同', '空气污染严重～昨天洗的车子，今天一层灰～真心伤不起。', '6', '2022-04-26', '09:37:38', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('5', '13512345688', '5', '5', '巴彦淖尔区奈曼旗', '扬尘飞沙扑面来，泥土气息撞满怀。', '5', '2022-05-26', '09:38:45', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('6', '13545645612', '6', '6', '浑南区彩霞路彩霞社区', '花朦胧，叶朦胧，医院排长队。', '4', '2022-06-26', '09:40:02', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('7', '13566987452', '6', '17', '清原满族自治县迎宾路', '每天漫天灰尘，出门一分钟，回来一身灰。', '6', '2022-07-26', '09:41:01', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('8', '13655669988', '7', '7', '集安区白山街白城社区', '环境污染，全球变暖，钓鱼人的环境越来越差。', '3', '2022-08-26', '09:42:02', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('9', '13688998874', '8', '8', '双城区海林路五常里', '近年来空气污染越发严重，PM2.5值越来越高，眼睛经常有异物感。', '4', '2022-08-26', '09:43:20', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('10', '13758745632', '9', '9', '徐汇区明水路集贤里', '环境脏了，脏了的不仅是环境，更是心情。', '5', '2022-09-26', '09:44:19', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('11', '13847895623', '10', '10', '江都区杜尔伯特街456号', 'PM2.5值越来越高，眼睛经常有异物感，导致眼部发病率急剧升高。', '5', '2022-09-26', '10:03:17', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('12', '13900240032', '11', '11', '金湖区响水路东海社区', '呼吸的空气里面都带各种污染，环境真是让人堪忧哦！', '6', '2022-10-26', '10:04:35', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('13', '13954754744', '12', '12', '天台区宁海街道1-123-3号', '雾霾，我们的生活条件也在不断提高，但是生活的环境真是不尽人意。', '4', '2022-01-26', '10:05:34', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('14', '14555889874', '13', '13', '南平区无为路7-789-9号', '身边都是乌烟瘴气，烟雾缭绕可能一个字，都会成为最致命的“导火线”。', '6', '2022-02-26', '10:06:31', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('15', '14955226688', '14', '14', '高安区永丰路玉山社区', '一阵狂风破青天，所谓雾霾成云烟，万千人财不胜冷，吾辈环工情何堪。', '3', '2022-02-26', '10:08:01', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('16', '15353245698', '15', '15', '临淄区胶南街444号', '连日阴雨，空气潮湿，阴云沉沉，心事重重。', '3', '2022-03-26', '10:09:08', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('17', '15544523687', '16', '16', '修武区登封街新乡社区', '身边都是乌烟瘴气，烟雾缭绕。', '4', '2022-04-26', '10:10:55', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('18', '15560023569', '1', '1', '怀柔区北辰街道78号', '人类的环境破坏已经变得千疮百孔，保护大自然，人人有责。', '5', '2022-04-26', '10:13:41', '1', '2022-11-27', '11:03:29', '2', null);
INSERT INTO `aqi_feedback` VALUES ('19', '15655881122', '9', '9', ' 巨鹿区灵寿路安国里', '环境被破坏，地球在哭嚎。本色皆可期，全靠你我他。', '3', '2022-05-26', '10:18:15', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('20', '15800556874', '11', '11', '盐山县南皮街道56号', '环境污染天天喊，其实污染在传染，污染水源植物减', '4', '2022-05-26', '10:19:25', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('21', '17345988896', '6', '6', '和平区盐泉路456号', '扬尘飞沙扑面来，泥土气息撞满怀。', '3', '2022-05-26', '10:20:29', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('22', '17522112211', '13', '13', '南和区邢台街好好社区', '穹顶之下，雾霾锁城。环境污染是一个摆在所有人面前的问题。', '5', '2022-06-26', '10:21:25', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('23', '17645614561', '6', '17', '东光区高峰会胡同', '月朦胧，鸟朦胧，空气雾霾浓。', '5', '2022-06-26', '10:22:25', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('24', '17733658965', '15', '15', '大城区文安路阳泉胡同', '天空灰蒙蒙的一片、空气里散发着刺鼻的味道，让人感到压抑。', '4', '2022-06-26', '10:23:44', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('25', '18065895234', '14', '14', '长治区阳高路421号', '呼吸的空气里面都带各种污染，环境真是让人堪忧哦！', '6', '2022-07-26', '10:24:40', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('26', '18165214789', '8', '8', '静乐区丰镇路789号', '沙尘风暴又雾霾，保护环境皆有责。', '3', '2022-07-26', '10:25:46', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('27', '18558743311', '16', '16', '杭锦旗土默特左旗乌拉特社区', '每天漫天灰尘，出门一分钟，回来一身灰。', '5', '2022-07-26', '10:27:16', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('28', '18655441236', '2', '2', '和龙区柳河街1-123-1号', '花朦胧，叶朦胧，医院排长队', '3', '2022-08-26', '10:28:26', '1', '2022-11-27', '11:04:08', '2', null);
INSERT INTO `aqi_feedback` VALUES ('29', '18925321123', '4', '4', '孙吴区廉颇路李牧社区', 'PM2.5值越来越高，眼睛经常有异物感，导致眼部发病率急剧升高。', '4', '2022-08-26', '10:29:46', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('30', '13147859658', '11', '11', '尚志区友谊路友谊社区', '环境污染了，污染了的不仅是环境，更是健康。', '5', '2022-08-26', '10:32:34', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('31', '13245871254', '10', '10', '仙居区仙女路仙人社区', '环境污染，全球变暖，钓鱼人的环境越来越差。', '3', '2022-09-26', '10:33:58', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('32', '13369852458', '8', '8', '界首区阜南街霍山街道', '一阵狂风破青天，所谓雾霾成云烟。', '5', '2022-09-26', '10:35:12', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('33', '13512345678', '16', '16', '肥西区分东路费义里', '清晨雾蒙蒙，世间万物皆胧罩，恰似人间仙境，雾霾满城，活吞天地，繁华遮尽，唯有心近。', '4', '2022-09-26', '10:36:56', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('34', '13512345688', '9', '9', '浦东区玉环路4-56-4号', '连日阴雨，空气潮湿，阴云沉沉，心事重重。', '4', '2022-09-26', '10:37:56', '9', '2022-11-25', '12:52:56', '1', null);
INSERT INTO `aqi_feedback` VALUES ('35', '13545645612', '4', '4', '庆元区景宁畲族自治县', '如果地球生态失衡，自然灾害就会增多。', '3', '2022-10-26', '10:39:13', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('36', '13566987452', '10', '10', '明光区六安路五河社区', '地球在哭泣，恶劣天气频现，全球气候变暖，爱护我们的自然环境。', '6', '2022-10-26', '10:39:57', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('37', '13655669988', '12', '12', '建瓯区邵武大街7-8-9号', '起起伏伏，跌跌荡荡。归于平静，波澜不惊。', '5', '2022-10-26', '10:41:05', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('38', '13688998874', '6', '17', '甘井子区凌风街乘风社区', '月黑风高，空气浑浊，难道是杀人夜？', '4', '2022-10-27', '16:29:26', '0', null, null, '0', null);
INSERT INTO `aqi_feedback` VALUES ('39', '13758745632', '4', '4', '西山区解放大路1-258-6号', '雾朦胧，鸟朦胧，一切都朦胧。', '3', '2022-11-03', '11:09:09', '4', '2022-11-25', '12:31:25', '1', null);
INSERT INTO `grid_city` VALUES ('1', '北京市', '1', null);
INSERT INTO `grid_city` VALUES ('2', '天津市', '2', null);
INSERT INTO `grid_city` VALUES ('3', '石家庄市', '3', null);
INSERT INTO `grid_city` VALUES ('4', '太原市', '4', null);
INSERT INTO `grid_city` VALUES ('5', '呼和浩特市', '5', null);
INSERT INTO `grid_city` VALUES ('6', '沈阳市', '6', null);
INSERT INTO `grid_city` VALUES ('7', '长春市', '7', null);
INSERT INTO `grid_city` VALUES ('8', '哈尔滨市', '8', null);
INSERT INTO `grid_city` VALUES ('9', '上海市', '9', null);
INSERT INTO `grid_city` VALUES ('10', '南京市', '10', null);
INSERT INTO `grid_city` VALUES ('11', '杭州市', '11', null);
INSERT INTO `grid_city` VALUES ('12', '合肥市', '12', null);
INSERT INTO `grid_city` VALUES ('13', '福州市', '13', null);
INSERT INTO `grid_city` VALUES ('14', '南昌市', '14', null);
INSERT INTO `grid_city` VALUES ('15', '济南市', '15', null);
INSERT INTO `grid_city` VALUES ('16', '郑州市', '16', null);
INSERT INTO `grid_city` VALUES ('17', '大连市', '6', null);
INSERT INTO `grid_member` VALUES ('1', '曹操', 'caocao', '123', '1', '1', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('2', '刘备', 'liubei', '123', '2', '2', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('3', '孙权', 'sunquan', '123', '3', '3', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('4', '关羽', 'guanyu', '123', '4', '4', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('5', '张飞', 'zhangfei', '123', '5', '5', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('6', '诸葛亮', 'zhugeliang', '123', '6', '6', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('7', '赵云', 'zhaoyun', '123', '7', '7', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('8', '黄忠', 'huangzhong', '123', '8', '8', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('9', '马超', 'machao', '123', '9', '9', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('10', '魏延', 'weiyan', '123', '10', '10', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('11', '法正', 'fazheng', '123', '11', '11', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('12', '庞统', 'pangtong', '123', '12', '12', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('13', '周瑜', 'zhouyu', '123', '13', '13', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('14', '司马懿', 'simayi', '123', '14', '14', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('15', '徐庶', 'xushu', '123', '15', '15', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('16', '曹丕', 'caopi', '123', '16', '16', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('17', '曹植', 'caozhi', '123', '6', '6', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('18', '郭嘉', 'guojia', '123', '6', '6', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('19', '荀彧', 'xunyu', '123', '6', '6', '12687643458', '1', null);
INSERT INTO `grid_member` VALUES ('20', '张辽', 'zhangliao', '123', '6', '17', '12687643458', '1', null);
INSERT INTO `grid_member` VALUES ('21', '徐晃', 'xuhuang', '123', '6', '17', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('22', '许褚', 'xuchu', '123', '6', '1', '12687643458', '1', null);
INSERT INTO `grid_member` VALUES ('23', '典韦', 'dianwei', '123', '1', '2', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('24', '张郃', 'zhanghe', '123', '2', '3', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('25', '于禁', 'yujin', '123', '3', '4', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('26', '乐进', 'yuejin', '123', '4', '5', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('27', '李典', 'lidian', '123', '5', '6', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('28', '曹仁', 'caoren', '123', '6', '7', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('29', '文聘', 'wenpin', '123', '7', '8', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('30', '曹洪', 'caohong', '123', '8', '9', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('31', '曹真', 'caozhen', '123', '9', '10', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('32', '夏侯渊', 'xiahouyuan', '123', '10', '10', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('33', '姜维', 'jiangwei', '123', '11', '11', '12687643458', '0', null);
INSERT INTO `grid_member` VALUES ('34', '陆逊', 'luxun', '123', '12', '12', '12687643458', '0', null);
INSERT INTO `grid_province` VALUES ('1', '北京市', '京', null);
INSERT INTO `grid_province` VALUES ('2', '天津市', '津', null);
INSERT INTO `grid_province` VALUES ('3', '河北省', '冀', null);
INSERT INTO `grid_province` VALUES ('4', '山西省', '晋', null);
INSERT INTO `grid_province` VALUES ('5', '内蒙古', '蒙', null);
INSERT INTO `grid_province` VALUES ('6', '辽宁省', '辽', null);
INSERT INTO `grid_province` VALUES ('7', '吉林省', '吉', null);
INSERT INTO `grid_province` VALUES ('8', '黑龙江省', '黑', null);
INSERT INTO `grid_province` VALUES ('9', '上海市', '沪', null);
INSERT INTO `grid_province` VALUES ('10', '江苏省', '苏', null);
INSERT INTO `grid_province` VALUES ('11', '浙江省', '浙', null);
INSERT INTO `grid_province` VALUES ('12', '安徽省', '皖', null);
INSERT INTO `grid_province` VALUES ('13', '福建省', '闽', null);
INSERT INTO `grid_province` VALUES ('14', '江西省', '赣', null);
INSERT INTO `grid_province` VALUES ('15', '山东省', '鲁', null);
INSERT INTO `grid_province` VALUES ('16', '河南省', '豫', null);
INSERT INTO `statistics` VALUES ('1', '1', '1', '怀柔区北辰街道78号', '425', '3', '42', '4', '56', '2', '4', '2022-04-26', '11:09:31', '1', '15560023569', '人类的环境破坏已经变得千疮百孔，保护大自然，人人有责。', null);
INSERT INTO `statistics` VALUES ('2', '1', '1', '朝阳区建国路123号', '123', '2', '23', '3', '144', '4', '4', '2022-01-26', '11:16:08', '1', '13045825698', '空气能见度不足，稍有异味。', null);
INSERT INTO `statistics` VALUES ('3', '2', '2', '和龙区柳河街1-123-1号', '124', '2', '6', '2', '45', '2', '2', '2022-08-26', '11:19:19', '2', '18655441236', '花朦胧，叶朦胧，医院排长队', null);
INSERT INTO `statistics` VALUES ('4', '2', '2', '塘沽区延庆街乐亭理', '111', '2', '61', '5', '89', '3', '5', '2022-02-26', '11:19:56', '2', '13147859658', '空气中似乎有粉尘，呼吸不畅，刺激。', null);
INSERT INTO `statistics` VALUES ('5', '3', '3', '昌平区临西路45-69号', '836', '5', '6', '2', '200', '5', '5', '2022-03-26', '11:21:38', '3', '13245871254', '月朦胧，鸟朦胧，空气雾霾浓。', null);
INSERT INTO `statistics` VALUES ('6', '4', '4', '庆元区景宁畲族自治县', '566', '4', '12', '3', '44', '2', '4', '2022-10-26', '11:22:47', '4', '13369852458', '如果地球生态失衡，自然灾害就会增多。', null);
INSERT INTO `statistics` VALUES ('7', '4', '4', '孙吴区廉颇路李牧社区', '78', '2', '56', '4', '111', '3', '4', '2022-08-26', '11:23:16', '4', '18925321123', 'PM2.5值越来越高，眼睛经常有异物感，导致眼部发病率急剧升高。', null);
INSERT INTO `statistics` VALUES ('8', '4', '4', '阳泉区天镇街平顺胡同', '564', '4', '222', '6', '78', '3', '6', '2022-04-26', '11:23:48', '4', '13369852458', '空气污染严重～昨天洗的车子，今天一层灰～真心伤不起。', null);
INSERT INTO `statistics` VALUES ('9', '5', '5', '巴彦淖尔区奈曼旗', '456', '3', '56', '4', '23', '1', '4', '2022-05-26', '11:26:21', '5', '13545645612', '扬尘飞沙扑面来，泥土气息撞满怀。', null);
INSERT INTO `statistics` VALUES ('10', '6', '6', '浑南区彩霞路彩霞社区', '122', '2', '12', '2', '23', '1', '2', '2022-06-26', '11:27:50', '6', '13566987452', '花朦胧，叶朦胧，医院排长队。', null);
INSERT INTO `statistics` VALUES ('11', '7', '7', '集安区白山街白城社区', '123', '2', '23', '3', '102', '3', '3', '2022-08-26', '11:28:42', '7', '13688998874', '环境污染，全球变暖，钓鱼人的环境越来越差。', null);
INSERT INTO `statistics` VALUES ('12', '8', '8', '界首区阜南街霍山街道', '456', '3', '86', '5', '45', '2', '5', '2022-09-26', '11:29:30', '8', '13045825698', '一阵狂风破青天，所谓雾霾成云烟。', null);
INSERT INTO `statistics` VALUES ('13', '8', '8', '双城区海林路五常里', '456', '3', '7', '2', '123', '4', '4', '2022-08-26', '11:29:59', '8', '13758745632', '近年来空气污染越发严重，PM2.5值越来越高，眼睛经常有异物感。', null);
INSERT INTO `statistics` VALUES ('14', '8', '8', '静乐区丰镇路789号', '56', '2', '13', '3', '45', '2', '3', '2022-07-26', '11:30:30', '8', '18165214789', '沙尘风暴又雾霾，保护环境皆有责。', null);
INSERT INTO `statistics` VALUES ('15', '9', '9', '浦东区玉环路4-56-4号', '456', '3', '7', '2', '178', '5', '5', '2022-09-26', '11:31:23', '9', '13245871254', '连日阴雨，空气潮湿，阴云沉沉，心事重重。', null);
INSERT INTO `statistics` VALUES ('16', '9', '9', '徐汇区明水路集贤里', '566', '4', '6', '2', '77', '3', '4', '2022-09-26', '11:31:53', '9', '13847895623', '环境脏了，脏了的不仅是环境，更是心情。', null);
INSERT INTO `statistics` VALUES ('17', '9', '9', '巨鹿区灵寿路安国里', '56', '2', '13', '3', '100', '3', '3', '2022-05-26', '11:32:28', '9', '15655881122', '环境被破坏，地球在哭嚎。本色皆可期，全靠你我他。', null);
INSERT INTO `statistics` VALUES ('18', '10', '10', '明光区六安路五河社区', '2234', '6', '5', '1', '456', '6', '6', '2022-10-26', '11:33:43', '10', '13545645612', '地球在哭泣，恶劣天气频现，全球气候变暖，爱护我们的自然环境。', null);
INSERT INTO `statistics` VALUES ('19', '10', '10', '江都区杜尔伯特街456号', '566', '4', '86', '5', '123', '4', '5', '2022-09-26', '11:34:16', '10', '13900240032', 'PM2.5值越来越高，眼睛经常有异物感，导致眼部发病率急剧升高。', null);
INSERT INTO `statistics` VALUES ('20', '11', '11', '盐山县南皮街道56号', '566', '4', '9', '2', '145', '4', '4', '2022-05-26', '11:35:12', '11', '15800556874', '环境污染天天喊，其实污染在传染，污染水源植物减', null);
INSERT INTO `statistics` VALUES ('21', '12', '12', '建瓯区邵武大街7-8-9号', '789', '4', '12', '3', '245', '5', '5', '2022-10-26', '11:36:08', '12', '13566987452', '起起伏伏，跌跌荡荡。归于平静，波澜不惊。', null);
INSERT INTO `statistics` VALUES ('22', '12', '12', '天台区宁海街道1-123-3号', '568', '4', '56', '4', '111', '3', '4', '2022-01-26', '11:36:34', '12', '14555889874', '雾霾，我们的生活条件也在不断提高，但是生活的环境真是不尽人意。', null);
INSERT INTO `statistics` VALUES ('23', '13', '13', '南和区邢台街好好社区', '1589', '5', '66', '5', '123', '4', '5', '2022-06-26', '11:40:26', '13', '17522112211', '穹顶之下，雾霾锁城。环境污染是一个摆在所有人面前的问题。', null);
INSERT INTO `statistics` VALUES ('24', '13', '13', '南平区无为路7-789-9号', '456', '3', '156', '6', '145', '4', '6', '2022-02-26', '11:41:15', '13', '14955226688', '身边都是乌烟瘴气，烟雾缭绕可能一个字，都会成为最致命的“导火线”。', null);
INSERT INTO `statistics` VALUES ('25', '14', '14', '长治区阳高路421号', '789', '4', '5', '1', '1234', '6', '6', '2022-07-26', '11:42:37', '14', '18065895234', '呼吸的空气里面都带各种污染，环境真是让人堪忧哦！', null);
INSERT INTO `statistics` VALUES ('26', '14', '14', '高安区永丰路玉山社区', '456', '3', '12', '3', '45', '2', '3', '2022-02-26', '11:43:10', '14', '15353245698', '一阵狂风破青天，所谓雾霾成云烟，万千人财不胜冷，吾辈环工情何堪。', null);
INSERT INTO `statistics` VALUES ('27', '15', '15', '大城区文安路阳泉胡同', '123', '2', '56', '4', '89', '3', '4', '2022-06-26', '11:44:00', '15', '17733658965', '天空灰蒙蒙的一片、空气里散发着刺鼻的味道，让人感到压抑。', null);
INSERT INTO `statistics` VALUES ('28', '15', '15', '临淄区胶南街444号', '196', '3', '13', '3', '100', '3', '3', '2022-03-26', '11:44:36', '15', '15544523687', '连日阴雨，空气潮湿，阴云沉沉，心事重重。', null);
INSERT INTO `statistics` VALUES ('29', '16', '16', '肥西区分东路费义里', '100', '2', '45', '4', '123', '4', '4', '2022-09-26', '12:51:39', '16', '13147859658', '清晨雾蒙蒙，世间万物皆胧罩，恰似人间仙境，雾霾满城，活吞天地...', null);
INSERT INTO `statistics` VALUES ('30', '16', '16', '杭锦旗土默特左旗乌拉特社区', '123', '2', '56', '4', '245', '5', '5', '2022-07-26', '12:52:18', '16', '18558743311', '每天漫天灰尘，出门一分钟，回来一身灰。', null);
INSERT INTO `statistics` VALUES ('31', '16', '16', '修武区登封街新乡社区', '455', '3', '12', '3', '123', '4', '4', '2022-04-26', '13:23:20', '16', '15544523687', '身边都是乌烟瘴气，烟雾缭绕。', null);
INSERT INTO `statistics` VALUES ('32', '6', '6', '和平区盐泉路456号', '223', '3', '8', '2', '78', '3', '3', '2022-05-26', '13:24:14', '17', '17345988896', '扬尘飞沙扑面来，泥土气息撞满怀。', null);
INSERT INTO `statistics` VALUES ('33', '6', '17', '东光区高峰会胡同', '456', '3', '45', '4', '78', '3', '4', '2022-06-26', '13:24:58', '18', '17645614561', '月朦胧，鸟朦胧，空气雾霾浓。', null);
INSERT INTO `statistics` VALUES ('34', '6', '17', '清原满族自治县迎宾路', '456', '3', '235', '6', '156', '5', '6', '2022-07-26', '13:26:14', '20', '13655669988', '每天漫天灰尘，出门一分钟，回来一身灰。', null);
INSERT INTO `statistics` VALUES ('35', '10', '10', '仙居区仙女路仙人社区', '455', '3', '6', '2', '89', '3', '3', '2022-09-26', '13:28:09', '32', '13147859658', '环境污染，全球变暖，钓鱼人的环境越来越差。', null);
INSERT INTO `statistics` VALUES ('36', '11', '11', '金湖区响水路东海社区', '456', '3', '78', '5', '365', '6', '6', '2022-10-26', '13:28:50', '33', '13954754744', '呼吸的空气里面都带各种污染，环境真是让人堪忧哦！', null);
INSERT INTO `statistics` VALUES ('37', '11', '11', '尚志区友谊路友谊社区', '897', '5', '23', '3', '156', '5', '5', '2022-08-26', '13:29:25', '33', '13045825698', '环境污染了，污染了的不仅是环境，更是健康。', null);
INSERT INTO `statistics` VALUES ('38', '6', '17', '甘井子区凌风街乘风社区', '564', '4', '25', '3', '123', '4', '4', '2022-10-27', '17:01:52', '21', '17345988896', '月黑风高，空气浑浊，难道是杀人夜？', null);
INSERT INTO `statistics` VALUES ('39', '4', '4', '西山区解放大路1-258-6号', '56', '2', '23', '3', '78', '3', '3', '2022-11-03', '11:10:57', '4', '17645614561', '雾朦胧，鸟朦胧，一切都朦胧。', null);
INSERT INTO `supervisor` VALUES ('13147859658', '123', '柯镇恶', '1984-12-09', '1', null);
INSERT INTO `supervisor` VALUES ('13245871254', '123', '朱聪', '1985-02-07', '1', null);
INSERT INTO `supervisor` VALUES ('13369852458', '123', '郭靖', '2000-10-12', '1', null);
INSERT INTO `supervisor` VALUES ('13512345678', '123', '梅超风', '1996-07-24', '0', null);
INSERT INTO `supervisor` VALUES ('13512345688', '123', '王五', '2001-06-06', '0', null);
INSERT INTO `supervisor` VALUES ('13545645612', '123', '谭处端', '1985-09-28', '1', null);
INSERT INTO `supervisor` VALUES ('13566987452', '123', '欧阳克', '2000-09-08', '1', null);
INSERT INTO `supervisor` VALUES ('13655669988', '123', '刘处玄', '1985-10-10', '1', null);
INSERT INTO `supervisor` VALUES ('13688998874', '123', '陆乘风', '1996-02-07', '1', null);
INSERT INTO `supervisor` VALUES ('13758745632', '123', '丘处机', '1985-11-14', '1', null);
INSERT INTO `supervisor` VALUES ('13847895623', '123', '王处一', '1986-05-25', '1', null);
INSERT INTO `supervisor` VALUES ('13900240032', '123', '郝大通', '1986-07-07', '1', null);
INSERT INTO `supervisor` VALUES ('13954754744', '123', '陈玄风', '1995-12-13', '1', null);
INSERT INTO `supervisor` VALUES ('14555889874', '123', '张阿生', '1985-12-17', '1', null);
INSERT INTO `supervisor` VALUES ('14955226688', '123', '段智兴', '1979-09-28', '1', null);
INSERT INTO `supervisor` VALUES ('15353245698', '123', '洪七公', '1980-02-16', '1', null);
INSERT INTO `supervisor` VALUES ('15544523687', '123', '韩宝驹', '1985-06-30', '1', null);
INSERT INTO `supervisor` VALUES ('15560023569', '123', '曲灵风', '1995-12-25', '1', null);
INSERT INTO `supervisor` VALUES ('15655881122', '123', '南希仁', '1985-10-02', '1', null);
INSERT INTO `supervisor` VALUES ('15800556874', '123', '孙不二', '1986-11-04', '0', null);
INSERT INTO `supervisor` VALUES ('17345988896', '123', '欧阳锋', '1980-11-13', '1', null);
INSERT INTO `supervisor` VALUES ('17522112211', '123', '马钰', '1984-06-21', '1', null);
INSERT INTO `supervisor` VALUES ('17645614561', '123', '瑛姑', '1986-04-20', '0', null);
INSERT INTO `supervisor` VALUES ('17733658965', '123', '黄药师', '1981-09-09', '1', null);
INSERT INTO `supervisor` VALUES ('18065895234', '123', '王重阳', '1970-10-25', '1', null);
INSERT INTO `supervisor` VALUES ('18165214789', '123', '周伯通', '1976-06-23', '1', null);
INSERT INTO `supervisor` VALUES ('18558743311', '123', '全金发', '1986-01-06', '1', null);
INSERT INTO `supervisor` VALUES ('18655441236', '123', '韩小莹', '1986-04-19', '0', null);
INSERT INTO `supervisor` VALUES ('18925321123', '123', '杨康', '2000-11-15', '1', null);
