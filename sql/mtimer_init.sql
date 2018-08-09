CREATE TABLE mtimer_tasks (
	`id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
	`group_id` int(10) UNSIGNED NOT NULL COMMENT '业务组ID',
	`group_name` varchar(20) COMMENT '业务组名称',
	`biz_id` int(10) UNSIGNED NOT NULL COMMENT '业务ID',
	`biz_name` varchar(20) COMMENT '业务名称',
	`cron_time` varchar(20) NOT NULL COMMENT 'cron表达式，用来执行定时任务',
	`status` tinyint(2) NOT NULL COMMENT '任务状态：0:new;1:已加载;2:中断;3:执行成功;4删除:',
	`ip` varchar(16) COMMENT '创建服务器地址',
	`param` varchar(256) COMMENT '回调参数，json格式',
	`ins_num` int(10) UNSIGNED COMMENT '实例分区号：用户多实例部署',
	`excution_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '任务执行时间，对应于cron_time',
	`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	UNIQUE KEY `uniq_mtimer_group_biz` (`group_id`,`biz_id`) USING BTREE
) ENGINE=`InnoDB` COMMENT='定时服务器调度任务表';