package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChattertonMessengerServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChattertonMessengerServer Suite")
}

var _ = Describe("MessageActions", func() {})
