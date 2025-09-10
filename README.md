# Test Release OSS

A dummy test application for testing GitHub workflows and CI/CD pipelines. This project is designed to replicate the workflow structure of [vcluster](https://github.com/loft-sh/vcluster) for testing purposes.

## Overview

This repository contains:
- **testapp**: A simple HTTP server application
- **testctl**: A CLI tool for managing test resources
- Complete GitHub Actions workflows for CI/CD
- Helm chart for Kubernetes deployment
- Comprehensive testing infrastructure

## Components

### Applications

1. **testapp** - Main server application
   - HTTP server on port 8080
   - Health check endpoint at `/health`
   - Version endpoint at `/version`

2. **testctl** - Command line interface
   - Create, delete, and list test resources
   - Version command for build information

### Development

```bash
# Build applications
go build cmd/testapp/main.go
go build cmd/testctl/main.go

# Run tests
./hack/test.sh

# Run locally
go run cmd/testapp/main.go start
```

### Using Just

```bash
# Build CLI tool
just build-cli-snapshot

# Build server image
just build-snapshot

# Run tests
just test

# Run linting
just lint

# Run Helm tests
just helm-test
```

### Docker

```bash
# Build image
docker build -t test-release-oss:dev .

# Run container
docker run --rm -p 8080:8080 test-release-oss:dev
```

### Helm

```bash
# Install chart
helm install my-test chart/

# Run chart tests
helm unittest chart/
```

## GitHub Workflows

This repository includes comprehensive GitHub Actions workflows:

- **Unit Tests**: Go and Helm testing
- **Linting**: Code quality checks
- **Release**: Automated releases with goreleaser
- **E2E Testing**: End-to-end testing (adapted from vcluster)
- **Security**: Vulnerability scanning and license checks

## Purpose

This is a test repository designed to validate GitHub Actions workflows without affecting production systems. It provides a minimal but complete application structure that can be built, tested, and deployed using the same patterns as more complex projects.

## License

Apache License 2.0 - see [LICENSE](LICENSE) file for details.
