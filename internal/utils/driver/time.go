package driver

import (
	"database/sql/driver"
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	"time"
)

type Time int64

func (t *Time) Scan(value interface{}) error {
	temp, _ := time.Parse(time.TimeOnly, value.(time.Time).Format(time.TimeOnly))
	*t = Time(temp.Unix())
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return t.Format(), nil
}

func (t Time) Format() string {
	return time.Unix(int64(t), 0).UTC().Format(time.TimeOnly)
}

func (t Time) Unix() int64 {
	return int64(t)
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
	*t = Time(tt.Unix())

	return nil
}

func (t Time) Time() time.Time {
	return time.Unix(int64(t), 0).UTC()
}
