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

type TargetHTTPSProxyInitParameters struct {

	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is //certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}.
	CertificateMap *string `json:"certificateMap,omitempty" tf:"certificate_map,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `json:"proxyBind,omitempty" tf:"proxy_bind,omitempty"`

	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is NONE.
	// Possible values are: NONE, ENABLE, DISABLE.
	QuicOverride *string `json:"quicOverride,omitempty" tf:"quic_override,omitempty"`

	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
}

type TargetHTTPSProxyObservation struct {

	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is //certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}.
	CertificateMap *string `json:"certificateMap,omitempty" tf:"certificate_map,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/targetHttpsProxies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `json:"proxyBind,omitempty" tf:"proxy_bind,omitempty"`

	// The unique identifier for the resource.
	ProxyID *float64 `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is NONE.
	// Possible values are: NONE, ENABLE, DISABLE.
	QuicOverride *string `json:"quicOverride,omitempty" tf:"quic_override,omitempty"`

	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SSLCertificates []*string `json:"sslCertificates,omitempty" tf:"ssl_certificates,omitempty"`

	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`
}

type TargetHTTPSProxyParameters struct {

	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is //certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}.
	// +kubebuilder:validation:Optional
	CertificateMap *string `json:"certificateMap,omitempty" tf:"certificate_map,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// +kubebuilder:validation:Optional
	ProxyBind *bool `json:"proxyBind,omitempty" tf:"proxy_bind,omitempty"`

	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is NONE.
	// Possible values are: NONE, ENABLE, DISABLE.
	// +kubebuilder:validation:Optional
	QuicOverride *string `json:"quicOverride,omitempty" tf:"quic_override,omitempty"`

	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	// +crossplane:generate:reference:type=SSLCertificate
	// +kubebuilder:validation:Optional
	SSLCertificates []*string `json:"sslCertificates,omitempty" tf:"ssl_certificates,omitempty"`

	// References to SSLCertificate to populate sslCertificates.
	// +kubebuilder:validation:Optional
	SSLCertificatesRefs []v1.Reference `json:"sslCertificatesRefs,omitempty" tf:"-"`

	// Selector for a list of SSLCertificate to populate sslCertificates.
	// +kubebuilder:validation:Optional
	SSLCertificatesSelector *v1.Selector `json:"sslCertificatesSelector,omitempty" tf:"-"`

	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	// +kubebuilder:validation:Optional
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	// +kubebuilder:validation:Optional
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`
}

// TargetHTTPSProxySpec defines the desired state of TargetHTTPSProxy
type TargetHTTPSProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetHTTPSProxyParameters `json:"forProvider"`
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
	InitProvider TargetHTTPSProxyInitParameters `json:"initProvider,omitempty"`
}

// TargetHTTPSProxyStatus defines the observed state of TargetHTTPSProxy.
type TargetHTTPSProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetHTTPSProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetHTTPSProxy is the Schema for the TargetHTTPSProxys API. Represents a TargetHttpsProxy resource, which is used by one or more global forwarding rule to route incoming HTTPS requests to a URL map.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TargetHTTPSProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.urlMap)",message="urlMap is a required parameter"
	Spec   TargetHTTPSProxySpec   `json:"spec"`
	Status TargetHTTPSProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetHTTPSProxyList contains a list of TargetHTTPSProxys
type TargetHTTPSProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetHTTPSProxy `json:"items"`
}

// Repository type metadata.
var (
	TargetHTTPSProxy_Kind             = "TargetHTTPSProxy"
	TargetHTTPSProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetHTTPSProxy_Kind}.String()
	TargetHTTPSProxy_KindAPIVersion   = TargetHTTPSProxy_Kind + "." + CRDGroupVersion.String()
	TargetHTTPSProxy_GroupVersionKind = CRDGroupVersion.WithKind(TargetHTTPSProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetHTTPSProxy{}, &TargetHTTPSProxyList{})
}
