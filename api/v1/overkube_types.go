/*
Copyright 2020 Nicolas Kowenski.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type workers struct {
	// Size is the size of the memcached deployment
	Size int32 `json:"workers"`
}
type controlPlaneEndpoint struct {
	// Nodes are the names of the memcached pods
	Nodes string `json:"controlPlaneEndpoint"`
}

// OverkubeSpec defines the desired state of Overkube
type OverkubeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Overkube. Edit Overkube_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// OverkubeStatus defines the observed state of Overkube
type OverkubeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Overkube is the Schema for the overkubes API
type Overkube struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OverkubeSpec   `json:"spec,omitempty"`
	Status OverkubeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OverkubeList contains a list of Overkube
type OverkubeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Overkube `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Overkube{}, &OverkubeList{})
}
