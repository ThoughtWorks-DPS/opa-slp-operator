# opa-slp-operator

The opa-sl-operator automates the configuration and deployment of the Styra Local Plane within individual kubernetes namespaces.  


## Usage

**requirements**

The operator expects that the necessary DAS System opa token will already exist in the namespace, published as a kubernetes secret. The template below uses the default naming conventions. You may further customize these in the customr resource request.  

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: {{ namespace }}-slp-credentials
  namespace: {{ namespace }}
  labels:
    app: {{ namespace }}-slp
    system-type: istio
    managed-by: opa-slp-operator
type: Opaque
stringData:
  das.yaml: |
    discovery:
      name: discovery
      prefix: /systems/{{ the das systemId associated with the namespace }}
      service: styra
    labels:
      system-id: {{ the das systemId associated with the namespace }}
      system-type: template.istio:1.0
    services:
    - credentials:
        bearer:
          token: {{ opa system access token }}
      name: styra
      url: https://{{ styra tenant }}.styra.com/v1
    - credentials:
        bearer:
          token: {{ opa system access token }}
      name: styra-bundles
      url: https://{{ styra tenant }}.styra.com/v1/bundles
```

### Deploy Operator

Use the `resources.yaml` file to create the opa-system namespace and deploy the operator.    

```bash
curl https://raw.githubusercontent.com/ThoughtWorks-DPS/opa-slp-operator/main/resources.yaml | kubectl apply -f -
```
Wait a few seconds and confirm the operator is running.  
```bash
$ kubectl get all -n opa-system
NAME                                                       READY   STATUS    RESTARTS        AGE
pod/opa-slp-operator-controller-manager-56474d7c8b-zdt6q   2/2     Running   0               13s

NAME                                                          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/opa-slp-operator-controller-manager-metrics-service   ClusterIP   10.100.65.100    <none>        8443/TCP   2d

NAME                                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/opa-slp-operator-controller-manager   1/1     1            1           2d

NAME                                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/opa-slp-operator-controller-manager-56474d7c8b   1         1         1       13s
```

Once deployed the opa-slp-operator will watch all namespaces for SlpDeployment custom resource requests. When a custom resource request is deployed the operator will deploy and instance of the Styra Local Plane to the requesting namespace. Open policy agents running in the namespace may now be configured to featch their policy bundles directly from the local slp service.  

### Apply SLP resource request

```bash
cat <<EOF | kubetl apply -f -
apiVersion: opa.twdps.io/v1alpha1
kind: SlpDeployment
metadata:
  name: slpdeployment-example
spec:
  tenant: my-styra-das-tenant
  systemId: my-das-system-id
  namespace: my-namespace-name
EOF
```

Check to see that the slp service is now running in your namespace.  
```bash
$ kubectl get all -n my-namespace
NAME                           READY   STATUS    RESTARTS   AGE
pod/my-namespace-slp-0         2/2     Running   0          2d1h

NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/my-namespace-slp    ClusterIP   10.106.169.56   <none>        8080/TCP   2d1h

NAME                                 READY   AGE
statefulset.apps/my-namespace-slp    1/1     2d1h
```

The slp is now ready to serve the policy bundles from the specified Styra System.  

Slp is deployed as a stateful set and will define a persistent volume claim for a scratch area to log decision results. Refer to the Styra [documentation](https://docs.styra.com/policies/policy-organization/systems/use-styra-local-plane) for additional details about the local plane service.  

Read [Custom SLP Resource parameters](doc/resource_request_parameters.yaml) for additional configuration parameters.  

**Please Note** This operator does not manage the mechanism whereby the opa-sidecar is deployed nor the EnvoyFilter necessary to direct traffic to the sidecar before the associated api. A team will either need to include opa in their deployment definition, or the cluster configuration will need to include management of namespace annotations, envoy filter, and a mutatingwebhook and admission controller to enable automatic injection with each deployment.  

#### Helm deployment alternative

_pending_  

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to Helm's [documentation](https://helm.sh/docs) to get started.  

Once Helm has been set up correctly, add the twdps repo as follows:  

```bash
helm repo add twdps https://ThoughtWorks-DPS.github.io/helm-charts && helm repo update
```

To deploy the opa-slp-operator chart:  
```bash
helm install opa-slp-operator twdps/opa-slp-operator  
```

See chart documentation for configurable parameters.  

This operator was built using the [operator-sdk](https://sdk.operatorframework.io)  
