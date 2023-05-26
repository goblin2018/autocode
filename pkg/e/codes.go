package e

// http 200 ok
// data: {
//		code: 200
//    msg :
//    data:
// }

var (
	OK = New(200, "ok")

	InvalidParams = New(400, "invalid params: ")
	Forbidden     = New(403, "forbidden: ")
	SystemError   = New(500, "system error")

	Unknown            = New(1000, "unknown error")
	DBError            = New(1001, "db error: ")
	QueryTooFrequently = New(1002, "请求太频繁")
	TokenError         = New(1003, "token error: ")

	UserAlreadyExists = New(2003, "user aleady exists: ")
	UserNotExists     = New(2004, "user not exists: ")
	InvalidPassword   = New(2005, "invalidPassword: ")
)
