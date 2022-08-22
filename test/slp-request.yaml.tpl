apiVersion: opa.twdps.io/v1alpha1
kind: SlpDeployment
metadata:
  name: ci-dev-slp
  namespace: ci-dev
spec:
  tenant: thoughtworks
  systemId: {{ op://empc-lab/svc-styra/ci-test-fixture-systemid }}
  namespace: ci-dev