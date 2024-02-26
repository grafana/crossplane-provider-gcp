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

type ProjectInitParameters struct {

	// Controls whether the 'default' network exists on the project. Defaults
	// to true, where it is created. Therefore, for quota purposes, you will still need to have 1
	// network slot available to create the project successfully, even if you set auto_create_network to
	// false.googleapis.com on the project to interact
	// with the GCE API and currently leaves it enabled.
	AutoCreateNetwork *bool `json:"autoCreateNetwork,omitempty" tf:"auto_create_network,omitempty"`

	// The alphanumeric ID of the billing account this project
	// belongs to.user) on the billing account.
	// See Google Cloud Billing API Access Control
	// for more details.
	BillingAccount *string `json:"billingAccount,omitempty" tf:"billing_account,omitempty"`

	// The numeric ID of the folder this project should be
	// created under. Only one of org_id or folder_id may be
	// specified. If the folder_id is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in cloudplatform to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the project.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The display name of the project.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// org_id or folder_id may be specified. If the org_id is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	// The numeric ID of the organization this project belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The project ID. Changing this forces a new project to be created.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	SkipDelete *bool `json:"skipDelete,omitempty" tf:"skip_delete,omitempty"`
}

type ProjectObservation struct {

	// Controls whether the 'default' network exists on the project. Defaults
	// to true, where it is created. Therefore, for quota purposes, you will still need to have 1
	// network slot available to create the project successfully, even if you set auto_create_network to
	// false.googleapis.com on the project to interact
	// with the GCE API and currently leaves it enabled.
	AutoCreateNetwork *bool `json:"autoCreateNetwork,omitempty" tf:"auto_create_network,omitempty"`

	// The alphanumeric ID of the billing account this project
	// belongs to.user) on the billing account.
	// See Google Cloud Billing API Access Control
	// for more details.
	BillingAccount *string `json:"billingAccount,omitempty" tf:"billing_account,omitempty"`

	// The numeric ID of the folder this project should be
	// created under. Only one of org_id or folder_id may be
	// specified. If the folder_id is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// an identifier for the resource with format projects/{{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the project.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The display name of the project.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The numeric identifier of the project.
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// org_id or folder_id may be specified. If the org_id is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	// The numeric ID of the organization this project belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The project ID. Changing this forces a new project to be created.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	SkipDelete *bool `json:"skipDelete,omitempty" tf:"skip_delete,omitempty"`
}

type ProjectParameters struct {

	// Controls whether the 'default' network exists on the project. Defaults
	// to true, where it is created. Therefore, for quota purposes, you will still need to have 1
	// network slot available to create the project successfully, even if you set auto_create_network to
	// false.googleapis.com on the project to interact
	// with the GCE API and currently leaves it enabled.
	// +kubebuilder:validation:Optional
	AutoCreateNetwork *bool `json:"autoCreateNetwork,omitempty" tf:"auto_create_network,omitempty"`

	// The alphanumeric ID of the billing account this project
	// belongs to.user) on the billing account.
	// See Google Cloud Billing API Access Control
	// for more details.
	// +kubebuilder:validation:Optional
	BillingAccount *string `json:"billingAccount,omitempty" tf:"billing_account,omitempty"`

	// The numeric ID of the folder this project should be
	// created under. Only one of org_id or folder_id may be
	// specified. If the folder_id is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in cloudplatform to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the project.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The display name of the project.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// org_id or folder_id may be specified. If the org_id is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	// The numeric ID of the organization this project belongs to.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The project ID. Changing this forces a new project to be created.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	SkipDelete *bool `json:"skipDelete,omitempty" tf:"skip_delete,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
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
	InitProvider ProjectInitParameters `json:"initProvider,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Project is the Schema for the Projects API. Allows management of a Google Cloud Platform project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
