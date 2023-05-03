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
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this WorkloadIdentityPool.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *WorkloadIdentityPool) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this WorkloadIdentityPool.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *WorkloadIdentityPool) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this WorkloadIdentityPool.
func (mg *WorkloadIdentityPool) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this WorkloadIdentityPoolProvider.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *WorkloadIdentityPoolProvider) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this WorkloadIdentityPoolProvider.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *WorkloadIdentityPoolProvider) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this WorkloadIdentityPoolProvider.
func (mg *WorkloadIdentityPoolProvider) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}