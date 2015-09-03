package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGonway(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gonway Suite")
}
