### How to execute the e2e tests locally

1. Start test environment:
    - Deploy testapp using helm:
      ```bash
      helm install testapp ./chart --namespace testapp --create-namespace
      ```
    
    - Or run testapp locally:
      ```bash
      go run cmd/testapp/main.go start
      ```
          
2. Run the e2e tests:
    ```bash
    cd test/<test_suite_path>
    TESTAPP_NAMESPACE=testapp go test -v -ginkgo.v
    ```

