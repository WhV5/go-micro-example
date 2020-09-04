/**
* @Author : henry
* @Data: 2020-09-04 15:03
* @Note:
**/

package main

import (
	"context"
	"encoding/json"
	logProto "github.com/go-micro-example/proto/log"
	proto "github.com/go-micro-example/proto/log"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.LogEvt) (err error) {
	// 业务逻辑
	log.Logf("[sub] 收到日志:", evt.Msg)

	return nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.learning.srv.log"))

	service.Init()

	micro.RegisterSubscriber("go.micro.learning.topic.log", service.Server(), &Sub{})

	if err := service.Run(); err != nil {
		panic(err)
	}
}
