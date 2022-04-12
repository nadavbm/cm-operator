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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpConfigMapSpec defines the desired state of OpConfigMap
type OpConfigMapSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ConfigMapName is the configMap name
	ConfigMapName string `json:"cmname,omitempty"`
	// FileName is the file name in configmap data
	FileName string `json:"filename,omitempty"`
	// FileData is the data of the file
	FileData string `json:"filedata,omitempty"`
}

// OpConfigMapStatus defines the observed state of OpConfigMap
type OpConfigMapStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OpConfigMap is the Schema for the opconfigmaps API
type OpConfigMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpConfigMapSpec   `json:"spec,omitempty"`
	Status OpConfigMapStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OpConfigMapList contains a list of OpConfigMap
type OpConfigMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpConfigMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpConfigMap{}, &OpConfigMapList{})
}
