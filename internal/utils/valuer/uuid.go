package valuer

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

func UUID[T uuid.UUID | string](v T) driver.Valuer {
	switch t := any(v).(type) {
	case string:
		return &uuidValuer[string]{v: t}
	case uuid.UUID:
		return &uuidValuer[uuid.UUID]{v: t}
	default:
		return &uuidValuer[string]{v: ""}
	}
}

type uuidValuer[T uuid.UUID | string] struct {
	v T
}

func (d *uuidValuer[T]) Value() (driver.Value, error) {
	switch t := any(d.v).(type) {
	case uuid.UUID:
		return t.String(), nil
	case string:
		return t, nil
	default:
		return "", nil
	}
}
