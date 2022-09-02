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

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// instances.start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type NodeGroupAutoscalingPolicyObservation struct {
}

type NodeGroupAutoscalingPolicyParameters struct {

	// Maximum size of the node group. Set to a value less than or equal
	// to 100 and greater than or equal to min-nodes.
	// +kubebuilder:validation:Optional
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// Minimum size of the node group. Must be less
	// than or equal to max-nodes. The default value is 0.
	// +kubebuilder:validation:Optional
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The autoscaling mode. Set to one of the following:
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type NodeGroupObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type NodeGroupParameters struct {

	// If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AutoscalingPolicy []NodeGroupAutoscalingPolicyParameters `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy,omitempty"`

	// An optional textual description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The initial number of nodes in the node group. One of initial_size or size must be specified.
	// +kubebuilder:validation:Optional
	InitialSize *float64 `json:"initialSize,omitempty" tf:"initial_size,omitempty"`

	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	// +kubebuilder:validation:Optional
	MaintenancePolicy *string `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// contains properties for the timeframe of maintenance
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The URL of the node template to which this node group belongs.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.NodeTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NodeTemplate *string `json:"nodeTemplate,omitempty" tf:"node_template,omitempty"`

	// Reference to a NodeTemplate in compute to populate nodeTemplate.
	// +kubebuilder:validation:Optional
	NodeTemplateRef *v1.Reference `json:"nodeTemplateRef,omitempty" tf:"-"`

	// Selector for a NodeTemplate in compute to populate nodeTemplate.
	// +kubebuilder:validation:Optional
	NodeTemplateSelector *v1.Selector `json:"nodeTemplateSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The total number of nodes in the node group. One of initial_size or size must be specified.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Zone where this node group is located
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroup is the Schema for the NodeGroups API. Represents a NodeGroup resource to manage a group of sole-tenant nodes.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroup_Kind             = "NodeGroup"
	NodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeGroup_Kind}.String()
	NodeGroup_KindAPIVersion   = NodeGroup_Kind + "." + CRDGroupVersion.String()
	NodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(NodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}
