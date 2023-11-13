package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNotLogin
	CodeTokenFormatErr
	CodeTokenParseErr
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "param is invalid",
	CodeUserExist:       "user is already existed",
	CodeUserNotExist:    "user is not existed",
	CodeInvalidPassword: "password is invalid",
	CodeServerBusy:      "server is busy",
	CodeNotLogin:        "login is necessary access for this page",
	CodeTokenFormatErr:  "token have format err",
	CodeTokenParseErr:   "token parse err",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
