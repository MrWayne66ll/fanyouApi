/*
Date: 2018-12-26 20:41:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for acl_app
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(128) CHARACTER SET utf8 NOT NULL COMMENT '花名拼音',
  `username_cn` varchar(128) CHARACTER SET utf8 NOT NULL COMMENT '花名中文',
  `email` varchar(256) CHARACTER SET utf8 NOT NULL COMMENT '用户邮箱',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `floor` varchar(64) CHARACTER SET utf8 NOT NULL COMMENT '用户所在楼层',
  `shelf` varchar(64) CHARACTER SET utf8 NOT NULL COMMENT '饭架编号',
  `active` tinyint(4) DEFAULT 1 NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `food`;
CREATE TABLE `food` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `foodName` varchar(256) CHARACTER SET utf8 NOT NULL COMMENT '食物名称',
  `status` varchar(128) CHARACTER SET utf8 DEFAULT 'release' COMMENT '食物状态，发布、被抢、被领取',
  `release_time` datetime NOT NULL COMMENT '发布时间',
  `get_time` datetime NOT NULL COMMENT '被获取时间',
  `food_type` varchar(256) CHARACTER SET utf8 NOT NULL COMMENT '食物种类，早、中、晚、夜宵',
  `comment` varchar(256) CHARACTER SET utf8 DEFAULT '快来取我吧~' COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;