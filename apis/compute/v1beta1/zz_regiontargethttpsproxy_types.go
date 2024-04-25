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

type RegionTargetHTTPSProxyInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type RegionTargetHTTPSProxyObservation struct {

	// URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, you may specify up to 15 certificates. Certificate manager certificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	// sslCertificates and certificateManagerCertificates fields can not be defined together.
	// Accepted format is //certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName} or just the self_link projects/{project}/locations/{location}/certificates/{resourceName}
	CertificateManagerCertificates []*string `json:"certificateManagerCertificates,omitempty" tf:"certificate_manager_certificates,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The unique identifier for the resource.
	ProxyID *float64 `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// The Region in which the created target https proxy should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer.
	// At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
	// sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	SSLCertificates []*string `json:"sslCertificates,omitempty" tf:"ssl_certificates,omitempty"`

	// A reference to the Region SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A reference to the RegionUrlMap resource that defines the mapping from URL
	// to the RegionBackendService.
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`
}

type RegionTargetHTTPSProxyParameters struct {

	// URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, you may specify up to 15 certificates. Certificate manager certificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	// sslCertificates and certificateManagerCertificates fields can not be defined together.
	// Accepted format is //certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName} or just the self_link projects/{project}/locations/{location}/certificates/{resourceName}
	// +kubebuilder:validation:Optional
	CertificateManagerCertificates []*string `json:"certificateManagerCertificates,omitempty" tf:"certificate_manager_certificates,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created target https proxy should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer.
	// At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
	// sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	// +crossplane:generate:reference:type=RegionSSLCertificate
	// +kubebuilder:validation:Optional
	SSLCertificates []*string `json:"sslCertificates,omitempty" tf:"ssl_certificates,omitempty"`

	// References to RegionSSLCertificate to populate sslCertificates.
	// +kubebuilder:validation:Optional
	SSLCertificatesRefs []v1.Reference `json:"sslCertificatesRefs,omitempty" tf:"-"`

	// Selector for a list of RegionSSLCertificate to populate sslCertificates.
	// +kubebuilder:validation:Optional
	SSLCertificatesSelector *v1.Selector `json:"sslCertificatesSelector,omitempty" tf:"-"`

	// A reference to the Region SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	// +kubebuilder:validation:Optional
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// A reference to the RegionUrlMap resource that defines the mapping from URL
	// to the RegionBackendService.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.RegionURLMap
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`

	// Reference to a RegionURLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapRef *v1.Reference `json:"urlMapRef,omitempty" tf:"-"`

	// Selector for a RegionURLMap in compute to populate urlMap.
	// +kubebuilder:validation:Optional
	URLMapSelector *v1.Selector `json:"urlMapSelector,omitempty" tf:"-"`
}

// RegionTargetHTTPSProxySpec defines the desired state of RegionTargetHTTPSProxy
type RegionTargetHTTPSProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionTargetHTTPSProxyParameters `json:"forProvider"`
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
	InitProvider RegionTargetHTTPSProxyInitParameters `json:"initProvider,omitempty"`
}

// RegionTargetHTTPSProxyStatus defines the observed state of RegionTargetHTTPSProxy.
type RegionTargetHTTPSProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionTargetHTTPSProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionTargetHTTPSProxy is the Schema for the RegionTargetHTTPSProxys API. Represents a RegionTargetHttpsProxy resource, which is used by one or more forwarding rules to route incoming HTTPS requests to a URL map.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionTargetHTTPSProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionTargetHTTPSProxySpec   `json:"spec"`
	Status            RegionTargetHTTPSProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionTargetHTTPSProxyList contains a list of RegionTargetHTTPSProxys
type RegionTargetHTTPSProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionTargetHTTPSProxy `json:"items"`
}

// Repository type metadata.
var (
	RegionTargetHTTPSProxy_Kind             = "RegionTargetHTTPSProxy"
	RegionTargetHTTPSProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionTargetHTTPSProxy_Kind}.String()
	RegionTargetHTTPSProxy_KindAPIVersion   = RegionTargetHTTPSProxy_Kind + "." + CRDGroupVersion.String()
	RegionTargetHTTPSProxy_GroupVersionKind = CRDGroupVersion.WithKind(RegionTargetHTTPSProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionTargetHTTPSProxy{}, &RegionTargetHTTPSProxyList{})
}
