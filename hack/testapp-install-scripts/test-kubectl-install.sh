#!/usr/bin/env bash

set -e

echo "Testing kubectl installation of testapp..."

# Create a test namespace
kubectl create namespace testapp-kubectl-test || true

# Install testapp using kubectl
testctl create testapp-kubectl \
  -n testapp-kubectl-test \
  --create-namespace \
  --debug \
  --connect=false

# Wait for the testapp to be ready
./hack/wait-for-pod.sh -l app=testapp-kubectl -n testapp-kubectl-test

# Clean up
testctl delete testapp-kubectl -n testapp-kubectl-test --delete-namespace

echo "kubectl installation test passed!"
