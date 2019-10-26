package goods

import (
	"time"

	"github.com/gomsa/old-service/util"
)

// TimeLayout 转换字符
const TimeLayout = "2006-01-02 15:04:05"

var (
	local = time.FixedZone("CST", 8*3600)
)

// TableName 对应数据库名称
func (goodsGroup *GoodsGroup) TableName() string {
	return `tBmPlu`
}

// TableName 对应数据库名称
func (goods *Goods) TableName() string {
	return `tBmPlu`
}

// Hander 数据预处理
func (g *Goods) Hander() {
	g.PluCode = util.TrimSpace(g.PluCode)
	g.PluName = util.TrimSpace(g.PluName)
	g.BarCode = util.TrimSpace(g.BarCode)
	g.Unit = util.TrimSpace(g.Unit)
	g.Spec = util.TrimSpace(g.Spec)
	g.Weigt = util.TrimSpace(g.Weigt)
	g.DepCode = util.TrimSpace(g.DepCode)
	g.Produce = util.TrimSpace(g.Produce)
	g.Grade = util.TrimSpace(g.Grade)
	g.SupCode = util.TrimSpace(g.SupCode)
	g.ClsCode = util.TrimSpace(g.ClsCode)
	g.BrandCode = util.TrimSpace(g.BrandCode)
	g.HJPrice = util.TrimSpace(g.HJPrice)
	g.WJPrice = util.TrimSpace(g.WJPrice)
	g.SPrice = util.TrimSpace(g.SPrice)
	g.HyPrice = util.TrimSpace(g.HyPrice)
	g.PfPrice = util.TrimSpace(g.PfPrice)
	g.JTaxRate = util.TrimSpace(g.JTaxRate)
	g.XTaxRate = util.TrimSpace(g.XTaxRate)
	g.CgyCode = util.TrimSpace(g.CgyCode)
	g.CgyName = util.TrimSpace(g.CgyName)
	g.XgDate = util.TrimSpace(g.XgDate)
	g.LrDate = util.TrimSpace(g.LrDate)
	g.UserCode = util.TrimSpace(g.UserCode)
	g.UserName = util.TrimSpace(g.UserName)
	g.PackPriceType = util.TrimSpace(g.PackPriceType)
	g.PackPriceRate = util.TrimSpace(g.PackPriceRate)
	g.PluStatus = util.TrimSpace(g.PluStatus)
}

// TableName 对应数据库名称
func (barcode *Barcode) TableName() string {
	return `tBmMulBar`
}

// Hander 数据预处理
func (b *Barcode) Hander() {
	b.BarCode = util.TrimSpace(b.BarCode)
	b.PluCode = util.TrimSpace(b.PluCode)
	b.DepCode = util.TrimSpace(b.DepCode)
	b.PluName = util.TrimSpace(b.PluName)
	b.PluAbbrName = util.TrimSpace(b.PluAbbrName)
	b.Spec = util.TrimSpace(b.Spec)
}
