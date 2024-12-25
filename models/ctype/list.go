package ctype

import (
	"database/sql/driver"
	"github.com/sirupsen/logrus"
	"strings"
)

type List []string

func (j *List) Scan(value interface{}) error {
	val, ok := value.([]uint8)
	logrus.Info("ok:", ok)
	if ok {
		logrus.Info("val:", val)
		*j = strings.Split(string(val), ",")
	}
	return nil
}

func (j List) Value() (driver.Value, error) {
	return strings.Join(j, ","), nil
}
