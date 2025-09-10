#!/usr/bin/env bash

set -e

echo "Testing helm installation of testapp..."

# Create a test namespace
kubectl create namespace testapp-helm-test || true

# Install testapp using helm
helm install testapp-helm ./chart \
  -n testapp-helm-test \
  --create-namespace \
  --wait

# Wait for the testapp to be ready
./hack/wait-for-pod.sh -l app.kubernetes.io/name=test-release-oss -n testapp-helm-test

# Clean up
helm uninstall testapp-helm -n testapp-helm-test
kubectl delete namespace testapp-helm-test

echo "helm installation test passed!"
