package http

type Response struct {
	Code    	int32       `json:"code"`
	Message		string      `json:"message"`
	Timestamp	int64       `json:"timestamp"`
	Data  		interface{} `json:"data"`
}

func (res *Response) Success(data interface{}) pkg.JSON {
	//res.Timestamp =
	//
	//res := Response{
	//	Code:    util.Success,
	//	Messag:     util.ErrorMsg(util.Success),
	//	Timestamp: datetime.Timestamp(),
	//	Data:    data,
	//}
	//
	//return pkg.JSON{
	//	Data: res,
	//}
}

func (Response) Error(err error, code ...int) pkg.JSON {
	//c := int(util.DefaultErr)
	//if len(code) > 0 {
	//	c = code[0]
	//}
	//
	//res := Resp{
	//	Code:    c,
	//	Msg:     err.Error(),
	//	NowTime: datetime.Timestamp(),
	//	Data:    nil,
	//}
	//
	//return pkg.JSON{
	//	Data: res,
	//}
}