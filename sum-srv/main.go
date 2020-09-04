/**
* @Author : henry
* @Data: 2020-09-04 10:07
* @Note:
**/

package main

import (
	"context"
	"encoding/json"
	logProto "github.com/go-micro-example/proto/log"
	"github.com/go-micro-example/proto/sum"
	"github.com/go-micro-example/sum-srv/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// 创建服务
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
	)

	// 初始化
	srv.Init(
		micro.BeforeStart(func() error {
			log.Log("[srv] 启动前的日志")
			return nil
		}),

		micro.AfterStart(func() error {
			log.Log("[srv] 启动后的日志")
			return nil
		}),
	)

	// 挂载接口
	_ = sum.RegisterSumHandler(srv.Server(), handler.Handler())

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

// 日志Wrapper
// 通过Broxer异步消息把日志推送到日志服务
func reqLogger(cli client.Client) server.HandlerWrapper {
	pub := micro.NewPublisher("go.micro.learning.topic.log", cli)
	// 初始化动作
	return func(handleFunc server.HandlerFunc) server.HandlerFunc {
		// 中间按动作
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Info("请求, 准备发送日志")
			evt := logProto.LogEvt{
				Msg: "Hello",
			}

			body, _ := json.Marshal(evt)

			pub.Publish(ctx, &broker.Message{
				Header: map[string]string{
					"serviceName": "sum",
				},
				Body: body,
			})

			return handleFunc(ctx, req, rsp)
		}
	}
}
