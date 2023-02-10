package go_version

// https://cloud.tencent.com/developer/article/2195168

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// FuzzEncodeJSON fuzzing test for `EncodeJSON`, case 1: one kind of input
func FuzzEncodeJSON(f *testing.F) {
	testCases := []string{"abc", "", "\\U000f59fcş", " "}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		convey.Convey("fuzz string", t, func() {
			_, err := EncodeJSON(input)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func EncodeJSON(input string) (string, error) {
	//return input, errors.New("error")

	return input, nil
}
