package service

import (
	"fmt"

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
	engine := repo.Engine.Table("tBmPlu")
	if req.ListQuery != nil {
		engine = engine.Where("XgDate?", &req.ListQuery.XgDate)
	}
	err = engine.Find(&goods, req.Goods)
	fmt.Println(req, goods)
	// goods := &models.Goods{
	// 	BarCode: good.BarCode,
	// }
	// err := GoodsInfo(repo.Engine, goods)
	// if err != nil {
	// 	return nil, err
	// }
	// g := &gpd.Good{
	// 	Id:       goods.Id,
	// 	Name:     goods.Name,
	// 	BarCode:  goods.BarCode,
	// 	Price:    goods.Price,
	// 	Standard: goods.Standard,
	// }
	// return g, err
	return goods, err
}
