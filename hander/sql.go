package hander

import (
	"context"
	"encoding/json"

	pd "github.com/gomsa/old-sql/proto/sql"

	"github.com/go-xorm/xorm"
)

// SQL 语句
type SQL struct {
	Engine *xorm.Engine
}

// Query 获取商品列表
func (srv *SQL) Query(ctx context.Context, req *pd.Request, res *pd.Response) (err error) {
	results, err := srv.Engine.Query(req.SQL)
	mjson, _ := json.Marshal(results)
	res.Results = string(mjson)
	return err
}
