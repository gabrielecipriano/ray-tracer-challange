package ray_tracer_challange_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRayTracerChallange(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RayTracerChallange Suite")
}
