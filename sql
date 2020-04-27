CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(40) DEFAULT '' COMMENT '用户名',
  `password` varchar(20) DEFAULT '' COMMENT '密码',
  `age` int(10) unsigned DEFAULT 0 COMMENT '年龄',
  `sex` varchar(20) DEFAULT '' COMMENT '性别',
  `avatar` varchar(255) DEFAULT '' COMMENT '头像',
  PRIMARY KEY(`id`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户管理';

CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned DEFAULT 0 COMMENT '用户ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `summary` text,
  `author` varchar(100) DEFAULT '' COMMENT '文章作者',
  `cover_url` varchar(255) DEFAULT '' COMMENT '文章封面',
  `created_on` int(11) DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(11) DEFAULT 0 COMMENT '修改时间',
  `deleted_on` int(11) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '文章管理';

# type 
  0: 文本
  1: 图片
  2: 视频
CREATE TABLE `article_detail` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned DEFAULT 0 COMMENT '文章ID',
  `type` int(10) unsigned DEFAULT 0 COMMENT '内容类型',
  `content` text,
  `desc` varchar(255) DEFAULT '' COMMENT '内容描述',
  `created_on` int (11) DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(11) DEFAULT 0 COMMENT '修改时间',
  `deleted_on` int(11) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '文章内容管理';