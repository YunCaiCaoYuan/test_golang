package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual的描述", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})
}

func TestStringSliceEqual2(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		
		Convey("true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("true when a ＝= nil  && b ＝= nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})
	})
}
