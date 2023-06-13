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

type NetworkObservation struct {

	// When set to true, the network is created in "auto subnet mode" and
	// it will create a subnet for each region automatically across the
	// 10.128.0.0/9 address range.
	// When set to false, the network is created in "custom subnet mode" so
	// the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks *bool `json:"autoCreateSubnetworks,omitempty" tf:"auto_create_subnetworks,omitempty"`

	// If set to true, default routes (0.0.0.0/0) will be deleted
	// immediately after network creation. Defaults to false.
	DeleteDefaultRoutesOnCreate *bool `json:"deleteDefaultRoutesOnCreate,omitempty" tf:"delete_default_routes_on_create,omitempty"`

	// An optional description of this resource. The resource must be
	// recreated to modify this field.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Enable ULA internal ipv6 on this network. Enabling this feature will assign
	// a /48 from google defined ULA prefix fd20::/20.
	EnableUlaInternalIPv6 *bool `json:"enableUlaInternalIpv6,omitempty" tf:"enable_ula_internal_ipv6,omitempty"`

	// The gateway address for default routing out of the network. This value
	// is selected by GCP.
	GatewayIPv4 *string `json:"gatewayIpv4,omitempty" tf:"gateway_ipv4,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/networks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When enabling ula internal ipv6, caller optionally can specify the /48 range
	// they want from the google defined ULA prefix fd20::/20. The input must be a
	// valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will
	// fail if the speficied /48 is already in used by another resource.
	// If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field.
	InternalIPv6Range *string `json:"internalIpv6Range,omitempty" tf:"internal_ipv6_range,omitempty"`

	// Maximum Transmission Unit in bytes. The default value is 1460 bytes.
	// The minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).
	// Note that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped
	// with an ICMP Fragmentation-Needed message if the packets are routed to the Internet or other VPCs
	// with varying MTUs.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// Set the order that Firewall Rules and Firewall Policies are evaluated.
	// Default value is AFTER_CLASSIC_FIREWALL.
	// Possible values are: BEFORE_CLASSIC_FIREWALL, AFTER_CLASSIC_FIREWALL.
	NetworkFirewallPolicyEnforcementOrder *string `json:"networkFirewallPolicyEnforcementOrder,omitempty" tf:"network_firewall_policy_enforcement_order,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The network-wide routing mode to use. If set to REGIONAL, this
	// network's cloud routers will only advertise routes with subnetworks
	// of this network in the same region as the router. If set to GLOBAL,
	// this network's cloud routers will advertise routes with all
	// subnetworks of this network, across regions.
	// Possible values are: REGIONAL, GLOBAL.
	RoutingMode *string `json:"routingMode,omitempty" tf:"routing_mode,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type NetworkParameters struct {

	// When set to true, the network is created in "auto subnet mode" and
	// it will create a subnet for each region automatically across the
	// 10.128.0.0/9 address range.
	// When set to false, the network is created in "custom subnet mode" so
	// the user can explicitly connect subnetwork resources.
	// +kubebuilder:validation:Optional
	AutoCreateSubnetworks *bool `json:"autoCreateSubnetworks,omitempty" tf:"auto_create_subnetworks,omitempty"`

	// If set to true, default routes (0.0.0.0/0) will be deleted
	// immediately after network creation. Defaults to false.
	// +kubebuilder:validation:Optional
	DeleteDefaultRoutesOnCreate *bool `json:"deleteDefaultRoutesOnCreate,omitempty" tf:"delete_default_routes_on_create,omitempty"`

	// An optional description of this resource. The resource must be
	// recreated to modify this field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Enable ULA internal ipv6 on this network. Enabling this feature will assign
	// a /48 from google defined ULA prefix fd20::/20.
	// +kubebuilder:validation:Optional
	EnableUlaInternalIPv6 *bool `json:"enableUlaInternalIpv6,omitempty" tf:"enable_ula_internal_ipv6,omitempty"`

	// When enabling ula internal ipv6, caller optionally can specify the /48 range
	// they want from the google defined ULA prefix fd20::/20. The input must be a
	// valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will
	// fail if the speficied /48 is already in used by another resource.
	// If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field.
	// +kubebuilder:validation:Optional
	InternalIPv6Range *string `json:"internalIpv6Range,omitempty" tf:"internal_ipv6_range,omitempty"`

	// Maximum Transmission Unit in bytes. The default value is 1460 bytes.
	// The minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).
	// Note that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped
	// with an ICMP Fragmentation-Needed message if the packets are routed to the Internet or other VPCs
	// with varying MTUs.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// Set the order that Firewall Rules and Firewall Policies are evaluated.
	// Default value is AFTER_CLASSIC_FIREWALL.
	// Possible values are: BEFORE_CLASSIC_FIREWALL, AFTER_CLASSIC_FIREWALL.
	// +kubebuilder:validation:Optional
	NetworkFirewallPolicyEnforcementOrder *string `json:"networkFirewallPolicyEnforcementOrder,omitempty" tf:"network_firewall_policy_enforcement_order,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The network-wide routing mode to use. If set to REGIONAL, this
	// network's cloud routers will only advertise routes with subnetworks
	// of this network in the same region as the router. If set to GLOBAL,
	// this network's cloud routers will advertise routes with all
	// subnetworks of this network, across regions.
	// Possible values are: REGIONAL, GLOBAL.
	// +kubebuilder:validation:Optional
	RoutingMode *string `json:"routingMode,omitempty" tf:"routing_mode,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Network is the Schema for the Networks API. Manages a VPC network or legacy network resource on GCP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSpec   `json:"spec"`
	Status            NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
