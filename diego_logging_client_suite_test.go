package diego_logging_client_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDiegoLoggingClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DiegoLoggingClient Suite")
}
