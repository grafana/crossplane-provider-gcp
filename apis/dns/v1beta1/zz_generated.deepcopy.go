//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGeoHealthCheckedTargetsObservation) DeepCopyInto(out *BackupGeoHealthCheckedTargetsObservation) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]HealthCheckedTargetsInternalLoadBalancersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGeoHealthCheckedTargetsObservation.
func (in *BackupGeoHealthCheckedTargetsObservation) DeepCopy() *BackupGeoHealthCheckedTargetsObservation {
	if in == nil {
		return nil
	}
	out := new(BackupGeoHealthCheckedTargetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGeoHealthCheckedTargetsParameters) DeepCopyInto(out *BackupGeoHealthCheckedTargetsParameters) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]HealthCheckedTargetsInternalLoadBalancersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGeoHealthCheckedTargetsParameters.
func (in *BackupGeoHealthCheckedTargetsParameters) DeepCopy() *BackupGeoHealthCheckedTargetsParameters {
	if in == nil {
		return nil
	}
	out := new(BackupGeoHealthCheckedTargetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGeoObservation) DeepCopyInto(out *BackupGeoObservation) {
	*out = *in
	if in.HealthCheckedTargets != nil {
		in, out := &in.HealthCheckedTargets, &out.HealthCheckedTargets
		*out = make([]BackupGeoHealthCheckedTargetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGeoObservation.
func (in *BackupGeoObservation) DeepCopy() *BackupGeoObservation {
	if in == nil {
		return nil
	}
	out := new(BackupGeoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGeoParameters) DeepCopyInto(out *BackupGeoParameters) {
	*out = *in
	if in.HealthCheckedTargets != nil {
		in, out := &in.HealthCheckedTargets, &out.HealthCheckedTargets
		*out = make([]BackupGeoHealthCheckedTargetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGeoParameters.
func (in *BackupGeoParameters) DeepCopy() *BackupGeoParameters {
	if in == nil {
		return nil
	}
	out := new(BackupGeoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoObservation) DeepCopyInto(out *GeoObservation) {
	*out = *in
	if in.HealthCheckedTargets != nil {
		in, out := &in.HealthCheckedTargets, &out.HealthCheckedTargets
		*out = make([]HealthCheckedTargetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoObservation.
func (in *GeoObservation) DeepCopy() *GeoObservation {
	if in == nil {
		return nil
	}
	out := new(GeoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoParameters) DeepCopyInto(out *GeoParameters) {
	*out = *in
	if in.HealthCheckedTargets != nil {
		in, out := &in.HealthCheckedTargets, &out.HealthCheckedTargets
		*out = make([]HealthCheckedTargetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoParameters.
func (in *GeoParameters) DeepCopy() *GeoParameters {
	if in == nil {
		return nil
	}
	out := new(GeoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckedTargetsInternalLoadBalancersObservation) DeepCopyInto(out *HealthCheckedTargetsInternalLoadBalancersObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckedTargetsInternalLoadBalancersObservation.
func (in *HealthCheckedTargetsInternalLoadBalancersObservation) DeepCopy() *HealthCheckedTargetsInternalLoadBalancersObservation {
	if in == nil {
		return nil
	}
	out := new(HealthCheckedTargetsInternalLoadBalancersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckedTargetsInternalLoadBalancersParameters) DeepCopyInto(out *HealthCheckedTargetsInternalLoadBalancersParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckedTargetsInternalLoadBalancersParameters.
func (in *HealthCheckedTargetsInternalLoadBalancersParameters) DeepCopy() *HealthCheckedTargetsInternalLoadBalancersParameters {
	if in == nil {
		return nil
	}
	out := new(HealthCheckedTargetsInternalLoadBalancersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckedTargetsObservation) DeepCopyInto(out *HealthCheckedTargetsObservation) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]InternalLoadBalancersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckedTargetsObservation.
func (in *HealthCheckedTargetsObservation) DeepCopy() *HealthCheckedTargetsObservation {
	if in == nil {
		return nil
	}
	out := new(HealthCheckedTargetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckedTargetsParameters) DeepCopyInto(out *HealthCheckedTargetsParameters) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]InternalLoadBalancersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckedTargetsParameters.
func (in *HealthCheckedTargetsParameters) DeepCopy() *HealthCheckedTargetsParameters {
	if in == nil {
		return nil
	}
	out := new(HealthCheckedTargetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalLoadBalancersObservation) DeepCopyInto(out *InternalLoadBalancersObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalLoadBalancersObservation.
func (in *InternalLoadBalancersObservation) DeepCopy() *InternalLoadBalancersObservation {
	if in == nil {
		return nil
	}
	out := new(InternalLoadBalancersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalLoadBalancersParameters) DeepCopyInto(out *InternalLoadBalancersParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalLoadBalancersParameters.
func (in *InternalLoadBalancersParameters) DeepCopy() *InternalLoadBalancersParameters {
	if in == nil {
		return nil
	}
	out := new(InternalLoadBalancersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryBackupObservation) DeepCopyInto(out *PrimaryBackupObservation) {
	*out = *in
	if in.BackupGeo != nil {
		in, out := &in.BackupGeo, &out.BackupGeo
		*out = make([]BackupGeoObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableGeoFencingForBackups != nil {
		in, out := &in.EnableGeoFencingForBackups, &out.EnableGeoFencingForBackups
		*out = new(bool)
		**out = **in
	}
	if in.Primary != nil {
		in, out := &in.Primary, &out.Primary
		*out = make([]PrimaryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrickleRatio != nil {
		in, out := &in.TrickleRatio, &out.TrickleRatio
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryBackupObservation.
func (in *PrimaryBackupObservation) DeepCopy() *PrimaryBackupObservation {
	if in == nil {
		return nil
	}
	out := new(PrimaryBackupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryBackupParameters) DeepCopyInto(out *PrimaryBackupParameters) {
	*out = *in
	if in.BackupGeo != nil {
		in, out := &in.BackupGeo, &out.BackupGeo
		*out = make([]BackupGeoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableGeoFencingForBackups != nil {
		in, out := &in.EnableGeoFencingForBackups, &out.EnableGeoFencingForBackups
		*out = new(bool)
		**out = **in
	}
	if in.Primary != nil {
		in, out := &in.Primary, &out.Primary
		*out = make([]PrimaryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrickleRatio != nil {
		in, out := &in.TrickleRatio, &out.TrickleRatio
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryBackupParameters.
func (in *PrimaryBackupParameters) DeepCopy() *PrimaryBackupParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryBackupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryInternalLoadBalancersObservation) DeepCopyInto(out *PrimaryInternalLoadBalancersObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryInternalLoadBalancersObservation.
func (in *PrimaryInternalLoadBalancersObservation) DeepCopy() *PrimaryInternalLoadBalancersObservation {
	if in == nil {
		return nil
	}
	out := new(PrimaryInternalLoadBalancersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryInternalLoadBalancersParameters) DeepCopyInto(out *PrimaryInternalLoadBalancersParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.NetworkURLRef != nil {
		in, out := &in.NetworkURLRef, &out.NetworkURLRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkURLSelector != nil {
		in, out := &in.NetworkURLSelector, &out.NetworkURLSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryInternalLoadBalancersParameters.
func (in *PrimaryInternalLoadBalancersParameters) DeepCopy() *PrimaryInternalLoadBalancersParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryInternalLoadBalancersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryObservation) DeepCopyInto(out *PrimaryObservation) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]PrimaryInternalLoadBalancersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryObservation.
func (in *PrimaryObservation) DeepCopy() *PrimaryObservation {
	if in == nil {
		return nil
	}
	out := new(PrimaryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryParameters) DeepCopyInto(out *PrimaryParameters) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]PrimaryInternalLoadBalancersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryParameters.
func (in *PrimaryParameters) DeepCopy() *PrimaryParameters {
	if in == nil {
		return nil
	}
	out := new(PrimaryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSet) DeepCopyInto(out *RecordSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSet.
func (in *RecordSet) DeepCopy() *RecordSet {
	if in == nil {
		return nil
	}
	out := new(RecordSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetList) DeepCopyInto(out *RecordSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RecordSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetList.
func (in *RecordSetList) DeepCopy() *RecordSetList {
	if in == nil {
		return nil
	}
	out := new(RecordSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetObservation) DeepCopyInto(out *RecordSetObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ManagedZone != nil {
		in, out := &in.ManagedZone, &out.ManagedZone
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RoutingPolicy != nil {
		in, out := &in.RoutingPolicy, &out.RoutingPolicy
		*out = make([]RoutingPolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetObservation.
func (in *RecordSetObservation) DeepCopy() *RecordSetObservation {
	if in == nil {
		return nil
	}
	out := new(RecordSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetParameters) DeepCopyInto(out *RecordSetParameters) {
	*out = *in
	if in.ManagedZone != nil {
		in, out := &in.ManagedZone, &out.ManagedZone
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RoutingPolicy != nil {
		in, out := &in.RoutingPolicy, &out.RoutingPolicy
		*out = make([]RoutingPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetParameters.
func (in *RecordSetParameters) DeepCopy() *RecordSetParameters {
	if in == nil {
		return nil
	}
	out := new(RecordSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetSpec) DeepCopyInto(out *RecordSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetSpec.
func (in *RecordSetSpec) DeepCopy() *RecordSetSpec {
	if in == nil {
		return nil
	}
	out := new(RecordSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetStatus) DeepCopyInto(out *RecordSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetStatus.
func (in *RecordSetStatus) DeepCopy() *RecordSetStatus {
	if in == nil {
		return nil
	}
	out := new(RecordSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyObservation) DeepCopyInto(out *RoutingPolicyObservation) {
	*out = *in
	if in.EnableGeoFencing != nil {
		in, out := &in.EnableGeoFencing, &out.EnableGeoFencing
		*out = new(bool)
		**out = **in
	}
	if in.Geo != nil {
		in, out := &in.Geo, &out.Geo
		*out = make([]GeoObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrimaryBackup != nil {
		in, out := &in.PrimaryBackup, &out.PrimaryBackup
		*out = make([]PrimaryBackupObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Wrr != nil {
		in, out := &in.Wrr, &out.Wrr
		*out = make([]WrrObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyObservation.
func (in *RoutingPolicyObservation) DeepCopy() *RoutingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyParameters) DeepCopyInto(out *RoutingPolicyParameters) {
	*out = *in
	if in.EnableGeoFencing != nil {
		in, out := &in.EnableGeoFencing, &out.EnableGeoFencing
		*out = new(bool)
		**out = **in
	}
	if in.Geo != nil {
		in, out := &in.Geo, &out.Geo
		*out = make([]GeoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrimaryBackup != nil {
		in, out := &in.PrimaryBackup, &out.PrimaryBackup
		*out = make([]PrimaryBackupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Wrr != nil {
		in, out := &in.Wrr, &out.Wrr
		*out = make([]WrrParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyParameters.
func (in *RoutingPolicyParameters) DeepCopy() *RoutingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrHealthCheckedTargetsInternalLoadBalancersObservation) DeepCopyInto(out *WrrHealthCheckedTargetsInternalLoadBalancersObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrHealthCheckedTargetsInternalLoadBalancersObservation.
func (in *WrrHealthCheckedTargetsInternalLoadBalancersObservation) DeepCopy() *WrrHealthCheckedTargetsInternalLoadBalancersObservation {
	if in == nil {
		return nil
	}
	out := new(WrrHealthCheckedTargetsInternalLoadBalancersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrHealthCheckedTargetsInternalLoadBalancersParameters) DeepCopyInto(out *WrrHealthCheckedTargetsInternalLoadBalancersParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPProtocol != nil {
		in, out := &in.IPProtocol, &out.IPProtocol
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerType != nil {
		in, out := &in.LoadBalancerType, &out.LoadBalancerType
		*out = new(string)
		**out = **in
	}
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrHealthCheckedTargetsInternalLoadBalancersParameters.
func (in *WrrHealthCheckedTargetsInternalLoadBalancersParameters) DeepCopy() *WrrHealthCheckedTargetsInternalLoadBalancersParameters {
	if in == nil {
		return nil
	}
	out := new(WrrHealthCheckedTargetsInternalLoadBalancersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrHealthCheckedTargetsObservation) DeepCopyInto(out *WrrHealthCheckedTargetsObservation) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]WrrHealthCheckedTargetsInternalLoadBalancersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrHealthCheckedTargetsObservation.
func (in *WrrHealthCheckedTargetsObservation) DeepCopy() *WrrHealthCheckedTargetsObservation {
	if in == nil {
		return nil
	}
	out := new(WrrHealthCheckedTargetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrHealthCheckedTargetsParameters) DeepCopyInto(out *WrrHealthCheckedTargetsParameters) {
	*out = *in
	if in.InternalLoadBalancers != nil {
		in, out := &in.InternalLoadBalancers, &out.InternalLoadBalancers
		*out = make([]WrrHealthCheckedTargetsInternalLoadBalancersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrHealthCheckedTargetsParameters.
func (in *WrrHealthCheckedTargetsParameters) DeepCopy() *WrrHealthCheckedTargetsParameters {
	if in == nil {
		return nil
	}
	out := new(WrrHealthCheckedTargetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrObservation) DeepCopyInto(out *WrrObservation) {
	*out = *in
	if in.HealthCheckedTargets != nil {
		in, out := &in.HealthCheckedTargets, &out.HealthCheckedTargets
		*out = make([]WrrHealthCheckedTargetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrObservation.
func (in *WrrObservation) DeepCopy() *WrrObservation {
	if in == nil {
		return nil
	}
	out := new(WrrObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WrrParameters) DeepCopyInto(out *WrrParameters) {
	*out = *in
	if in.HealthCheckedTargets != nil {
		in, out := &in.HealthCheckedTargets, &out.HealthCheckedTargets
		*out = make([]WrrHealthCheckedTargetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WrrParameters.
func (in *WrrParameters) DeepCopy() *WrrParameters {
	if in == nil {
		return nil
	}
	out := new(WrrParameters)
	in.DeepCopyInto(out)
	return out
}
