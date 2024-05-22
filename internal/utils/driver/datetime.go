package driver

import (
	"database/sql/driver"
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"time"
)

type Datetime int64

func NewDateTime[T string | time.Time](n T) Datetime {
	t := gbconv.Time(n)
	t, _ = time.Parse(time.DateTime, t.Format(time.DateTime))
	return Datetime(t.Unix())
}

func (d *Datetime) Scan(value interface{}) error {
	if value == nil {
		*d = 0
		return nil
	}
	v := value.(time.Time)
	*d = Datetime(v.Unix())
	return nil
}

func (d Datetime) Value() (driver.Value, error) {
	if d == 0 {
		return nil, nil
	}
	return d.Format(), nil
}

func (d Datetime) Format() string {
	return time.Unix(int64(d), 0).UTC().Format(time.DateTime)
}

func (d Datetime) Unix() int64 {
	return int64(d)
}

func (d Datetime) MarshalJSON() ([]byte, error) {
	if d == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d *Datetime) UnmarshalJSON(data []byte) error {
	dataStr := gbstr.Replace(string(data), `"`, "")
	t, err := time.Parse(time.DateTime, dataStr)
	if err != nil {
		return err
	}
	*d = Datetime(t.Unix())

	return nil
}

func (d Datetime) Time() time.Time {
	return time.Unix(int64(d), 0).UTC()
}

func (d Datetime) After(u Date) bool {
	return d.Time().After(u.Time())
}

func (d Datetime) AddDate(years int, months int, days int) Datetime {
	return Datetime(d.Time().AddDate(years, months, days).Unix())
}
