#!/usr/bin/env bats

@test "slp liveness" {
  run bash -c "kubectl get all -n ci-dev"
  [[ "${output}" =~ "Running" ]]
}

@test "slp readiness" {
  run bash -c "kubectl logs pod/ci-dev-slp-0 -n ci-dev --tail=1"
  [[ "${output}" =~ "Start serving end-point" ]]
}
