
SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for a_menu
-- ----------------------------
DROP TABLE IF EXISTS `a_menu`;
CREATE TABLE `a_menu`
(
    `id`          int          NOT NULL AUTO_INCREMENT,
    `gid`         int          NOT NULL DEFAULT '0',
    `name`        varchar(20)  NOT NULL DEFAULT '',
    `icon`        varchar(100) NOT NULL DEFAULT '',
    `use_vip`     tinyint(1) NOT NULL DEFAULT '0' COMMENT '使用权限限制 0全部',
    `click_type`  tinyint(1) NOT NULL DEFAULT '1' COMMENT '点击类型 1跳转 2调用函数',
    `click_func`  varchar(50)  NOT NULL DEFAULT '0' COMMENT '函数标识 小程序端提前封装',
    `path`        varchar(100) NOT NULL DEFAULT '' COMMENT '打开的页面路径',
    `app_id`      varchar(20)  NOT NULL DEFAULT '' COMMENT '小程序appid',
    `extra_data`  varchar(100) NOT NULL DEFAULT '' COMMENT '需要传递给目标小程序的数据 json',
    `env_version` varchar(100) NOT NULL DEFAULT '' COMMENT '要打开的小程序版本',
    `short_link`  varchar(100) NOT NULL DEFAULT '' COMMENT '小程序链接',
    `sort`        int(11) NOT NULL DEFAULT '0',
    `is_lock`     tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`   int(11) NOT NULL DEFAULT '0',
    `created_at`  timestamp NULL DEFAULT NULL,
    `updated_at`  timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单表';


-- ----------------------------
-- Table structure for a_menu_group
-- ----------------------------
DROP TABLE IF EXISTS `a_menu_group`;
CREATE TABLE `a_menu_group`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `name`       varchar(20) NOT NULL DEFAULT '',
    `sort`       int(11) NOT NULL DEFAULT '0',
    `is_lock`    tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`  int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='菜单分组';

-- ----------------------------
-- Table structure for a_user
-- ----------------------------
CREATE TABLE `a_user`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `openid`     char(32)     NOT NULL DEFAULT '',
    `unionid`    char(32)     NOT NULL,
    `user_name`  varchar(18)  NOT NULL DEFAULT '',
    `nick_name`  varchar(50)  NOT NULL DEFAULT '',
    `password`   char(32)     NOT NULL DEFAULT '',
    `head_img`   varchar(100) NOT NULL DEFAULT '',
    `vip`        int(11) unsigned DEFAULT '0' COMMENT '会员级别',
    `is_lock`    tinyint(1) NOT NULL DEFAULT '1',
    `vip_end_at` timestamp NULL DEFAULT NULL 'vip到期时间',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `oid_idx` (`unionid`,`openid`) USING BTREE,
    KEY          `u_idx` (`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

CREATE TABLE `a_user_contact`
(
    `uid`          int(11) unsigned NOT NULL,
    `superior`     int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级uid',
    `superior_two` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级运营中心',
    `updated_at`   timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户关系';

CREATE TABLE `a_ppt_type`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `name`       varchar(20) NOT NULL DEFAULT '',
    `sort`       int(11) NOT NULL DEFAULT '0',
    `is_lock`    tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`  int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='ppt分类';

CREATE TABLE `a_ppt`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT,
    `tid`          int(11) NOT NULL DEFAULT '0' COMMENT '分类',
    `name`         varchar(20)  NOT NULL DEFAULT '',
    `img_url`      varchar(150) NOT NULL DEFAULT '' COMMENT '封面',
    `desc_content` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
    `sort`         int(11) NOT NULL DEFAULT '0',
    `file_url`     varchar(100) NOT NULL DEFAULT '' COMMENT '文件地址',
    `ai_dou`       int(5) NOT NULL DEFAULT '0' COMMENT '消耗爱享豆',
    `is_lock`      tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`    int(11) NOT NULL DEFAULT '0',
    `created_at`   timestamp NULL DEFAULT NULL,
    `updated_at`   timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `name` (`name`),
    KEY            `tid` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='ppt主表';

CREATE TABLE `a_ppt_content`
(
    `id`      int(11) NOT NULL AUTO_INCREMENT,
    `pid`     int(11) NOT NULL DEFAULT '0',
    `img_url` varchar(100) NOT NULL DEFAULT '',
    `content` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='ppt内容表';

CREATE TABLE `a_content`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `scene`      int(4) NOT NULL DEFAULT '0' COMMENT '场景',
    `name`       varchar(20) NOT NULL DEFAULT '',
    `content`    text,
    `admin_uid`  int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='文本库';

CREATE TABLE `a_feedback`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `uid`        int(11) NOT NULL,
    `content`    varchar(255) DEFAULT '' COMMENT '反馈内容',
    `remark`     varchar(255) DEFAULT '' COMMENT '处理结果备注',
    `admin_uid`  int NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='反馈';

CREATE TABLE `a_ai_dou`
(
    `uid`          int(11) NOT NULL,
    `ai_dou`       int(11) NOT NULL DEFAULT '0' COMMENT '可用爱豆',
    `total_ai_dou` int(11) NOT NULL DEFAULT '0' COMMENT '累计爱豆',
    PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='爱豆表';

CREATE TABLE `a_ai_dou_log`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uid`        int(11) unsigned NOT NULL COMMENT 'uid',
    `source`     tinyint              DEFAULT '0' COMMENT '来源 1ppt 2壁纸 3充值',
    `ai_dou`     int(11) unsigned NOT NULL COMMENT '结算爱豆',
    `ord_sn`     varchar(18) NOT NULL DEFAULT '' COMMENT '订单号',
    `sid`        int(11) NOT NULL DEFAULT '0' COMMENT '资源id',
    `content`    varchar(100)         DEFAULT '' COMMENT '获得说明',
    `direction`  tinyint unsigned DEFAULT '1' COMMENT '类型:1收入,2支出',
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `id_index` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='爱豆变动记录表';

CREATE TABLE `a_wallet`
(
    `uid`           int(11) NOT NULL,
    `balance`       int(11) NOT NULL DEFAULT '0' COMMENT '可用余额',
    `total_balance` int(11) NOT NULL DEFAULT '0' COMMENT '累计余额',
    PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='余额表';

CREATE TABLE `a_wallet_log`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uid`        int(11) unsigned NOT NULL COMMENT 'uid',
    `source`     tinyint(1) DEFAULT '0' COMMENT '来源 1开通VIP ',
    `balance`    int(11) unsigned NOT NULL COMMENT '结算余额',
    `oid`        int(11) NOT NULL DEFAULT '0' COMMENT '订单ID',
    `content`    varchar(100) DEFAULT '' COMMENT '获得说明',
    `direction`  tinyint unsigned DEFAULT '1' COMMENT '类型:1收入,2支出',
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `id_index` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='余额变动记录表';


CREATE TABLE `a_share_resource_type`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(20) NOT NULL DEFAULT '',
    `sort`       int(11) NOT NULL DEFAULT '0',
    `is_lock`    tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`  int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='朋友圈素材分类';

CREATE TABLE `a_share_resource`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `tid`        int(11) unsigned NOT NULL COMMENT '分类id',
    `img_url`    varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
    `content`    varchar(255)          DEFAULT '' COMMENT '文案内容',
    `sort`       int(11) NOT NULL DEFAULT '0',
    `is_lock`    tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`  int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `id_index` (`tid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='朋友圈素材';

CREATE TABLE `a_admin_user`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username`   varchar(50)  NOT NULL COMMENT '用户名',
    `password`   varchar(32)  NOT NULL COMMENT '密码',
    `name`       varchar(100) NOT NULL COMMENT '名称',
    `avatar`     varchar(100) NOT NULL DEFAULT '' COMMENT '头像',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `c_admin_username_unique` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员';
-- admin admin123
INSERT INTO `a_admin_user`(`id`, `username`, `password`, `name`, `avatar`, `created_at`, `updated_at`)
VALUES (1, 'admin', '0192023a7bbd73250516f069df18b500', 'admin', '', NULL, NULL);

CREATE TABLE `a_admin_menu`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `pid`        int(11) NOT NULL DEFAULT '0',
    `name`       varchar(50) NOT NULL DEFAULT '',
    `path`       varchar(50) NOT NULL DEFAULT '',
    `sort`       int(11) NOT NULL DEFAULT '0',
    `icon`       varchar(30) NOT NULL DEFAULT '',
    `is_show`    tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示',
    `is_lock`    tinyint(1) NOT NULL DEFAULT '1',
    `admin_uid`  int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='后台菜单';

CREATE TABLE `a_order`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uid`           int(11) NOT NULL DEFAULT '0',
    `order_no`      char(21) NOT NULL COMMENT '订单号',
    `scene`         tinyint(2) NOT NULL COMMENT '订单场景: 开通vip、充值爱豆...',
    `total_price`   int(11) NOT NULL DEFAULT '0' COMMENT '总价(分)',
    `actual_amount` int(11) NOT NULL DEFAULT '0' COMMENT '实际金额(分)',
    `pay_type`      tinyint(1) DEFAULT '1' COMMENT '订单支付方式 1微信支付 2后台付费 3后台免费',
    `pay_status`    tinyint(1) DEFAULT '1' COMMENT '订单状态 1未支付 2已支付 3失败',
    `remark`        varchar(255) DEFAULT '' COMMENT '订单备注',
    `admin_uid`     int(11) DEFAULT NULL DEFAULT '0',
    `created_at`    timestamp NULL DEFAULT NULL,
    `updated_at`    timestamp NULL DEFAULT NULL,
    `deleted_at`    timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY             `uid_idx` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

CREATE TABLE `a_order_vip`
(
    `id`    int(10) unsigned NOT NULL AUTO_INCREMENT,
    `oid`   int(10) unsigned NOT NULL COMMENT '订单id',
    `level` tinyint(2) unsigned NOT NULL COMMENT 'vip等级',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单详情:开通vip';

CREATE TABLE `a_order_aidou`
(
    `id`     int(10) unsigned NOT NULL AUTO_INCREMENT,
    `oid`    int(10) unsigned NOT NULL COMMENT '订单id',
    `ai_dou` int unsigned NOT NULL COMMENT '爱豆数量',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单详情:充值爱享豆';

SET
FOREIGN_KEY_CHECKS = 1;