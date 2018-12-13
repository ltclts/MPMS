package structure



type Response struct {
	Error int8   `json:"error"`
	Msg   string `json:"msg"`
	Info  Map    `json:"info"`
}

type Map map[string]interface{}