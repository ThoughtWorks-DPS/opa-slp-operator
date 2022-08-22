#!/usr/bin/env bats

@test "operator liveness" {
  run bash -c "kubectl get all -n opa-system"
  [[ "${output}" =~ "2/2" ]]
}

@test "operator readiness" {
  run bash -c "kubectl logs -n opa-system $(kubectl get pods -o jsonpath={..metadata.name} -n opa-system)"
  [[ "${output}" =~ "Starting workers" ]]
}
