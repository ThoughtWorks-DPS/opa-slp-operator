The SlpDeployment custom resource request supports the following parameters and overrides:   

```yaml
apiVersion: opa.twdps.io/v1alpha1
kind: SlpDeployment
metadata:
  name: slpdeployment-sample
spec:
  # required fields
  tenant: my-styra-das-tenant   # used to construct the url used by slp to fetch policies from Styra DAS
  systemId: my-das-system-id    # the system id for the DAS System where my policies are stored
  namespace: my-namespace-name  # expected to be the name of the namespace where the custom resource request is deployed

  # optional overrides

  # by default the operator will expect the secret filename to be <namespace>-<chart name>-credentials
  # and the secret key name to be das.yaml
  # you may overide this settings
  systemCredentialsFilenameOverride: "" 
  systemCredentialsSecretNameOverride: ""

  # may override the cluster default storage class,
  # specify the storage class name for the persistent volume template
  storageClassName: ""

  # default service name = my-namespace-slp
  nameOverride: ""

  image:
    repository: styra/styra-local-plane
    pullPolicy: Always
    # Overrides the image tag whose default is the chart appVersion.
    tag: ""
  imagePullSecrets: []

  service:
    type: ClusterIP
    port: 8080

  serviceAccount:
    create: true
    annotations: {}

  podAnnotations: {}

  resources: {}
    # To set container resources, uncomment the following
    # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
    # limits:
    #   cpu: 100m
    #   memory: 128Mi
    # requests:
    #   cpu: 100m
    #   memory: 128Mi

  nodeSelector: {}

  tolerations: []

  affinity: {}
```