package uitl

import (
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

// TimeLayout 转换字符
const TimeLayout = "2006-01-02 15:04:05"

// XormWhere Xorm 条件查询
func XormWhere(session *xorm.Session, value string) *xorm.Session {
	where := strings.Split(value, ",")
	for _, whe := range where {
		split := strings.Split(whe, "|")
		var val interface{}
		param := split[0] + split[1] + `?`
		switch split[2] {
		case `string`:
			val = split[3]
		case `time`:
			val, _ = time.ParseInLocation(TimeLayout, split[3], time.Local)
		}
		session.Where(param, val)
	}
	return session
}
