package response

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func Result(success bool, code int, data interface{}, msg string) Response {
	return Response{
		success,
		code,
		data,
		msg,
	}
}

func Ok(data interface{}, msg string, code int) Response {
	if msg == "" {
		msg = "成功"
	}
	return Result(true, code, data, msg)
}
func Fail(msg string, code int) Response {
	if msg == "" {
		msg = "失败"
	}
	return Result(false, code, nil, msg)
}
