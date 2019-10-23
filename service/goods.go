package service

import (
	"fmt"
	pd "github.com/gomsa/old-service/proto/goods"
	"github.com/go-xorm/xorm"
)

//Goods 商品仓库接口
type Goods interface {
	List(*pd.Request) (*pd.Response, error)
}

// GoodsRepository 商品仓库
type GoodsRepository struct {
	Engine *xorm.Engine
}

// List 条件获取商品
func (repo *GoodsRepository) List(req *pd.Request) (res *pd.Response, err error) {
	fmt.Println(req,res)
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
	return res, err
}
