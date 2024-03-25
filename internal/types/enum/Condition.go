package enum

import "database/sql/driver"

type Operator string

const (
	Eq  Operator = "=="
	Gt  Operator = ">"
	Gte Operator = ">="
	Lt  Operator = "<"
	Lte Operator = "<="
)

func (e *Operator) Scan(value interface{}) error {
	*e = Operator(value.(string))
	return nil
}

func (e Operator) Value() (driver.Value, error) {
	return string(e), nil
}
