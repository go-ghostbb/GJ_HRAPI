package mathx_test

import (
	gbtest "ghostbb.io/gb/test/gb_test"
	"hrapi/internal/utils/mathx"
	"testing"
)

func TestRoundx(t *testing.T) {
	gbtest.C(t, func(t *gbtest.T) {
		test1 := 6.7
		test2 := 6.4
		test3 := 6.5
		test4 := 6.72
		test5 := 6.45
		test6 := 6.56
		test7 := 6.458

		t.Assert(mathx.Roundx(test1, 0), 7.0)
		t.Assert(mathx.Roundx(test2, 0), 6.0)
		t.Assert(mathx.Roundx(test3, 0), 6.5)
		t.Assert(mathx.Roundx(test4, 1), 6.7)
		t.Assert(mathx.Roundx(test5, 1), 6.45)
		t.Assert(mathx.Roundx(test6, 1), 6.6)
		t.Assert(mathx.Roundx(test7, 1), 6.45)
	})
}
