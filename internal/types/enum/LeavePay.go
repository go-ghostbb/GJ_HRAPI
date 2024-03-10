package enum

import "database/sql/driver"

type LeavePay string

const (
	None LeavePay = "none"
	Half LeavePay = "half"
	All  LeavePay = "all"
)

func (l *LeavePay) Scan(value interface{}) error {
	*l = LeavePay(value.(string))
	return nil
}

func (l LeavePay) Value() (driver.Value, error) {
	return string(l), nil
}
