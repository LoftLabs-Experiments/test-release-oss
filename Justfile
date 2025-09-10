set positional-arguments

timestamp := `date +%s`

GOOS := env("GOOS", `go env GOOS`)
GOARCH := env("GOARCH", `go env GOARCH`)
GOBIN := env("GOBIN", `go env GOPATH`+"/bin")

DIST_FOLDER := if GOARCH == "amd64" { "dist/testapp_linux_amd64_v1" } else if GOARCH == "arm64" { "dist/testapp_linux_arm64_v8.0" } else { "unknown" }
DIST_FOLDER_CLI := if GOARCH == "amd64" { "dist/testctl_" + GOOS + "_amd64_v1" } else if GOARCH == "arm64" { "dist/testctl_" + GOOS + "_arm64_v8.0" } else { "unknown" }

_default:
  @just --list

# --- Build ---

# Build the testctl binary
build-cli-snapshot:
  goreleaser build --id testctl --single-target --snapshot --clean
  mv {{DIST_FOLDER_CLI}}/testctl {{GOBIN}}/testctl

# Build the testapp binary (we force linux here to allow building on mac os or windows)
build-snapshot:
  GOOS=linux goreleaser build --id testapp --single-target --snapshot --clean
  cp Dockerfile.release {{DIST_FOLDER}}/Dockerfile
  cd {{DIST_FOLDER}} && docker buildx build --load . -t ghcr.io/loftlabs-experiments/test-release-oss:dev-next

# --- Test ---

# Run unit tests
test:
  go test -race -v ./...

# Run linting
lint:
  golangci-lint run

# --- Development ---

# Run the test application locally
dev:
  go run cmd/testapp/main.go start

# Build and run in docker
dev-docker:
  docker build -t test-release-oss:dev .
  docker run --rm -p 8080:8080 test-release-oss:dev

# Clean up build artifacts
clean:
  rm -rf dist/
  docker rmi -f ghcr.io/loftlabs-experiments/test-release-oss:dev-next || true

# Generate Go modules
mod-tidy:
  go mod tidy

# Format code
fmt:
  go fmt ./...

# --- Helm ---

# Run helm unit tests
helm-test:
  helm unittest chart/

# --- Release ---

# Embed chart for release builds
embed-chart version:
  echo "Embedding chart for version {{version}}"
  # This would embed the Helm chart into the binary for releases
  # In vcluster this generates embedded chart data

# Clean release artifacts
clean-release:
  echo "Cleaning release artifacts"
  rm -rf release/* || true
  mkdir -p release

# Copy release assets
copy-assets:
  echo "Copying release assets"
  mkdir -p release
  cp chart/values.schema.json release/ || true

# Generate testapp latest images
generate-testapp-latest-images version:
  echo "Generating testapp latest images for version {{version}}"
  # This would generate image variants for the release

# Generate testapp optional images  
generate-testapp-optional-images version:
  echo "Generating testapp optional images for version {{version}}"
  # This would generate additional image variants

# Generate matrix specific images
generate-matrix-specific-images version:
  echo "Generating matrix specific images for version {{version}}"
  # This would generate images for different Kubernetes versions