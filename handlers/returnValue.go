package handlers

var errorMap = map[int]string{
	0: "success",
	1: "File not found",
	2: "Permission denied",
	3: "Internal server error",
	// 添加更多的错误代码和对应的消息
}

type ErrorMessageType struct {
	Code    int
	Message string
}

var MessageOk = ErrorMessageType{Code: 0, Message: "success"}

// 错误代码 1xx为游戏服务器错误
var MessageErrorExecutionFailed = ErrorMessageType{Code: 101, Message: "execution failed"}

// 666其他错误

var MessageErrorUnknown = ErrorMessageType{Code: 666, Message: "unknown error"}
