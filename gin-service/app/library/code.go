package library

const (
	SUCCESS       = 1
	ERROR         = 0
	InvalidParams = 40001

	UserNotfound = 10001
	UserIsLock   = 10002
	UserIsDel    = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004

	SystemRateLimitErr  = 30001
	SystemWxMiniIsCheck = 30002
	SystemIsClose       = 30003
	SystemVersionUpdate = 30004
)

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	SystemRateLimitErr:         "不能频繁操作",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	UserNotfound:               "用户不存在",
	UserIsLock:                 "账户已被锁定，当前无法使用，请联系站长！",
	UserIsDel:                  "账户已被拉黑，当前无法使用，请联系站长！",
	SystemWxMiniIsCheck:        "小程序审核中",
	SystemIsClose:              "系统维护中",
	SystemVersionUpdate:        "系统已更新",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
