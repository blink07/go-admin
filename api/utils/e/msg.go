package e

var MsgFlags = map[int]string{
	SUCCESS:"OK",
	ERROR: "Fail",
	INVALID_PARAMS: "请求参数错误",


	DATA_INSERT_INOT_FAIL:"数据插入失败",
	USERNAME_OR_PASSWORD:"用户名或密码错误",
	TOKEN_NOT_VALID: "token无效",

	//file
	FILE_PARAM_GET_FAIL:"文件参数获取失败",
	FILE_NOT_STANDARD:"文件不符合规范",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:"检查图片失败",
	FILE_UPLOAD_FAIL:"文件上传失败",
}


func GetErrMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR]
}
