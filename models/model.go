package models

type FinalState struct {
	Code int
	Mes  interface{}
	Data interface{}
}

func Success(data interface{}) (f FinalState) {
	f.Code = 0
	f.Mes = "success"
	f.Data = data
	return f
}

func CharFail(data interface{}) (f FinalState) {
	f.Code = 4001
	f.Mes = "字符不服合规则"
	f.Data = data
	return f
}
