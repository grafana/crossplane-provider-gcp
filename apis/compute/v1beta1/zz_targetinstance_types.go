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

type TargetInstanceObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. projects/project/zones/zone/instances/instance) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is NO_NAT.
	// Possible values are NO_NAT.
	NATPolicy *string `json:"natPolicy,omitempty" tf:"nat_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// URL of the zone where the target instance resides.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TargetInstanceParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. projects/project/zones/zone/instances/instance) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is NO_NAT.
	// Possible values are NO_NAT.
	// +kubebuilder:validation:Optional
	NATPolicy *string `json:"natPolicy,omitempty" tf:"nat_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// URL of the zone where the target instance resides.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

// TargetInstanceSpec defines the desired state of TargetInstance
type TargetInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetInstanceParameters `json:"forProvider"`
}

// TargetInstanceStatus defines the observed state of TargetInstance.
type TargetInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetInstance is the Schema for the TargetInstances API. Represents a TargetInstance resource which defines an endpoint instance that terminates traffic of certain protocols.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TargetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetInstanceSpec   `json:"spec"`
	Status            TargetInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetInstanceList contains a list of TargetInstances
type TargetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetInstance `json:"items"`
}

// Repository type metadata.
var (
	TargetInstance_Kind             = "TargetInstance"
	TargetInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetInstance_Kind}.String()
	TargetInstance_KindAPIVersion   = TargetInstance_Kind + "." + CRDGroupVersion.String()
	TargetInstance_GroupVersionKind = CRDGroupVersion.WithKind(TargetInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetInstance{}, &TargetInstanceList{})
}