# RobotChat
机器人聊天

恰逢春节之际，回山中过年，穷山僻壤无法与外界交流，恐佳节不能与好友相送祝福，遂写此程序，使微信与之挂钩，能自动答复之。

post  /robotchat 
传参 
userid  //随便，没有要求
content //对话内容

回参
userid  //传过来得uid
content //机器人回话

###################################################
################ config.toml 文件 #################
###################################################
#被忽略提交的文件内容
[logger]
opLogPath = "./logs"
errLogPath = "./logs"
logLevel = "trace"
logFileMaxSize = 500
logFileMaxBackups = 10
logFileMaxAge = 30

[mysql]
user = "user"
database = "database"
password = "*****"
address = "127.0.0.1:3306"

[redis]
password = ""
address = "127.0.0.1:6379"
maxIdle = 500
maxActive = 500


[requrl]
xiaoyingurl      = "url1"
liaotiannvpuurl = "url2"
englishurl      = ""
###################################################
###################################################