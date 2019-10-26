package main

import (
	// 公共引入
	k8s "github.com/micro/examples/kubernetes/go/micro"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/gomsa/old-sql/hander"
	db "github.com/gomsa/old-sql/providers/database"

	PB "github.com/gomsa/old-sql/proto/sql"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 商品服务实现
	PB.RegisterSQLSHandler(srv.Server(), &hander.SQL{db.Engine})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}