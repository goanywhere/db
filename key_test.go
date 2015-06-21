package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkNewKey(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NewKey()
	}
}

func TestNewKey(t *testing.T) {
	key := NewKey()
	Convey("db.Key basic test", t, func() {
		So(len(key), ShouldEqual, 12)
		So(len(key.Hex()), ShouldEqual, 24)
	})
}
