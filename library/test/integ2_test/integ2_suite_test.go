package integ2


import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegTest Suite")
}
