package dto

type ResponseDto struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func Success(data interface{}) ResponseDto {
	msg := ResponseDto{
		Code:  0,
		Msg:   "SUCCESS",
		Data:  data,
		Count: 0,
	}
	return msg
}
func SuccessOk() ResponseDto {
	msg := ResponseDto{
		Code:  0,
		Msg:   "SUCCESS",
		Data:  "SUCCESS",
		Count: 0,
	}
	return msg
}
func SuccessMessage(message string, data interface{}) ResponseDto {
	msg := ResponseDto{
		Code:  0,
		Msg:   message,
		Data:  data,
		Count: 0,
	}
	return msg
}

func SuccessCount(data interface{}, count int64) ResponseDto {
	msg := ResponseDto{
		Code:  0,
		Msg:   "SUCCESS",
		Data:  data,
		Count: count,
	}
	return msg
}

func Fail(msg string) ResponseDto {
	msgObj := ResponseDto{
		Code: 1,
		Msg:  msg,
	}
	return msgObj
}
func FailMsg(msg string) ResponseDto {
	msgObj := ResponseDto{
		Code: 1,
		Msg:  msg,
	}
	return msgObj
}
func FailCode(code int) ResponseDto {
	msgObj := ResponseDto{
		Code: code,
		Msg:  "error",
	}
	return msgObj
}
func FailCodeAndMsg(code int, msg string) ResponseDto {
	msgObj := ResponseDto{
		Code: code,
		Msg:  msg,
	}
	return msgObj
}

type NullDto struct {
}
