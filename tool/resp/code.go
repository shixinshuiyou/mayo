package resp

const (
	CodeGrpcFail            = "FAIL"
	CodeGrpcSuccess         = "SUCCESS"
	CodeSuccess             = 0
	CodeUnknownError        = -1
	CodeLoginError          = 40001
	CodeLoginSessionIdError = 40002
)

var MsgCodeMap = map[int]string{
	CodeUnknownError:        "未知错误",
	CodeLoginError:          "login失败",
	CodeLoginSessionIdError: "无效的sessionId",
	CodeSuccess:             "操作成功",
}
