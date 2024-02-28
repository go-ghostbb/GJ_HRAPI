package password_test

import (
	gbtest "ghostbb.io/gb/test/gb_test"
	"hrapi/internal/utils/password"
	"testing"
)

func Test_Hash(t *testing.T) {
	gbtest.C(t, func(t *gbtest.T) {
		t.Assert(password.Check(password.Hash("123456"), "123456"), true)
	})
}

func Test_Check(t *testing.T) {
	gbtest.C(t, func(t *gbtest.T) {
		t.Assert(password.Check(`$2a$10$3zL.F5U8QrWPy6y4Z.WIouEwTZwBtyTA9la5EFzrEaUMimi5ngQc2`, "123456"), true)
		t.Assert(password.Check(`$2a$10$3zL.F5U8QrWPy6y4Z.WIouEwTZwBtyTA9la5EFzrEaUMimi5ngQc2`, "654321"), false)
	})
}
