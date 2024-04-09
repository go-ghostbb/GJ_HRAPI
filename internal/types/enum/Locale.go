package enum

import "database/sql/driver"

type Locale string

const (
	ZhTW Locale = "zh_TW"
	ZhCN Locale = "zh_CN"
	EN   Locale = "en"
)

func (l *Locale) Scan(value interface{}) error {
	*l = Locale(value.(string))
	return nil
}

func (l Locale) Value() (driver.Value, error) {
	return string(l), nil
}
