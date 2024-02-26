// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type LiteReservationInitParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`
}

type LiteReservationObservation struct {

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/reservations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite reservation.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`
}

type LiteReservationParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite reservation.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	// +kubebuilder:validation:Optional
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`
}

// LiteReservationSpec defines the desired state of LiteReservation
type LiteReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiteReservationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LiteReservationInitParameters `json:"initProvider,omitempty"`
}

// LiteReservationStatus defines the observed state of LiteReservation.
type LiteReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiteReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LiteReservation is the Schema for the LiteReservations API. A named resource representing a shared pool of capacity.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type LiteReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.throughputCapacity) || (has(self.initProvider) && has(self.initProvider.throughputCapacity))",message="spec.forProvider.throughputCapacity is a required parameter"
	Spec   LiteReservationSpec   `json:"spec"`
	Status LiteReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiteReservationList contains a list of LiteReservations
type LiteReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiteReservation `json:"items"`
}

// Repository type metadata.
var (
	LiteReservation_Kind             = "LiteReservation"
	LiteReservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LiteReservation_Kind}.String()
	LiteReservation_KindAPIVersion   = LiteReservation_Kind + "." + CRDGroupVersion.String()
	LiteReservation_GroupVersionKind = CRDGroupVersion.WithKind(LiteReservation_Kind)
)

func init() {
	SchemeBuilder.Register(&LiteReservation{}, &LiteReservationList{})
}
