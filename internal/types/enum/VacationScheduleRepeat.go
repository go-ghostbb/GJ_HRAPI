package enum

import "database/sql/driver"

type VacationScheduleRepeat string

const (
	RepeatNone  VacationScheduleRepeat = "none"
	RepeatWeek  VacationScheduleRepeat = "week"
	RepeatMonth VacationScheduleRepeat = "month"
	RepeatYear  VacationScheduleRepeat = "year"
)

func (v *VacationScheduleRepeat) Scan(value interface{}) error {
	*v = VacationScheduleRepeat(value.(string))
	return nil
}

func (v VacationScheduleRepeat) Value() (driver.Value, error) {
	return string(v), nil
}
