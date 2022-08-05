/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StyraLocalPlaneSpec defines the desired state of StyraLocalPlane
type StyraLocalPlaneSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Tenant      string `json:"tenant,omitempty"`      // Styra Tenant
	SystemName  string `json:"systemName,omitempty"`  // Sytra System name
	AccessToken string `json:"accessToken,omitempty"` // Styra access Token
	Namespace   string `json:"namespace,omitempty"`   // Namespace where SLP should be deployed
	Tag         string `json:"tag,omitempty"`         // styra/styra-local-plane image tag to use
}

// StyraLocalPlaneStatus defines the observed state of StyraLocalPlane
type StyraLocalPlaneStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Ready  corev1.ConditionStatus `json:"ready"`            // slp status
	Status string                 `json:"status,omitempty"` // slp pod status
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StyraLocalPlane is the Schema for the styralocalplanes API
type StyraLocalPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StyraLocalPlaneSpec   `json:"spec,omitempty"`
	Status StyraLocalPlaneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StyraLocalPlaneList contains a list of StyraLocalPlane
type StyraLocalPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StyraLocalPlane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StyraLocalPlane{}, &StyraLocalPlaneList{})
}
