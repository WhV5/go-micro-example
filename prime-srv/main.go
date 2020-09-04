/**
* @Author : henry
* @Data: 2020-09-04 10:48
* @Note:
**/

package main

import (
	"github.com/go-micro-example/prime-srv/handler"
	"github.com/go-micro-example/proto/prime"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 创建服务
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
	)

	// 初始化
	srv.Init()

	// 挂载接口
	_ = prime.RegisterPrimeHandler(srv.Server(), handler.Handler())

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
