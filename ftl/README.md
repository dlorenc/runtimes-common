### FTL Integration Tests
To run the FTL integration tests, run the following commands locally from the root directory:

```shell
python ftl/ftl_node_integration_tests_yaml.py | gcloud container builds submit --config /dev/fd/0 .
python ftl/ftl_php_integration_tests_yaml.py | gcloud container builds submit --config /dev/fd/0 .
gcloud container builds submit --config ftl/ftl_python_integration_tests.yaml .
```
