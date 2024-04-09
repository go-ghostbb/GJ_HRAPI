package driver

import (
	"database/sql/driver"
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	"time"
)

type Date time.Time

func (d *Date) Scan(value interface{}) error {
	*d = Date(value.(time.Time))
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Format(), nil
}

func (d Date) Format() string {
	return time.Time(d).Format(time.DateOnly)
}

func (d Date) Unix() int64 {
	return time.Time(d).Unix()
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	dataStr := gbstr.Replace(string(data), `"`, "")
	t, err := time.Parse(time.DateOnly, dataStr)
	if err != nil {
		return err
	}
	*d = Date(t)

	return nil
}
