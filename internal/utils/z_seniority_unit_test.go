package utils_test

import (
	gbtest "ghostbb.io/gb/test/gb_test"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"hrapi/internal/utils"
	"testing"
)

func TestCalcSeniority(t *testing.T) {
	gbtest.C(t, func(t *gbtest.T) {
		hireDate := gbconv.Time("2020-03-27")
		currDate := gbconv.Time("2024-01-01")
		years, months, days := utils.CalcSeniority(hireDate, currDate)
		t.Assert(years, 3)
		t.Assert(months, 9)
		t.Assert(days, 5)

		hireDate = gbconv.Time("2020-03-27")
		currDate = gbconv.Time("2024-12-31")
		years, months, days = utils.CalcSeniority(hireDate, currDate)
		t.Assert(years, 4)
		t.Assert(months, 9)
		t.Assert(days, 4)

		hireDate = gbconv.Time("2018-05-03")
		currDate = gbconv.Time("2024-09-08")
		years, months, days = utils.CalcSeniority(hireDate, currDate)
		t.Assert(years, 6)
		t.Assert(months, 4)
		t.Assert(days, 5)
	})
}
