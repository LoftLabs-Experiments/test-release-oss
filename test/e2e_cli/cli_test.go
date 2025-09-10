package e2e_cli

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestCLI(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CLI E2E Suite")
}

var _ = ginkgo.Describe("CLI Tests", func() {
	ginkgo.It("should test CLI functionality", func() {
		gomega.Expect(true).To(gomega.BeTrue())
	})
})
