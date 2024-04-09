package driver

import (
	"database/sql/driver"
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"time"
)

type Time time.Time

func (t *Time) Scan(value interface{}) error {
	temp, _ := time.Parse(time.TimeOnly, value.(time.Time).Format(time.TimeOnly))
	*t = Time(gbconv.Time(temp))
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return t.Format(), nil
}

func (t Time) Format() string {
	return time.Time(t).Format(time.TimeOnly)
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format())), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	dataStr := gbstr.Replace(string(data), `"`, "")
	tt, err := time.Parse(time.TimeOnly, dataStr)
	if err != nil {
		return err
	}
	*t = Time(tt)

	return nil
}
