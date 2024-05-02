package driver

import (
	"database/sql/driver"
)

type String string

func NewString(s string) String {
	return String(s)
}

func (s *String) Scan(value interface{}) error {
	v := value.(string)
	*s = String(v)
	return nil
}

func (s String) Value() (driver.Value, error) {
	return string(s), nil
}
