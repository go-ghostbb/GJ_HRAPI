package enum

import "database/sql/driver"

type LeaveGroupCondition string

const (
	ConditionMonth LeaveGroupCondition = "month"
	ConditionYear  LeaveGroupCondition = "year"
)

func (l *LeaveGroupCondition) Scan(value interface{}) error {
	*l = LeaveGroupCondition(value.(string))
	return nil
}

func (l LeaveGroupCondition) Value() (driver.Value, error) {
	return string(l), nil
}
