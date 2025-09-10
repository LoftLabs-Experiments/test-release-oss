package e2e_basic

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestBasic(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Basic E2E Suite")
}

var _ = ginkgo.Describe("Basic Tests", func() {
	ginkgo.It("should run basic functionality tests", func() {
		gomega.Expect(true).To(gomega.BeTrue())
	})
})
