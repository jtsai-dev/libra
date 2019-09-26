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
