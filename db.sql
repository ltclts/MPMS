CREATE TABLE `user` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
`email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
`phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
`password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
`sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
`status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '有效状态',
`type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '类型',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';