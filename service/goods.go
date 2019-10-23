package service

import (
	"github.com/gomsa/tools/uitl"

	"github.com/go-xorm/xorm"
	pd "github.com/gomsa/old-service/proto/goods"
)

//Goods 商品仓库接口
type Goods interface {
	List(*pd.Request) ([]*pd.Goods, error)
}

// GoodsRepository 商品仓库
type GoodsRepository struct {
	Engine *xorm.Engine
}

// List 条件获取商品
func (repo *GoodsRepository) List(req *pd.Request) (goods []*pd.Goods, err error) {
	session := repo.Engine.Table("tBmPlu")
	if req.ListQuery != nil {
		if req.ListQuery.Where != `` {
			session = uitl.XormWhere(session, req.ListQuery.Where)
		}
	}
	err = session.Find(&goods, req.Goods)
	return goods, err
}
