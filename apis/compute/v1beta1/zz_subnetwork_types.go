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

type SecondaryIPRangeInitParameters struct {

	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range"`

	// The name associated with this subnetwork secondary range, used
	// when adding an alias IP range to a VM instance. The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	RangeName *string `json:"rangeName,omitempty" tf:"range_name"`
}

type SecondaryIPRangeObservation struct {

	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// The name associated with this subnetwork secondary range, used
	// when adding an alias IP range to a VM instance. The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	RangeName *string `json:"rangeName,omitempty" tf:"range_name,omitempty"`
}

type SecondaryIPRangeParameters struct {

	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range"`

	// The name associated with this subnetwork secondary range, used
	// when adding an alias IP range to a VM instance. The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	// +kubebuilder:validation:Optional
	RangeName *string `json:"rangeName,omitempty" tf:"range_name"`
}

type SubnetworkInitParameters struct {

	// Typically packets destined to IPs within the subnetwork range that do not match
	// existing resources are dropped and prevented from leaving the VPC.
	// Setting this field to true will allow these packets to match dynamic routes injected
	// via BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap *bool `json:"allowSubnetCidrRoutesOverlap,omitempty" tf:"allow_subnet_cidr_routes_overlap,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIPv6Prefix *string `json:"externalIpv6Prefix,omitempty" tf:"external_ipv6_prefix,omitempty"`

	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path.
	// Possible values are: EXTERNAL, INTERNAL.
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// This field denotes the VPC flow logging options for this subnetwork. If
	// logging is enabled, logs are exported to Cloud Logging. Flow logging
	// isn't supported if the subnet purpose field is set to subnetwork is
	// REGIONAL_MANAGED_PROXY or GLOBAL_MANAGED_PROXY.
	// Structure is documented below.
	LogConfig []SubnetworkLogConfigInitParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIPGoogleAccess *bool `json:"privateIpGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty"`

	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIPv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty" tf:"private_ipv6_google_access,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. This field can be either PRIVATE_RFC_1918, REGIONAL_MANAGED_PROXY, GLOBAL_MANAGED_PROXY, PRIVATE_SERVICE_CONNECT or PRIVATE_NAT(Beta).
	// A subnet with purpose set to REGIONAL_MANAGED_PROXY is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	// A subnetwork in a given region with purpose set to GLOBAL_MANAGED_PROXY is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	// A subnetwork with purpose set to PRIVATE_SERVICE_CONNECT reserves the subnet for hosting a Private Service Connect published service.
	// A subnetwork with purpose set to PRIVATE_NAT is used as source range for Private NAT gateways.
	// Note that REGIONAL_MANAGED_PROXY is the preferred setting for all regional Envoy load balancers.
	// If unspecified, the purpose defaults to PRIVATE_RFC_1918.
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The role of subnetwork.
	// Currently, this field is only used when purpose is REGIONAL_MANAGED_PROXY.
	// The value can be set to ACTIVE or BACKUP.
	// An ACTIVE subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	// A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are: ACTIVE, BACKUP.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges.
	// Note: This field uses attr-as-block mode to avoid
	// breaking users during the 0.12 upgrade. To explicitly send a list
	// of zero objects you must use the following syntax:
	// example=[]
	// For more details about this behavior, see this section.
	// Structure is documented below.
	SecondaryIPRange []SecondaryIPRangeInitParameters `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	// If not specified IPV4_ONLY will be used.
	// Possible values are: IPV4_ONLY, IPV4_IPV6.
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`
}

type SubnetworkLogConfigInitParameters struct {

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Toggles the aggregation interval for collecting flow logs. Increasing the
	// interval time will reduce the amount of generated flow logs for long
	// lasting connections. Default is an interval of 5 seconds per connection.
	// Default value is INTERVAL_5_SEC.
	// Possible values are: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
	AggregationInterval *string `json:"aggregationInterval,omitempty" tf:"aggregation_interval,omitempty"`

	// Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
	// https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.
	// The default value is 'true', which evaluates to include everything.
	FilterExpr *string `json:"filterExpr,omitempty" tf:"filter_expr,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// The value of the field must be in [0, 1]. Set the sampling rate of VPC
	// flow logs within the subnetwork where 1.0 means all collected logs are
	// reported and 0.0 means no logs are reported. Default is 0.5 which means
	// half of all collected logs are reported.
	FlowSampling *float64 `json:"flowSampling,omitempty" tf:"flow_sampling,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Configures whether metadata fields should be added to the reported VPC
	// flow logs.
	// Default value is INCLUDE_ALL_METADATA.
	// Possible values are: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA, CUSTOM_METADATA.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// List of metadata fields that should be added to reported logs.
	// Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.
	MetadataFields []*string `json:"metadataFields,omitempty" tf:"metadata_fields,omitempty"`
}

type SubnetworkLogConfigObservation struct {

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Toggles the aggregation interval for collecting flow logs. Increasing the
	// interval time will reduce the amount of generated flow logs for long
	// lasting connections. Default is an interval of 5 seconds per connection.
	// Default value is INTERVAL_5_SEC.
	// Possible values are: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
	AggregationInterval *string `json:"aggregationInterval,omitempty" tf:"aggregation_interval,omitempty"`

	// Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
	// https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.
	// The default value is 'true', which evaluates to include everything.
	FilterExpr *string `json:"filterExpr,omitempty" tf:"filter_expr,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// The value of the field must be in [0, 1]. Set the sampling rate of VPC
	// flow logs within the subnetwork where 1.0 means all collected logs are
	// reported and 0.0 means no logs are reported. Default is 0.5 which means
	// half of all collected logs are reported.
	FlowSampling *float64 `json:"flowSampling,omitempty" tf:"flow_sampling,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Configures whether metadata fields should be added to the reported VPC
	// flow logs.
	// Default value is INCLUDE_ALL_METADATA.
	// Possible values are: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA, CUSTOM_METADATA.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// List of metadata fields that should be added to reported logs.
	// Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.
	MetadataFields []*string `json:"metadataFields,omitempty" tf:"metadata_fields,omitempty"`
}

type SubnetworkLogConfigParameters struct {

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Toggles the aggregation interval for collecting flow logs. Increasing the
	// interval time will reduce the amount of generated flow logs for long
	// lasting connections. Default is an interval of 5 seconds per connection.
	// Default value is INTERVAL_5_SEC.
	// Possible values are: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
	// +kubebuilder:validation:Optional
	AggregationInterval *string `json:"aggregationInterval,omitempty" tf:"aggregation_interval,omitempty"`

	// Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
	// https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.
	// The default value is 'true', which evaluates to include everything.
	// +kubebuilder:validation:Optional
	FilterExpr *string `json:"filterExpr,omitempty" tf:"filter_expr,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// The value of the field must be in [0, 1]. Set the sampling rate of VPC
	// flow logs within the subnetwork where 1.0 means all collected logs are
	// reported and 0.0 means no logs are reported. Default is 0.5 which means
	// half of all collected logs are reported.
	// +kubebuilder:validation:Optional
	FlowSampling *float64 `json:"flowSampling,omitempty" tf:"flow_sampling,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Configures whether metadata fields should be added to the reported VPC
	// flow logs.
	// Default value is INCLUDE_ALL_METADATA.
	// Possible values are: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA, CUSTOM_METADATA.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// List of metadata fields that should be added to reported logs.
	// Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.
	// +kubebuilder:validation:Optional
	MetadataFields []*string `json:"metadataFields,omitempty" tf:"metadata_fields,omitempty"`
}

type SubnetworkObservation struct {

	// Typically packets destined to IPs within the subnetwork range that do not match
	// existing resources are dropped and prevented from leaving the VPC.
	// Setting this field to true will allow these packets to match dynamic routes injected
	// via BGP even if their destinations match existing subnet ranges.
	AllowSubnetCidrRoutesOverlap *bool `json:"allowSubnetCidrRoutesOverlap,omitempty" tf:"allow_subnet_cidr_routes_overlap,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The range of external IPv6 addresses that are owned by this subnetwork.
	ExternalIPv6Prefix *string `json:"externalIpv6Prefix,omitempty" tf:"external_ipv6_prefix,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// The gateway address for default routes to reach destination addresses
	// outside this subnetwork.
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path.
	// Possible values are: EXTERNAL, INTERNAL.
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// The range of internal IPv6 addresses that are owned by this subnetwork.
	IPv6CidrRange *string `json:"ipv6CidrRange,omitempty" tf:"ipv6_cidr_range,omitempty"`

	// The internal IPv6 address range that is assigned to this subnetwork.
	InternalIPv6Prefix *string `json:"internalIpv6Prefix,omitempty" tf:"internal_ipv6_prefix,omitempty"`

	// This field denotes the VPC flow logging options for this subnetwork. If
	// logging is enabled, logs are exported to Cloud Logging. Flow logging
	// isn't supported if the subnet purpose field is set to subnetwork is
	// REGIONAL_MANAGED_PROXY or GLOBAL_MANAGED_PROXY.
	// Structure is documented below.
	LogConfig []SubnetworkLogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIPGoogleAccess *bool `json:"privateIpGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty"`

	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIPv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty" tf:"private_ipv6_google_access,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. This field can be either PRIVATE_RFC_1918, REGIONAL_MANAGED_PROXY, GLOBAL_MANAGED_PROXY, PRIVATE_SERVICE_CONNECT or PRIVATE_NAT(Beta).
	// A subnet with purpose set to REGIONAL_MANAGED_PROXY is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	// A subnetwork in a given region with purpose set to GLOBAL_MANAGED_PROXY is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	// A subnetwork with purpose set to PRIVATE_SERVICE_CONNECT reserves the subnet for hosting a Private Service Connect published service.
	// A subnetwork with purpose set to PRIVATE_NAT is used as source range for Private NAT gateways.
	// Note that REGIONAL_MANAGED_PROXY is the preferred setting for all regional Envoy load balancers.
	// If unspecified, the purpose defaults to PRIVATE_RFC_1918.
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The GCP region for this subnetwork.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The role of subnetwork.
	// Currently, this field is only used when purpose is REGIONAL_MANAGED_PROXY.
	// The value can be set to ACTIVE or BACKUP.
	// An ACTIVE subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	// A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are: ACTIVE, BACKUP.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges.
	// Note: This field uses attr-as-block mode to avoid
	// breaking users during the 0.12 upgrade. To explicitly send a list
	// of zero objects you must use the following syntax:
	// example=[]
	// For more details about this behavior, see this section.
	// Structure is documented below.
	SecondaryIPRange []SecondaryIPRangeObservation `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	// If not specified IPV4_ONLY will be used.
	// Possible values are: IPV4_ONLY, IPV4_IPV6.
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`
}

type SubnetworkParameters struct {

	// Typically packets destined to IPs within the subnetwork range that do not match
	// existing resources are dropped and prevented from leaving the VPC.
	// Setting this field to true will allow these packets to match dynamic routes injected
	// via BGP even if their destinations match existing subnet ranges.
	// +kubebuilder:validation:Optional
	AllowSubnetCidrRoutesOverlap *bool `json:"allowSubnetCidrRoutesOverlap,omitempty" tf:"allow_subnet_cidr_routes_overlap,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The range of external IPv6 addresses that are owned by this subnetwork.
	// +kubebuilder:validation:Optional
	ExternalIPv6Prefix *string `json:"externalIpv6Prefix,omitempty" tf:"external_ipv6_prefix,omitempty"`

	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path.
	// Possible values are: EXTERNAL, INTERNAL.
	// +kubebuilder:validation:Optional
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// This field denotes the VPC flow logging options for this subnetwork. If
	// logging is enabled, logs are exported to Cloud Logging. Flow logging
	// isn't supported if the subnet purpose field is set to subnetwork is
	// REGIONAL_MANAGED_PROXY or GLOBAL_MANAGED_PROXY.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig []SubnetworkLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	// +kubebuilder:validation:Optional
	PrivateIPGoogleAccess *bool `json:"privateIpGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty"`

	// The private IPv6 google access type for the VMs in this subnet.
	// +kubebuilder:validation:Optional
	PrivateIPv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty" tf:"private_ipv6_google_access,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. This field can be either PRIVATE_RFC_1918, REGIONAL_MANAGED_PROXY, GLOBAL_MANAGED_PROXY, PRIVATE_SERVICE_CONNECT or PRIVATE_NAT(Beta).
	// A subnet with purpose set to REGIONAL_MANAGED_PROXY is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	// A subnetwork in a given region with purpose set to GLOBAL_MANAGED_PROXY is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	// A subnetwork with purpose set to PRIVATE_SERVICE_CONNECT reserves the subnet for hosting a Private Service Connect published service.
	// A subnetwork with purpose set to PRIVATE_NAT is used as source range for Private NAT gateways.
	// Note that REGIONAL_MANAGED_PROXY is the preferred setting for all regional Envoy load balancers.
	// If unspecified, the purpose defaults to PRIVATE_RFC_1918.
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The GCP region for this subnetwork.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The role of subnetwork.
	// Currently, this field is only used when purpose is REGIONAL_MANAGED_PROXY.
	// The value can be set to ACTIVE or BACKUP.
	// An ACTIVE subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	// A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are: ACTIVE, BACKUP.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges.
	// Note: This field uses attr-as-block mode to avoid
	// breaking users during the 0.12 upgrade. To explicitly send a list
	// of zero objects you must use the following syntax:
	// example=[]
	// For more details about this behavior, see this section.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SecondaryIPRange []SecondaryIPRangeParameters `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	// If not specified IPV4_ONLY will be used.
	// Possible values are: IPV4_ONLY, IPV4_IPV6.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`
}

// SubnetworkSpec defines the desired state of Subnetwork
type SubnetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetworkParameters `json:"forProvider"`
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
	InitProvider SubnetworkInitParameters `json:"initProvider,omitempty"`
}

// SubnetworkStatus defines the observed state of Subnetwork.
type SubnetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnetwork is the Schema for the Subnetworks API. A VPC network is a virtual version of the traditional physical networks that exist within and between physical data centers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Subnetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipCidrRange) || has(self.initProvider.ipCidrRange)",message="ipCidrRange is a required parameter"
	Spec   SubnetworkSpec   `json:"spec"`
	Status SubnetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetworkList contains a list of Subnetworks
type SubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnetwork `json:"items"`
}

// Repository type metadata.
var (
	Subnetwork_Kind             = "Subnetwork"
	Subnetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnetwork_Kind}.String()
	Subnetwork_KindAPIVersion   = Subnetwork_Kind + "." + CRDGroupVersion.String()
	Subnetwork_GroupVersionKind = CRDGroupVersion.WithKind(Subnetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnetwork{}, &SubnetworkList{})
}
