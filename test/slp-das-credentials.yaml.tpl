apiVersion: v1
kind: Secret
metadata:
  name: ci-dev-slp-credentials
  namespace: ci-dev
  labels:
    app: ci-dev-slp
    version: 0.4.4
    system-type: istio
type: Opaque
stringData:
  das.yaml: |
    discovery:
      name: discovery
      prefix: /systems/{{ op://empc-lab/svc-styra/ci-test-fixture-systemid }}
      service: styra
    labels:
      system-id: {{ op://empc-lab/svc-styra/ci-test-fixture-systemid }}
      system-type: template.istio:1.0
    services:
    - credentials:
        bearer:
          token: {{ op://empc-lab/svc-styra/ci-test-fixture-opatoken }}
      name: styra
      url: https://thoughtworks.styra.com/v1
    - credentials:
        bearer:
          token: {{ op://empc-lab/svc-styra/ci-test-fixture-opatoken }}
      name: styra-bundles
      url: https://thoughtworks.styra.com/v1/bundles
