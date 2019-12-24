use robot_chat;

drop table if exists chat_user;
CREATE TABLE `chat_user` (
    `uid` VARCHAR(50) NOT NULL,
    `robot_status` INT NOT NULL default 0 COMMENT "机器人开关状态",
    `greet_status` INT NOT NULL default 0 COMMENT "拜年问候状态",
    PRIMARY KEY (`uid`)
)DEFAULT CHARSET=utf8;

drop table if exists chat_keyworld;
CREATE TABLE `chat_keyworld` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `keyid` VARCHAR(300) NOT NULL,
    `back_world` VARCHAR(500) NOT NULL default "等我回来" COMMENT "指定关键词回复",
    PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8;

drop table if exists chat_badworld;
CREATE TABLE `chat_badworld` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `bad_world` VARCHAR(500) NOT NULL default "00" COMMENT "机器人回复错误",
    PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8;

drop table if exists chat_exchangeworld;
CREATE TABLE `chat_exchangeworld` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `keyid` VARCHAR(300) NOT NULL,
    `exchange_world` VARCHAR(500) NOT NULL default "00" COMMENT "机器人回复中需要替换的词",
    PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8;