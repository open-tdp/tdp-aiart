package server

import (
	"github.com/kardianos/service"
	"github.com/open-tdp/go-helper/dborm"

	"tdp-aiart/cmd/args"
	"tdp-aiart/module/httpd"
	"tdp-aiart/module/migrator"
)

type program struct{}

func (p *program) Start(s service.Service) error {

	svclog.Info("TDP Server start")

	go p.run()
	return nil

}

func (p *program) Stop(s service.Service) error {

	svclog.Info("TDP Server stop")

	return nil

}

func (p *program) run() {

	// 连接数据库
	dborm.Connect(&args.Database)

	// 实施自动迁移
	migrator.Deploy()

	// 启动HTTP服务
	httpd.Daemon()

}
