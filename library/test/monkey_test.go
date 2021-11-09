package test

import (
	. "bou.ke/monkey"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func IsEqual(a, b string) bool {
	fmt.Println("not monkey")
	return a == b
}

func TestIsEqual(t *testing.T) {
	Convey("test is equal", t, func() {
		Convey("for patch true", func() {
			guard := Patch(IsEqual, func(_, _ string) bool {
				fmt.Println("monkey...")
				return true
			})
			defer guard.Unpatch()
			ok := IsEqual("hello", "world")
			So(ok, ShouldBeTrue)
		})
	})
}

/*
func TestExec(t *testing.T) {
	Convey("test has digit", t, func() {
		Convey("for succ", func() {
			outputExpect := "xxx-vethName100-yyy"
			guard := Patch(osencap.Exec, func(_ string, _ ...string) (string, error) {
				return outputExpect, nil
			})
			defer guard.Unpatch()
			output, err := osencap.Exec(any, any)
			So(output, ShouldEqual, outputExpect)
			So(err, ShouldBeNil)
		})
	})
}*/
