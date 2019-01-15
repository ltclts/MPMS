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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

CREATE TABLE `config` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '类型',
`content` text NOT NULL DEFAULT '' COMMENT '内容',
`desc` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='配置表';

CREATE TABLE `menu` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父节点id',
`type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '类型 1-一级菜单 2-二级菜单',
`name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
`name_en` varchar(100) NOT NULL DEFAULT '' COMMENT '英文名',
`uri` varchar(100) NOT NULL DEFAULT '' COMMENT '路由或者图标',
`sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
`user_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '用户类别 同user.type',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_type_parent_id` (`type`,`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单表';

CREATE TABLE `flow` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`refer_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '相关id',
`refer_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '相关类型',
`status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '操作',
`content` text NOT NULL DEFAULT '' COMMENT '变更内容',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_refer_id` (`refer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通用流水表';

CREATE TABLE `relation` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`refer_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '相关id',
`refer_id_others` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '相关另外的id',
`refer_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '关系类型',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_refer_id` (`refer_id`),
KEY `idx_refer_id_others` (`refer_id_others`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通用关系表';

CREATE TABLE `company` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
`short_name` varchar(50) NOT NULL DEFAULT '' COMMENT '简称',
`remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
`status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0-初始状态 1-启用 2-禁用',
`expire_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '过期时间',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_name` (`name`),
KEY `idx_short_name` (`short_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='公司表';

CREATE TABLE `mini_program` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
`remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
`appid` varchar(100) NOT NULL DEFAULT '' COMMENT '小程序appid',
`company_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '公司id',
`status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0-初始状态 1-启用 2-禁用',
`type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '小程序类型 0-未知 1-名片展示',
`content` text NOT NULL DEFAULT '' COMMENT '内容',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_company_id` (`company_id`),
KEY `idx_appid` (`appid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='小程序表';

CREATE TABLE `mini_program_version` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`mini_program_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '小程序id',
`code` varchar(100) NOT NULL DEFAULT '' COMMENT '版本号',
`status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0-编辑中 1-审核中 2-已审核 3-已上线 4-已下线',
`is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
`creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
`created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '更新时间',
PRIMARY KEY (`id`),
KEY `idx_mini_program_id` (`mini_program_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='小程序版本表';

