/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ConfigurationSetParameters defines the desired state of ConfigurationSet
type ConfigurationSetParameters struct {
	// Region is which region the ConfigurationSet will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A data structure that contains the name of the configuration set.
	// +kubebuilder:validation:Required
	ConfigurationSet                 *ConfigurationSet_SDK `json:"configurationSet"`
	CustomConfigurationSetParameters `json:",inline"`
}

// ConfigurationSetSpec defines the desired state of ConfigurationSet
type ConfigurationSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigurationSetParameters `json:"forProvider"`
}

// ConfigurationSetObservation defines the observed state of ConfigurationSet
type ConfigurationSetObservation struct {
}

// ConfigurationSetStatus defines the observed state of ConfigurationSet.
type ConfigurationSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigurationSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationSet is the Schema for the ConfigurationSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigurationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationSetSpec   `json:"spec"`
	Status            ConfigurationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationSetList contains a list of ConfigurationSets
type ConfigurationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationSet `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationSetKind             = "ConfigurationSet"
	ConfigurationSetGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigurationSetKind}.String()
	ConfigurationSetKindAPIVersion   = ConfigurationSetKind + "." + GroupVersion.String()
	ConfigurationSetGroupVersionKind = GroupVersion.WithKind(ConfigurationSetKind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationSet{}, &ConfigurationSetList{})
}
