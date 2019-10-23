package database

import (
	"log"

	"github.com/go-xorm/xorm"
	"github.com/gomsa/tools/env"
	txorm "github.com/gomsa/tools/xorm"
)

// Engine 管理包
var (
	// Engine 连接
	Engine *xorm.Engine
)

func init() {
	var err error
	conf := &txorm.Config{
		Driver: env.Getenv("DB_DRIVER", "odbc"),
		// Host 主机地址
		Host: env.Getenv("DB_HOST", "172.168.2.2"),
		// Port 主机端口
		Port: env.Getenv("DB_PORT", "21433"),
		// User 用户名
		User: env.Getenv("DB_USER", "sa"),
		// Password 密码
		Password: env.Getenv("DB_PASSWORD", ""),
		// DbName 数据库名称
		DbName: env.Getenv("DB_NAME", "stmis1"),
		// Charset 数据库编码
		Charset: env.Getenv("DB_CHARSET", "GBK"),
	}
	Engine, err = txorm.Connection(conf)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
}
