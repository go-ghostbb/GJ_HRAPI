package driver

import (
	"database/sql/driver"
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"time"
)

type Datetime int64

func NewDateTime[T string | time.Time](n T, format ...string) Datetime {
	t := gbconv.Time(n, format...)
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

func (d Datetime) Format(layout ...string) string {
	if len(layout) == 0 {
		return time.Unix(int64(d), 0).UTC().Format(time.DateTime)
	}
	return time.Unix(int64(d), 0).UTC().Format(layout[0])
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
	if dataStr == "null" || len(dataStr) == 0 {
		*d = Datetime(0)
		return nil
	}
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

func (d Datetime) After(u Datetime) bool {
	return d.Time().After(u.Time())
}

func (d Datetime) Before(u Datetime) bool {
	return d.Time().Before(u.Time())
}

func (d Datetime) AddDate(years int, months int, days int) Datetime {
	return Datetime(d.Time().AddDate(years, months, days).Unix())
}

func (d Datetime) IsZero() bool {
	return d.Unix() == 0
}
