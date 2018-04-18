CREATE TABLE `problem` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `flag` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1-普通题目 2-用户题目',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '申请状态',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目提供者',
  `difficulty` varchar(40) NOT NULL DEFAULT '' COMMENT '题目难度',
  `case_data` varchar(200) NOT NULL DEFAULT '' COMMENT '测试数据',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '题目标题',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '题目描述',
  `input_des` varchar(300) NOT NULL DEFAULT '' COMMENT '输入描述',
  `output_des` varchar(300) NOT NULL DEFAULT '' COMMENT '输出描述',
  `input_case` varchar(200) NOT NULL DEFAULT '' COMMENT '测试输入',
  `output_case` varchar(200) NOT NULL DEFAULT '' COMMENT '测试输出',
  `hint` varchar(300) DEFAULT NULL COMMENT '题目提示(可以为对样例输入输出的解释)',
  `time_limit` int(11) NOT NULL DEFAULT '0' COMMENT '时间限制',
  `memory_limit` int(11) NOT NULL DEFAULT '0' COMMENT '内存限制',
  `tag` varchar(200) DEFAULT NULL COMMENT '题目标签',
  `is_special_judge` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否特判 1-特判 2-非特判',
  `special_judge_source` varchar(100) DEFAULT NULL COMMENT '特判程序源代码',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '标准程序',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_difficulty` (`difficulty`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_code` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`problem_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目ID',
	`user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
	`save_code` varchar(100) NOT NULL DEFAULT '' COMMENT '保存代码',
	PRIMARY KEY (`id`),
  	UNIQUE KEY `uniq_user` (`problem_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_collection` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `problem_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目ID',
    `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sys_config` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `sys_key` varchar(100) NOT NULL DEFAULT '' COMMENT '键',
    `sys_value` varchar(100) NOT NULL DEFAULT '' COMMENT '值',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_sys_key` (`sys_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
