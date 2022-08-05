# opa-slp-operator

Deploys Styra Local Plane

### development

Initial skaffold using operator-sdk.  

# initialize a new empty operator template
1. operator-sdk init --domain twdps --repo github.com/ThoughtWorks-DPS/opa-slp-operator

# generate new API and controller
1. operator-sdk create api --group opa --version v1alpha1 --kind StyraLocalPlane --resource=true --controller=true

# Define resouce and status Struct fields

Add fields to type StyraLocalPlane struct
```go
	SystemName  string `json:"systemName,omitempty"`  // (required) Sytra System name
	AccessToken string `json:"accessToken,omitempty"` // (required) Styra access Token
	Namespace   string `json:"namespace,omitempty"`   // (optional) Namespace where SLP should be deployed
	Tag         string `json:"tag,omitempty"`         // (optional) styra/styra-local-plane image tag to use
```
```bash
$ make generate
```

Add fields to type StyraLocalPlaneStatus struct
```go
	Ready corev1.ConditionStatus `json:"ready"`  // slp status       
	Status string `json:"status,omitempty"`      // slp pod status
```
```bash
$ make generate
```

Generate manifests
```bash
$ make manifests
```

# customize controller

# in the defaults folder

set the namespace for the operator to run = opa-system

## modify makefile

IMG
VERSION
IMAGE_TAG_BASE