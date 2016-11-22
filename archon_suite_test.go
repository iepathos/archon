package archon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArchon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Archon Suite")
}
