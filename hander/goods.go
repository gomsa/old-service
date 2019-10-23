package hander

import (
	"context"

	pd "github.com/gomsa/old-service//proto/goods"
	"github.com/gomsa/old-service//service"
)

// Goods 商品服务处理
type Goods struct {
	Repo service.Goods
}

// List 获取商品列表
func (srv *Goods) List(ctx context.Context, req *pd.Request, res *gpd.Response) (err error) {
	good, err := srv.Repo.List(req)
	if err != nil {
		return err
	}
	res.Good = good
	return err
}
