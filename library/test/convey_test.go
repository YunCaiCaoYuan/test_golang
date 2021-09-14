package test

import (
	. "github.com/prashantv/gostub"
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
/*
func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual的描述", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}

func TestStringSliceEqual(t *testing.T) {
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
 */

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		Convey("should return true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return true when a ＝= nil  && b ＝= nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a ＝= nil  && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a != nil  && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}

/*
func TestFuncDemo(t *testing.T) {
	Convey("TestFuncDemo", t, func() {
		Convey("for succ", func() {
			stubs := Stub(&num, 150)
			defer stubs.Reset()
			stubs.StubFunc(&Exec,"xxx-vethName100-yyy", nil)
			var liLei = `{"name":"LiLei", "age":"21"}`
			stubs.StubFunc(&adapter.Marshal, []byte(liLei), nil)
			stubs.StubFunc(&DestroyResource)
			//several So assert
		})

		Convey("for fail when num is too small", func() {
			stubs := Stub(&num, 50)
			defer stubs.Reset()
			//several So assert
		})

		Convey("for fail when Exec error", func() {
			stubs := Stub(&num, 150)
			defer stubs.Reset()
			stubs.StubFunc(&Exec, "", ErrAny)
			//several So assert
		})

		Convey("for fail when Marshal error", func() {
			stubs := Stub(&num, 150)
			defer stubs.Reset()
			stubs.StubFunc(&Exec,"xxx-vethName100-yyy", nil)
			stubs.StubFunc(&adapter.Marshal, nil, ErrAny)
			//several So assert
		})

	})
}
 */
