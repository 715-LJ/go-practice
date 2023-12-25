package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data"`
}

type ChaosSimulateInject_Response struct {
	Code   int
	Result string
	Reason string
	Data   interface{}
}

func main() {
	url := "http://1.1.1.1/chaos/simulate/inject"
	method := "POST"
	reqStr, _ := json.Marshal("{}")

	payload := strings.NewReader(string(reqStr))
	httpClient(url, method, payload)
}

func httpClient(url string, method string, payload *strings.Reader) (*ChaosSimulateInject_Response, error) {

	var response Response
	var InjectRes = &ChaosSimulateInject_Response{Code: 0, Result: "", Reason: "", Data: nil}

	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("请求Address：%S 错误：", err)

	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &response); err != nil {
		return InjectRes, err
	}
	return nil, nil
}
