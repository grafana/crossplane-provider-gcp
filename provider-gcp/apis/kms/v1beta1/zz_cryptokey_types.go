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

type CryptoKeyObservation struct {

	// an identifier for the resource with format {{key_ring}}/cryptoKeys/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CryptoKeyParameters struct {

	// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	// If not specified at creation time, the default duration is 24 hours.
	// +kubebuilder:validation:Optional
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty" tf:"destroy_scheduled_duration,omitempty"`

	// Whether this key may contain imported versions only.
	// +kubebuilder:validation:Optional
	ImportOnly *bool `json:"importOnly,omitempty" tf:"import_only,omitempty"`

	// The KeyRing that this key belongs to.
	// Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/kms/v1beta1.KeyRing
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyRing *string `json:"keyRing,omitempty" tf:"key_ring,omitempty"`

	// Reference to a KeyRing in kms to populate keyRing.
	// +kubebuilder:validation:Optional
	KeyRingRef *v1.Reference `json:"keyRingRef,omitempty" tf:"-"`

	// Selector for a KeyRing in kms to populate keyRing.
	// +kubebuilder:validation:Optional
	KeyRingSelector *v1.Selector `json:"keyRingSelector,omitempty" tf:"-"`

	// Labels with user-defined metadata to apply to this resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The immutable purpose of this CryptoKey. See the
	// purpose reference
	// for possible inputs.
	// Default value is ENCRYPT_DECRYPT.
	// Possible values are ENCRYPT_DECRYPT, ASYMMETRIC_SIGN, and ASYMMETRIC_DECRYPT.
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter s . It must be greater than a day .
	// +kubebuilder:validation:Optional
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the google_kms_key_ring_import_job resource to import the CryptoKeyVersion.
	// +kubebuilder:validation:Optional
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty" tf:"skip_initial_version_creation,omitempty"`

	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VersionTemplate []VersionTemplateParameters `json:"versionTemplate,omitempty" tf:"version_template,omitempty"`
}

type VersionTemplateObservation struct {
}

type VersionTemplateParameters struct {

	// The algorithm to use when creating a version based on this template.
	// See the algorithm reference for possible inputs.
	// +kubebuilder:validation:Required
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL". Defaults to "SOFTWARE".
	// +kubebuilder:validation:Optional
	ProtectionLevel *string `json:"protectionLevel,omitempty" tf:"protection_level,omitempty"`
}

// CryptoKeySpec defines the desired state of CryptoKey
type CryptoKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CryptoKeyParameters `json:"forProvider"`
}

// CryptoKeyStatus defines the observed state of CryptoKey.
type CryptoKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CryptoKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKey is the Schema for the CryptoKeys API. A
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CryptoKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CryptoKeySpec   `json:"spec"`
	Status            CryptoKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyList contains a list of CryptoKeys
type CryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CryptoKey `json:"items"`
}

// Repository type metadata.
var (
	CryptoKey_Kind             = "CryptoKey"
	CryptoKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CryptoKey_Kind}.String()
	CryptoKey_KindAPIVersion   = CryptoKey_Kind + "." + CRDGroupVersion.String()
	CryptoKey_GroupVersionKind = CRDGroupVersion.WithKind(CryptoKey_Kind)
)

func init() {
	SchemeBuilder.Register(&CryptoKey{}, &CryptoKeyList{})
}
