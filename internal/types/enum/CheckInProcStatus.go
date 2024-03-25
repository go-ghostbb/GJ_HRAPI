package enum

import "database/sql/driver"

type CheckInProcStatus string

const (
	ProcNormal   CheckInProcStatus = "normal"
	NotProcessed CheckInProcStatus = "not processed"
	Processed    CheckInProcStatus = "processed"
)

func (e *CheckInProcStatus) Scan(value interface{}) error {
	*e = CheckInProcStatus(value.(string))
	return nil
}

func (e CheckInProcStatus) Value() (driver.Value, error) {
	return string(e), nil
}
