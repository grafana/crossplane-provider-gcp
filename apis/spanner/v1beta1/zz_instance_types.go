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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceObservation struct {

	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form regional-europe-west1 , us-central etc.
	// In order to obtain a valid list please consult the
	// Configuration section of the docs.
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// When deleting a spanner instance, this boolean option will delete all backups of this instance.
	// This must be set to true if you created a backup manually in the console.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// an identifier for the resource with format {{project}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The number of nodes allocated to this instance.
	NumNodes *float64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// The number of processing units allocated to this instance.
	ProcessingUnits *float64 `json:"processingUnits,omitempty" tf:"processing_units,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Instance status: CREATING or READY.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type InstanceParameters struct {

	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form regional-europe-west1 , us-central etc.
	// In order to obtain a valid list please consult the
	// Configuration section of the docs.
	// +kubebuilder:validation:Optional
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// When deleting a spanner instance, this boolean option will delete all backups of this instance.
	// This must be set to true if you created a backup manually in the console.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The number of nodes allocated to this instance.
	// +kubebuilder:validation:Optional
	NumNodes *float64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// The number of processing units allocated to this instance.
	// +kubebuilder:validation:Optional
	ProcessingUnits *float64 `json:"processingUnits,omitempty" tf:"processing_units,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. An isolated set of Cloud Spanner resources on which databases can be hosted.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.config)",message="config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}