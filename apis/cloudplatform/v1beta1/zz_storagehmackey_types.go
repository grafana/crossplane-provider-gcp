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

type StorageHMACKeyObservation struct {

	// The access ID of the HMAC Key.
	AccessID *string `json:"accessId,omitempty" tf:"access_id,omitempty"`

	// an identifier for the resource with format projects/{{project}}/hmacKeys/{{access_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The email address of the key's associated service account.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is ACTIVE.
	// Possible values are: ACTIVE, INACTIVE.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type StorageHMACKeyParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The email address of the key's associated service account.
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractEmail()
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is ACTIVE.
	// Possible values are: ACTIVE, INACTIVE.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

// StorageHMACKeySpec defines the desired state of StorageHMACKey
type StorageHMACKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageHMACKeyParameters `json:"forProvider"`
}

// StorageHMACKeyStatus defines the observed state of StorageHMACKey.
type StorageHMACKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageHMACKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageHMACKey is the Schema for the StorageHMACKeys API. The hmacKeys resource represents an HMAC key within Cloud Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type StorageHMACKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageHMACKeySpec   `json:"spec"`
	Status            StorageHMACKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageHMACKeyList contains a list of StorageHMACKeys
type StorageHMACKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageHMACKey `json:"items"`
}

// Repository type metadata.
var (
	StorageHMACKey_Kind             = "StorageHMACKey"
	StorageHMACKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageHMACKey_Kind}.String()
	StorageHMACKey_KindAPIVersion   = StorageHMACKey_Kind + "." + CRDGroupVersion.String()
	StorageHMACKey_GroupVersionKind = CRDGroupVersion.WithKind(StorageHMACKey_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageHMACKey{}, &StorageHMACKeyList{})
}