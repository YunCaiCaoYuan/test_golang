package integ_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Substring matching", func() {
	type SubstringCase struct {
		String    string
		Substring string
		Count     int
	}

	DescribeTable("counting substring matches",
		func(c SubstringCase) {
			Ω(strings.Count(c.String, c.Substring)).Should(BeNumerically("==", c.Count))
		},
		Entry("with no matching substring", SubstringCase{
			String:    "the sixth sheikh's sixth sheep's sick",
			Substring: "emir",
			Count:     0,
		}),
		Entry("with one matching substring", SubstringCase{
			String:    "the sixth sheikh's sixth sheep's sick",
			Substring: "sheep",
			Count:     1,
		}),
		Entry("with many matching substring", SubstringCase{
			String:    "the sixth sheikh's sixth sheep's sick",
			Substring: "si",
			Count:     3,
		}),
	)
})
