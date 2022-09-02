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

type TargetSSLProxyObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/targetSslProxies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier for the resource.
	ProxyID *float64 `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type TargetSSLProxyParameters struct {

	// A reference to the BackendService resource.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.BackendService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// Reference to a BackendService in compute to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceRef *v1.Reference `json:"backendServiceRef,omitempty" tf:"-"`

	// Selector for a BackendService in compute to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceSelector *v1.Selector `json:"backendServiceSelector,omitempty" tf:"-"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is NONE.
	// Possible values are NONE and PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one
	// SSL certificate must be specified.
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
	// the TargetSslProxy resource. If not set, the TargetSslProxy
	// resource will not have any SSL policy configured.
	// +kubebuilder:validation:Optional
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
}

// TargetSSLProxySpec defines the desired state of TargetSSLProxy
type TargetSSLProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetSSLProxyParameters `json:"forProvider"`
}

// TargetSSLProxyStatus defines the observed state of TargetSSLProxy.
type TargetSSLProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetSSLProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetSSLProxy is the Schema for the TargetSSLProxys API. Represents a TargetSslProxy resource, which is used by one or more global forwarding rule to route incoming SSL requests to a backend service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TargetSSLProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetSSLProxySpec   `json:"spec"`
	Status            TargetSSLProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetSSLProxyList contains a list of TargetSSLProxys
type TargetSSLProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetSSLProxy `json:"items"`
}

// Repository type metadata.
var (
	TargetSSLProxy_Kind             = "TargetSSLProxy"
	TargetSSLProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetSSLProxy_Kind}.String()
	TargetSSLProxy_KindAPIVersion   = TargetSSLProxy_Kind + "." + CRDGroupVersion.String()
	TargetSSLProxy_GroupVersionKind = CRDGroupVersion.WithKind(TargetSSLProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetSSLProxy{}, &TargetSSLProxyList{})
}
