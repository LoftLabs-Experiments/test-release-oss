package main

import (
	"os"

	"github.com/LoftLabs-Experiments/test-release-oss/cmd/testctl/cmd"
	"github.com/LoftLabs-Experiments/test-release-oss/pkg/telemetry"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var version = "0.0.1"

func main() {
	telemetry.SetVersion(version)

	cmd.Execute()
	os.Exit(0)
}
