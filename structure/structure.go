package structure

type Response struct {
	Error int8              `json:"error"`
	Msg   string            `json:"msg"`
	Info  StringToObjectMap `json:"info"`
}

type StringToObjectMap map[string]interface{}

type Uint8ToObjectMap map[uint8]interface{}

type Uint8ToStringMap map[uint8]string

type Uint8ToInt64 map[uint8]int64

type Array []interface{}
