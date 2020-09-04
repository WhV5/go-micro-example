/**
* @Author : henry
* @Data: 2020-09-04 10:41
* @Note:
**/

package handler

import (
	"context"
	"github.com/go-micro-example/prime-srv/service"
	"github.com/go-micro-example/proto/prime"
)

type handler struct {
}

func (h handler) GetPrime(ctx context.Context, req *prime.PrimeRequest, rsp *prime.PrimeResponse) error {
	inputs := make([]int64, 0)

	var i int64 = 0
	for ; i <= req.Input; i++ {
		inputs = append(inputs, i)
	}

	rsp.Output = service.GetPrime(inputs...)

	return nil
}

func Handler() handler {
	return handler{}
}
