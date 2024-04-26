package driver

import (
	"database/sql/driver"
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"time"
)

type Date int64

func NewDate[T string | time.Time](n T) Date {
	t := gbconv.Time(n)
	t, _ = time.Parse(time.DateOnly, t.Format(time.DateOnly))
	return Date(t.Unix())
}

func (d *Date) Scan(value interface{}) error {
	v := value.(time.Time)
	*d = Date(v.Unix())
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Format(), nil
}

func (d Date) Format() string {
	return time.Unix(int64(d), 0).UTC().Format(time.DateOnly)
}

func (d Date) Unix() int64 {
	return int64(d)
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
	*d = Date(t.Unix())

	return nil
}

func (d Date) DateTime(t Time) time.Time {
	result, _ := time.Parse(time.DateTime, d.Format()+" "+t.Format())
	return result
}

func (d Date) Time() time.Time {
	return time.Unix(int64(d), 0).UTC()
}

func (d Date) After(u Date) bool {
	return d.Time().After(u.Time())
}

func (d Date) AddDate(years int, months int, days int) Date {
	return Date(d.Time().AddDate(years, months, days).Unix())
}
