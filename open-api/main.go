/**
* @Author : henry
* @Data: 2020-09-04 11:49
* @Note:
**/

package main

import (
	"context"
	"encoding/json"
	"github.com/go-micro-example/proto/prime"
	"github.com/go-micro-example/proto/sum"
	"github.com/micro/go-micro/v2"
	api "github.com/micro/go-micro/v2/api/proto"
	"strconv"
)

var (
	srvClient   sum.SumService
	primeClient prime.PrimeService
)

type Open struct {
}

func (open Open) Fetch(ctx context.Context, req *api.Request, rsp *api.Response) error {
	sumInputStr := req.Get["sum"].Values[0]
	primeInputStr := req.Get["prime"].Values[0]

	sunInput, _ := strconv.ParseInt(sumInputStr, 10, 10)
	primeInput, _ := strconv.ParseInt(primeInputStr, 10, 10)

	sumReq := &sum.SumRequest{Input: sunInput}
	primeReq := &prime.PrimeRequest{Input: primeInput}

	// 调用客户端
	sumRsp, _ := srvClient.GetSum(ctx, sumReq)
	primeRsp, _ := primeClient.GetPrime(ctx, primeReq)

	ret, _ := json.Marshal(map[string]interface{}{
		"sum":   sumRsp,
		"prime": primeRsp,
	})

	rsp.Body = string(ret)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.api.open"),
	)

	srvClient = sum.NewSumService("", service.Client())
	primeClient = prime.NewPrimeService("go.micro.learning.srv.prime", service.Client())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
