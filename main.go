package main

import (
	"flag"
	"fmt"
	"log"

	"newspeak-chat/internal/config"
	"newspeak-chat/internal/handler"
	"newspeak-chat/internal/svc"

	"github.com/joho/godotenv"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/message-api.yaml", "the config file")

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ 没有找到 .env 文件")
	}

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 设置超时时间为5分钟
	c.RestConf.Timeout = 300000
	server := rest.MustNewServer(c.RestConf, rest.WithCors(
		"*",
		"Content-Type, Accept, Authorization",
		"GET, POST, PUT, DELETE, OPTIONS",
		"true",
	))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
