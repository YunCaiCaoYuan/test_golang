package integ2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func test2() int {
	return 2
}

var _ = Describe("test2", func() {
	It("test2", func() {
		//Expect(test2()).To(Equal(2))
		ExpectWithOffset(2, test2()).To(Equal(3)) // 堆栈偏移
	})
})

