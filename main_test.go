package main

import (
	"context"
	"log"
	"testing"

	"github.com/gomsa/old-sql/hander"
	sPD "github.com/gomsa/old-sql/proto/sql"
	db "github.com/gomsa/old-sql/providers/database"
)

func TestSQL(t *testing.T) {
	// 商品仓库 db 接口实现
	h := hander.SQL{db.Engine}
	req := &sPD.Request{
		SQL: `select TOP 1 * from tBmPlu WHERE PluCode<999 AND (PluStatus=0 or PluStatus=1)`,
	}
	res := &sPD.Response{}
	err := h.Query(context.TODO(), req, res)
	log.Println(res)
	if err != nil {
		t.Errorf("Query goods failed, %v!", err)
	}
}

// func TestGoodsList(t *testing.T) {
// 	// 商品仓库 db 接口实现
// 	repo := &service.GoodsRepository{db.Engine}
// 	h := hander.Goods{repo}
// 	req := &gPD.Request{
// 		ListQuery: &gPD.ListQuery{
// 			Where: `XgDate>='2019-10-01 00:00:00',PluStatus=0 or PluStatus=1`,
// 		},
// 		Goods: &gPD.Goods{
// 			// PluCode: `3301029`,
// 		},
// 	}
// 	res := &gPD.Response{}
// 	err := h.List(context.TODO(), req, res)
// 	log.Println(res)
// 	if err != nil {
// 		t.Errorf("Query goods failed, %v!", err)
// 	}
// }
