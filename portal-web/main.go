/**
* @Author : henry
* @Data: 2020-09-04 10:55
* @Note:
**/

package main

import (
	"context"
	"github.com/go-micro-example/proto/sum"
	"github.com/micro/go-micro/v2/web"
	"net/http"
	"strconv"
)

var (
	srvClient sum.SumService
)

func main() {
	service := web.NewService(
		web.Name("go.micro.learning.web.portal"),
		web.Address(":8888"),
		web.StaticDir("html"),
	)

	service.Init()

	srvClient = sum.NewSumService("go.micro.learning.srv.sum", service.Options().Service.Client())
	service.HandleFunc("/sum", Sum)

	service.Run()
}

func Sum(w http.ResponseWriter, r *http.Request) {
	inputString := r.URL.Query().Get("input")
	input, _ := strconv.ParseInt(inputString, 10, 10)

	req := &sum.SumRequest{Input: input}

	// 客户端
	rsp, err := srvClient.GetSum(context.Background(), req)
	if err != nil {
		// ignore
	}

	w.Write([]byte(strconv.Itoa(int(rsp.Output))))
}
