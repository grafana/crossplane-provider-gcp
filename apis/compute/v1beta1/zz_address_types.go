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

<<<<<<< HEAD
type AddressInitParameters struct {

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any. Set by the API if undefined.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of address to reserve.
	// Note: if you set this argument's value as INTERNAL you need to leave the network_tier argument unset in that resource block.
	// Default value is EXTERNAL.
	// Possible values are: INTERNAL, EXTERNAL.
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	// This argument should not be used when configuring Internal addresses, because network tier cannot be set for internal traffic; it's always Premium.
	// Possible values are: PREMIUM, STANDARD.
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// The prefix length if the resource represents an IP range.
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of this resource, which can be one of the following values.
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`
}

=======
>>>>>>> 5366d83e (Enable the following resources:)
type AddressObservation struct {

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any. Set by the API if undefined.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of address to reserve.
	// Note: if you set this argument's value as INTERNAL you need to leave the network_tier argument unset in that resource block.
	// Default value is EXTERNAL.
	// Possible values are: INTERNAL, EXTERNAL.
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/addresses/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL of the network in which to reserve the address. This field
	// can only be used with INTERNAL type with the VPC_PEERING and
	// IPSEC_INTERCONNECT purposes.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	// This argument should not be used when configuring Internal addresses, because network tier cannot be set for internal traffic; it's always Premium.
	// Possible values are: PREMIUM, STANDARD.
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// The prefix length if the resource represents an IP range.
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of this resource, which can be one of the following values.
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The URLs of the resources that are using this address.
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type AddressParameters struct {

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any. Set by the API if undefined.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of address to reserve.
	// Note: if you set this argument's value as INTERNAL you need to leave the network_tier argument unset in that resource block.
	// Default value is EXTERNAL.
	// Possible values are: INTERNAL, EXTERNAL.
	// +kubebuilder:validation:Optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The URL of the network in which to reserve the address. This field
	// can only be used with INTERNAL type with the VPC_PEERING and
	// IPSEC_INTERCONNECT purposes.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	// This argument should not be used when configuring Internal addresses, because network tier cannot be set for internal traffic; it's always Premium.
	// Possible values are: PREMIUM, STANDARD.
	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// The prefix length if the resource represents an IP range.
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of this resource, which can be one of the following values.
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	// +crossplane:generate:reference:type=Subnetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

// AddressSpec defines the desired state of Address
type AddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressParameters `json:"forProvider"`
<<<<<<< HEAD
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AddressInitParameters `json:"initProvider,omitempty"`
=======
>>>>>>> 5366d83e (Enable the following resources:)
}

// AddressStatus defines the observed state of Address.
type AddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Address is the Schema for the Addresss API. Represents an Address resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddressSpec   `json:"spec"`
	Status            AddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressList contains a list of Addresss
type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Address `json:"items"`
}

// Repository type metadata.
var (
	Address_Kind             = "Address"
	Address_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Address_Kind}.String()
	Address_KindAPIVersion   = Address_Kind + "." + CRDGroupVersion.String()
	Address_GroupVersionKind = CRDGroupVersion.WithKind(Address_Kind)
)

func init() {
	SchemeBuilder.Register(&Address{}, &AddressList{})
}
