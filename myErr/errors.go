package myErr

const (
	RespOK              = 0
	RespError           = 4001001
	RespErrServer       = 5001001
	RespErrNetWork      = 5001002
	RespErrNoUser       = 5002001
	RespErrInvalidParam = 5002002
	RespErrMysql        = 5006001
	RespErrRedis        = 5006003

	ROBOTCLOSE = 9009001
)

var MsgFlags = map[int]string{
	RespOK:              "ok",
	RespError:           "error",
	RespErrServer:       "server err",
	RespErrNetWork:      "network err",
	RespErrNoUser:       "no user",
	RespErrInvalidParam: "InvalidParam err",
	RespErrMysql:        "mysql err",
	RespErrRedis:        "redis err",
	ROBOTCLOSE:          "ROBOTCLOSE",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[RespError]
}
