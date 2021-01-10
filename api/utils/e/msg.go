package e

var MsgFlags = map[int]string{
	SUCCESS:"OK",
	ERROR: "Fail",
	INVALID_PARAMS: "请求参数错误",


	DATA_INSERT_INOT_FAIL:"数据插入失败",
	USERNAME_OR_PASSWORD:"用户名或密码错误",
	TOKEN_NOT_VALID: "token无效",
}


func GetErrMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR]
}
