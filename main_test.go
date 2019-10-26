package main

import (
	"context"
	"log"
	"testing"

	"github.com/gomsa/old-sql/hander"
	gPD "github.com/gomsa/old-sql/proto/goods"
	db "github.com/gomsa/old-sql/providers/database"
	"github.com/gomsa/old-sql/service"
)

func TestGoodsList(t *testing.T) {
	// 商品仓库 db 接口实现
	repo := &service.GoodsRepository{db.Engine}
	h := hander.Goods{repo}
	req := &gPD.Request{
		ListQuery: &gPD.ListQuery{
			Where: `XgDate>='2019-10-01 00:00:00',PluStatus=0 or PluStatus=1`,
		},
		Goods: &gPD.Goods{
			// PluCode: `3301029`,
		},
	}
	res := &gPD.Response{}
	err := h.List(context.TODO(), req, res)
	log.Println(res)
	if err != nil {
		t.Errorf("Query goods failed, %v!", err)
	}
}
