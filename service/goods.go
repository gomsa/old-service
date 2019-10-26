package service

import (
	"strings"

	"github.com/go-xorm/xorm"
	pd "github.com/gomsa/old-sql/proto/goods"
)

//Goods 商品仓库接口
type Goods interface {
	List(*pd.Request) ([]*pd.GoodsGroup, error)
}

// GoodsRepository 商品仓库
type GoodsRepository struct {
	Engine *xorm.Engine
}

// List 条件获取商品
func (repo *GoodsRepository) List(req *pd.Request) (goodsGroup []*pd.GoodsGroup, err error) {
	goods := make([]*pd.Goods, 0)
	session := repo.Engine.NewSession()
	if req.ListQuery != nil {
		if req.ListQuery.Where != `` {
			for _, value := range strings.Split(req.ListQuery.Where, ",") {
				session.Where(value)
			}
		}
	}
	err = session.Find(&goods, req.Goods)
	if err != nil {
		return nil, err
	}
	// 查询全部多条码数据
	barcode := make([]*pd.Barcode, 0)
	err = session.Find(&barcode)
	if err != nil {
		return nil, err
	}
	for _, bar := range barcode {
		bar.Hander()
	}
	// 整合多条码数据
	for _, good := range goods {
		good.Hander()
		var bars = make([]*pd.Barcode, 0)
		for _, bar := range barcode {
			if bar.PluCode == good.PluCode {
				bars = append(bars, bar)
			}
		}
		goodsGroup = append(goodsGroup, &pd.GoodsGroup{
			Goods:   good,
			Barcode: bars,
		})
	}
	return goodsGroup, err
}
