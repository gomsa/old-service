package hander

import (
	"context"

	pd "github.com/gomsa/old-sql/proto/goods"
	"github.com/gomsa/old-sql/service"
)

// Goods 商品服务处理
type Goods struct {
	Repo service.Goods
}

// List 获取商品列表
func (srv *Goods) List(ctx context.Context, req *pd.Request, res *pd.Response) (err error) {
	goodsGroup, err := srv.Repo.List(req)
	if err != nil {
		return err
	}
	res.GoodsGroup = goodsGroup
	return err
}
