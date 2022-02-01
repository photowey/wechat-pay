/*
 * @Author: photowey
 * @Date: 2022-02-01 16:57:21
 * @LastEditTime: 2022-02-01 17:03:51
 * @LastEditors: photowey
 * @Description: okhttp_test.go
 * @FilePath: \wechat-pay\internal\okhttp\okhttp_test.go
 * Copyright (c) 2022 by photowey<photowey@gmail.com>, All Rights Reserved.
 */

package okhttp_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"photowey.com/wechat-pay/internal/okhttp"
)

func TestPost(t *testing.T) {
	response, err := doPostRequest()
	if err != nil {
		t.Logf("the request is error,message is:%s", err)
	}
	t.Logf("the response is:%s", response)
}

func TestGet(t *testing.T) {
	response, err := doGetRequest()
	if err != nil {
		t.Logf("the request is error,message is:%s", err)
	}
	t.Logf("the response is:%s", response)

}

func doPostRequest() (string, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accep"] = "*/*"

	helloPayload := &HelloPayload{
		Name: "photoshark",
		Age:  25,
	}

	bosyStr, _ := json.Marshal(helloPayload)

	response, err := okhttp.Post("http://192.168.217.1:7923/gokhttp/hello", bosyStr, headers)

	return response, err
}

func doGetRequest() (string, error) {
	headers := make(map[string]string)
	headers["Accept"] = "*/*"

	parameters := make(map[string]string)
	parameters["name"] = "photoshark"
	parameters["age"] = "25"

	response, err := okhttp.Get("http://192.168.217.1:7923/gokhttp/hello", parameters, nil)
	if err != nil {
		fmt.Println("err>>", err)
	}
	var helloBody HelloBody
	if err := json.Unmarshal([]byte(response), &helloBody); err == nil {
		fmt.Println("body.id>>", helloBody.Id)
		fmt.Println("body.name>>", helloBody.Name)
		fmt.Println("body.age>>", helloBody.Age)
	} else {
		fmt.Println("err>>", err)
	}

	return response, err
}

type HelloPayload struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type HelloBody struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
