package e

var MsgFlags = map[int]string{
	SUCCESS:"OK",
	ERROR: "Fail",
	INVALID_PARAMS: "请求参数错误",
}


func GetErrMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR]
}
