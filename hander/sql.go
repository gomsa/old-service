package hander

import (
	"context"
	"encoding/json"

	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/go-xorm/xorm"

	pd "github.com/gomsa/old-sql/proto/sql"

	"github.com/gomsa/old-sql/util"
)

// SQL 语句
type SQL struct {
	Engine *xorm.Engine
}

// Query 获取商品列表
func (srv *SQL) Query(ctx context.Context, req *pd.Request, res *pd.Response) (err error) {
	results, err := srv.Engine.QueryString(req.SQL)
	for _, items := range results {
		for k, item := range items {
			items[k] = util.TrimSpace(gbk.Decode(item))
		}
	}
	mjson, _ := json.Marshal(results)
	res.Results = string(mjson)
	return err
}
