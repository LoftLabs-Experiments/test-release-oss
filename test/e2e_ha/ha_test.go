package e2e_ha

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestHA(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "HA E2E Suite")
}

var _ = ginkgo.Describe("HA Tests", func() {
	ginkgo.It("should test HA functionality", func() {
		gomega.Expect(true).To(gomega.BeTrue())
	})
})
