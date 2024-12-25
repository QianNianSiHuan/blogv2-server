package ctype

import (
	"database/sql/driver"
	"github.com/sirupsen/logrus"
	"strings"
)

type List []string

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *List) Scan(value interface{}) error {
	val, ok := value.([]uint8)
	logrus.Info("ok:", ok)
	if ok {
		logrus.Info("val:", val)
		*j = strings.Split(string(val), ",")
	}
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j List) Value() (driver.Value, error) {
	return strings.Join(j, ","), nil
}
