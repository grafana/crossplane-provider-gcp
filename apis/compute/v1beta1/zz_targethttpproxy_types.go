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

type TargetHTTPProxyInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `json:"proxyBind,omitempty" tf:"proxy_bind,omitempty"`
}

type TargetHTTPProxyObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HTTPKeepAliveTimeoutSec *float64 `json:"httpKeepAliveTimeoutSec,omitempty" tf:"http_keep_alive_timeout_sec,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/targetHttpProxies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `json:"proxyBind,omitempty" tf:"proxy_bind,omitempty"`

	// The unique identifier for the resource.
	ProxyID *float64 `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`
}

type TargetHTTPProxyParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	// +kubebuilder:validation:Optional
	HTTPKeepAliveTimeoutSec *float64 `json:"httpKeepAliveTimeoutSec,omitempty" tf:"http_keep_alive_timeout_sec,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// +kubebuilder:validation:Optional
	ProxyBind *bool `json:"proxyBind,omitempty" tf:"proxy_bind,omitempty"`

	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	// +kubebuilder:validation:Optional
	URLMap *string `json:"urlMap,omitempty" tf:"url_map,omitempty"`
}

// TargetHTTPProxySpec defines the desired state of TargetHTTPProxy
type TargetHTTPProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetHTTPProxyParameters `json:"forProvider"`
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
	InitProvider TargetHTTPProxyInitParameters `json:"initProvider,omitempty"`
}

// TargetHTTPProxyStatus defines the observed state of TargetHTTPProxy.
type TargetHTTPProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetHTTPProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetHTTPProxy is the Schema for the TargetHTTPProxys API. Represents a TargetHttpProxy resource, which is used by one or more global forwarding rule to route incoming HTTP requests to a URL map.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TargetHTTPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.urlMap)",message="urlMap is a required parameter"
	Spec   TargetHTTPProxySpec   `json:"spec"`
	Status TargetHTTPProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetHTTPProxyList contains a list of TargetHTTPProxys
type TargetHTTPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetHTTPProxy `json:"items"`
}

// Repository type metadata.
var (
	TargetHTTPProxy_Kind             = "TargetHTTPProxy"
	TargetHTTPProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetHTTPProxy_Kind}.String()
	TargetHTTPProxy_KindAPIVersion   = TargetHTTPProxy_Kind + "." + CRDGroupVersion.String()
	TargetHTTPProxy_GroupVersionKind = CRDGroupVersion.WithKind(TargetHTTPProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetHTTPProxy{}, &TargetHTTPProxyList{})
}
