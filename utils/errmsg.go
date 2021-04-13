package utils

const (
	SUCCSE = 200
	ERROR  = 500

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	ERROR_PASSWORD_EDIT    = 1009
	ERROR_PASSWORD_NOTSAME = 1010
	ERROR_PASSWORD_FAIL    = 1011
	ERROR_USER_FAIL        = 1012
	ERROR_BOTH_WRONG       = 1013
	ERROR_PASSWORD_EDFL    = 1014
)

var codeMsg = map[int]string{
	SUCCSE:                 "成功辣",
	ERROR:                  "失败辣",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_PASSWORD_EDIT:    "密码修改成功",
	ERROR_PASSWORD_FAIL:    "密码加密失败",
	ERROR_PASSWORD_NOTSAME: "输入的密码不一致",
	ERROR_USER_FAIL:        "注册失败",
	ERROR_BOTH_WRONG:       "帐号或密码错误",
	ERROR_PASSWORD_EDFL:    "密码修改失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
