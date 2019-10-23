package main

import (
	"context"
	"log"
	"testing"

	"github.com/gomsa/old-service/hander"
	gPD "github.com/gomsa/old-service/proto/goods"
	db "github.com/gomsa/old-service/providers/database"
	"github.com/gomsa/old-service/service"
)

func TestOrderPaySum(t *testing.T) {
	// 商品仓库 db 接口实现
	repo := &service.GoodsRepository{db.Engine}
	h := hander.Goods{repo}
	req := &gPD.Request{
		Goods: &gPD.Goods{
			PluCode:`123`,
		},
	}
	res := &gPD.Response{}
	err := h.List(context.TODO(), req, res)
	log.Println(res.Total)
	if err != nil {
		t.Errorf("Query goods failed, %v!", err)
	}
}
