package main

import (
	// 公共引入
	k8s "github.com/micro/examples/kubernetes/go/micro"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/gomsa/old-service/hander"
	db "github.com/gomsa/old-service/providers/database"
	"github.com/gomsa/old-service/service"

	goodsPB "github.com/gomsa/old-service/proto/goods"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 商品服务实现
	repo := &service.GoodsRepository{db.Engine}
	goodsPB.RegisterGoodssHandler(srv.Server(), &hander.Goods{repo})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
