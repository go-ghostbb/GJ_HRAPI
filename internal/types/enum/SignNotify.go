package enum

import "database/sql/driver"

type SignNotify string

const (
	SignOffPlusNotify SignNotify = "sign-off plus notification"
	NotifyOnly        SignNotify = "notification only"
)

func (l *SignNotify) Scan(value interface{}) error {
	*l = SignNotify(value.(string))
	return nil
}

func (l SignNotify) Value() (driver.Value, error) {
	return string(l), nil
}
