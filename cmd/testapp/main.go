package main

import (
	"github.com/LoftLabs-Experiments/test-release-oss/cmd/testapp/cmd"

	// Make sure dep tools picks up these dependencies
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	cmd.RunRoot()
}
