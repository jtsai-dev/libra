/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:17:19
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-13 10:25:17
 */
package enums

const (
	Success = iota + 1
	Fail
	AuthTokenTimeout
	AuthTokenInvalid
	LoginFail
	ParamsInvalid
	ParamsError
	DataBlank
	DataRepeat
	PermissionError
)

var respCodeflags = map[int]string{
	Success:          "ok",
	Fail:             "fail",
	AuthTokenTimeout: "token timeout",
	AuthTokenInvalid: "token invalid",
	LoginFail:        "the name or password is incorrect",
	ParamsInvalid:    "params invalid",
	ParamsError:      "params error",
	DataBlank:        "data blank",
	DataRepeat:       "data repeat",
	PermissionError:  "no permission",
}

func GetRespCodeDesc(value int) string {
	desc, ok := respCodeflags[value]
	if ok {
		return desc
	}

	return respCodeflags[Fail]
}

const (
	Normal = iota
	Frozen
	Deleted
)

var statusFlags = map[int]string{
	Normal:  "Normal",
	Frozen:  "Frozen",
	Deleted: "Deleted",
}

func GetStatusDesc(value int) string {
	desc, ok := statusFlags[value]
	if ok {
		return desc
	}

	return statusFlags[Fail]
}

const (
	UnSubscribed = iota
	Subscribed
)

var subscribeStatusFlags = map[int]string{
	Subscribed:   "Subscribed",
	UnSubscribed: "UnSubscribed",
}

func GetSubscribeStatusDesc(value int) string {
	desc, ok := subscribeStatusFlags[value]
	if ok {
		return desc
	}

	return subscribeStatusFlags[Fail]
}
