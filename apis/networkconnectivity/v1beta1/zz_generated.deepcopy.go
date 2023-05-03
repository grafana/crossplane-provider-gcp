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
func (in *Hub) DeepCopyInto(out *Hub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hub.
func (in *Hub) DeepCopy() *Hub {
	if in == nil {
		return nil
	}
	out := new(Hub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubList) DeepCopyInto(out *HubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubList.
func (in *HubList) DeepCopy() *HubList {
	if in == nil {
		return nil
	}
	out := new(HubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubObservation) DeepCopyInto(out *HubObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
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
	if in.RoutingVpcs != nil {
		in, out := &in.RoutingVpcs, &out.RoutingVpcs
		*out = make([]RoutingVpcsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.UniqueID != nil {
		in, out := &in.UniqueID, &out.UniqueID
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubObservation.
func (in *HubObservation) DeepCopy() *HubObservation {
	if in == nil {
		return nil
	}
	out := new(HubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubParameters) DeepCopyInto(out *HubParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubParameters.
func (in *HubParameters) DeepCopy() *HubParameters {
	if in == nil {
		return nil
	}
	out := new(HubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubSpec) DeepCopyInto(out *HubSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubSpec.
func (in *HubSpec) DeepCopy() *HubSpec {
	if in == nil {
		return nil
	}
	out := new(HubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubStatus) DeepCopyInto(out *HubStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubStatus.
func (in *HubStatus) DeepCopy() *HubStatus {
	if in == nil {
		return nil
	}
	out := new(HubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancesObservation) DeepCopyInto(out *InstancesObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachine != nil {
		in, out := &in.VirtualMachine, &out.VirtualMachine
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancesObservation.
func (in *InstancesObservation) DeepCopy() *InstancesObservation {
	if in == nil {
		return nil
	}
	out := new(InstancesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancesParameters) DeepCopyInto(out *InstancesParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachine != nil {
		in, out := &in.VirtualMachine, &out.VirtualMachine
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineRef != nil {
		in, out := &in.VirtualMachineRef, &out.VirtualMachineRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualMachineSelector != nil {
		in, out := &in.VirtualMachineSelector, &out.VirtualMachineSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancesParameters.
func (in *InstancesParameters) DeepCopy() *InstancesParameters {
	if in == nil {
		return nil
	}
	out := new(InstancesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedInterconnectAttachmentsObservation) DeepCopyInto(out *LinkedInterconnectAttachmentsObservation) {
	*out = *in
	if in.SiteToSiteDataTransfer != nil {
		in, out := &in.SiteToSiteDataTransfer, &out.SiteToSiteDataTransfer
		*out = new(bool)
		**out = **in
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedInterconnectAttachmentsObservation.
func (in *LinkedInterconnectAttachmentsObservation) DeepCopy() *LinkedInterconnectAttachmentsObservation {
	if in == nil {
		return nil
	}
	out := new(LinkedInterconnectAttachmentsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedInterconnectAttachmentsParameters) DeepCopyInto(out *LinkedInterconnectAttachmentsParameters) {
	*out = *in
	if in.SiteToSiteDataTransfer != nil {
		in, out := &in.SiteToSiteDataTransfer, &out.SiteToSiteDataTransfer
		*out = new(bool)
		**out = **in
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedInterconnectAttachmentsParameters.
func (in *LinkedInterconnectAttachmentsParameters) DeepCopy() *LinkedInterconnectAttachmentsParameters {
	if in == nil {
		return nil
	}
	out := new(LinkedInterconnectAttachmentsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedRouterApplianceInstancesObservation) DeepCopyInto(out *LinkedRouterApplianceInstancesObservation) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]InstancesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SiteToSiteDataTransfer != nil {
		in, out := &in.SiteToSiteDataTransfer, &out.SiteToSiteDataTransfer
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedRouterApplianceInstancesObservation.
func (in *LinkedRouterApplianceInstancesObservation) DeepCopy() *LinkedRouterApplianceInstancesObservation {
	if in == nil {
		return nil
	}
	out := new(LinkedRouterApplianceInstancesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedRouterApplianceInstancesParameters) DeepCopyInto(out *LinkedRouterApplianceInstancesParameters) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]InstancesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SiteToSiteDataTransfer != nil {
		in, out := &in.SiteToSiteDataTransfer, &out.SiteToSiteDataTransfer
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedRouterApplianceInstancesParameters.
func (in *LinkedRouterApplianceInstancesParameters) DeepCopy() *LinkedRouterApplianceInstancesParameters {
	if in == nil {
		return nil
	}
	out := new(LinkedRouterApplianceInstancesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedVPNTunnelsObservation) DeepCopyInto(out *LinkedVPNTunnelsObservation) {
	*out = *in
	if in.SiteToSiteDataTransfer != nil {
		in, out := &in.SiteToSiteDataTransfer, &out.SiteToSiteDataTransfer
		*out = new(bool)
		**out = **in
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedVPNTunnelsObservation.
func (in *LinkedVPNTunnelsObservation) DeepCopy() *LinkedVPNTunnelsObservation {
	if in == nil {
		return nil
	}
	out := new(LinkedVPNTunnelsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedVPNTunnelsParameters) DeepCopyInto(out *LinkedVPNTunnelsParameters) {
	*out = *in
	if in.SiteToSiteDataTransfer != nil {
		in, out := &in.SiteToSiteDataTransfer, &out.SiteToSiteDataTransfer
		*out = new(bool)
		**out = **in
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedVPNTunnelsParameters.
func (in *LinkedVPNTunnelsParameters) DeepCopy() *LinkedVPNTunnelsParameters {
	if in == nil {
		return nil
	}
	out := new(LinkedVPNTunnelsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingVpcsObservation) DeepCopyInto(out *RoutingVpcsObservation) {
	*out = *in
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingVpcsObservation.
func (in *RoutingVpcsObservation) DeepCopy() *RoutingVpcsObservation {
	if in == nil {
		return nil
	}
	out := new(RoutingVpcsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingVpcsParameters) DeepCopyInto(out *RoutingVpcsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingVpcsParameters.
func (in *RoutingVpcsParameters) DeepCopy() *RoutingVpcsParameters {
	if in == nil {
		return nil
	}
	out := new(RoutingVpcsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spoke) DeepCopyInto(out *Spoke) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spoke.
func (in *Spoke) DeepCopy() *Spoke {
	if in == nil {
		return nil
	}
	out := new(Spoke)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spoke) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeList) DeepCopyInto(out *SpokeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spoke, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeList.
func (in *SpokeList) DeepCopy() *SpokeList {
	if in == nil {
		return nil
	}
	out := new(SpokeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpokeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeObservation) DeepCopyInto(out *SpokeObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Hub != nil {
		in, out := &in.Hub, &out.Hub
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LinkedInterconnectAttachments != nil {
		in, out := &in.LinkedInterconnectAttachments, &out.LinkedInterconnectAttachments
		*out = make([]LinkedInterconnectAttachmentsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedRouterApplianceInstances != nil {
		in, out := &in.LinkedRouterApplianceInstances, &out.LinkedRouterApplianceInstances
		*out = make([]LinkedRouterApplianceInstancesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedVPNTunnels != nil {
		in, out := &in.LinkedVPNTunnels, &out.LinkedVPNTunnels
		*out = make([]LinkedVPNTunnelsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
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
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.UniqueID != nil {
		in, out := &in.UniqueID, &out.UniqueID
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeObservation.
func (in *SpokeObservation) DeepCopy() *SpokeObservation {
	if in == nil {
		return nil
	}
	out := new(SpokeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeParameters) DeepCopyInto(out *SpokeParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Hub != nil {
		in, out := &in.Hub, &out.Hub
		*out = new(string)
		**out = **in
	}
	if in.HubRef != nil {
		in, out := &in.HubRef, &out.HubRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.HubSelector != nil {
		in, out := &in.HubSelector, &out.HubSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LinkedInterconnectAttachments != nil {
		in, out := &in.LinkedInterconnectAttachments, &out.LinkedInterconnectAttachments
		*out = make([]LinkedInterconnectAttachmentsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedRouterApplianceInstances != nil {
		in, out := &in.LinkedRouterApplianceInstances, &out.LinkedRouterApplianceInstances
		*out = make([]LinkedRouterApplianceInstancesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedVPNTunnels != nil {
		in, out := &in.LinkedVPNTunnels, &out.LinkedVPNTunnels
		*out = make([]LinkedVPNTunnelsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeParameters.
func (in *SpokeParameters) DeepCopy() *SpokeParameters {
	if in == nil {
		return nil
	}
	out := new(SpokeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeSpec) DeepCopyInto(out *SpokeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeSpec.
func (in *SpokeSpec) DeepCopy() *SpokeSpec {
	if in == nil {
		return nil
	}
	out := new(SpokeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeStatus) DeepCopyInto(out *SpokeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeStatus.
func (in *SpokeStatus) DeepCopy() *SpokeStatus {
	if in == nil {
		return nil
	}
	out := new(SpokeStatus)
	in.DeepCopyInto(out)
	return out
}