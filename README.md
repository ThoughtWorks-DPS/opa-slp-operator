# opa-slp-operator

Deploys Styra Local Plane configured to match the Styra system pattern where all opa sidecars in a namespace share the same styra System and it's associated policy configuration.  

example:  

Assume a development team at book-info company known as the `publications` team since they own the product domain of _publications_ within book-info SaaS. Further assume that the publications team's path to production is through the following environments:  

dev => qa => production  

and that these environments equate to the matching namespaces of:

publications-dev => publications-qa => publications-prod  

The publications team then has a single git repository named `publications-authorization-policies` in which their team specific rego policies are maintained. There may be a styra Stack or shared libraries containing additional policies that will be included in their policy bundle.

The team would create matching Systems in their organization's Styra tenant, tied to the associated branches in the publications-authorization-policies_ repo:  

Systems:
- publications-dev => dev branch
- publications-qa => qa branch
- publications-prod => prod branch

Typically, it is recommended that the policy release pipeline for the policies repo follow a pattern of:

The team develop their policies within the repo using trunk-based-development against the `main` (or `master`) branch. The repo pipeline will then merge all changes pushed to `main` into the dev branch.  

To trigger a release to production, the team tags the main branch with the appropriate semantic version (or other versioning scheme as desired). Tagging causes the pipeline to first merge with the QA branch, and then upon approval further merge with the prod branch.  

Each of these branches being pulled in the respective styra System by the system automation.  

> Within Styra, the publications-dev system can either immediately publish all such changes out to the slp's pulling from publications-dev or wait for human interaction in the DAS UI depending on the workflow desired by the team. They may wish to take advantage of Styra's playback features before committing changes to all or only some of the environments.  

## requirements or assumptions

* This operator expects that the necessary credentials will already exist in the namespace, published as a kubernetes secret in the following format:  

_note: always confirm the url endpoint for your styra tenant_  

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

* This operator does not manage the mechanism whereby the opa-sidecar is deployed nor the EnvoyFilter necessary to direct traffic to the sidecar before the associated api.  

A team will either need to include these in their deployment, or the cluster configuration process will need to include management of namespace annotations, envoyfilter deployment, and the existence of a mutatingwebhook and admission controller to enable automatic injection with each deployment.  

### development

Inited using operator-sdk.  

# initialize a new empty operator helm template
1. operator-sdk init --plugins helm --domain twdps.io --group opa --version v1alpha1 --kind SlpDeployment

## to use this helm-managed slp deploy operator

- depends upon a credentials secrets already being deployed in the namespace where you deploy a resource request
- the credentials secret must have the following format

