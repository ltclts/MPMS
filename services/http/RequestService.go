package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Param map[string]interface{}
type Response struct {
	Error int8
	Msg   string
	Info  Map
}

type Map map[string]interface{}

func Post(param Param, url string) Response {
	fmt.Println("params_map :", param)
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(param)

	fmt.Println("params_str :", body)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return Response{1, err.Error(), Map{}}
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return Response{2, err.Error(), Map{}}
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	var response Response
	err = json.Unmarshal(respBody, &response)
	fmt.Println(response, resp.Status)
	if err != nil {
		return Response{3, err.Error(), Map{}}
	}
	return Response{0, "ok", response.Info}
}
