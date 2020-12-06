// +build !ignore_autogenerated

/*
Copyright 2020 IBM Co..

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
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/common"
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/utils/status"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ByAlphabetical) DeepCopyInto(out *ByAlphabetical) {
	{
		in := &in
		*out = make(ByAlphabetical, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByAlphabetical.
func (in ByAlphabetical) DeepCopy() ByAlphabetical {
	if in == nil {
		return nil
	}
	out := new(ByAlphabetical)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSVNamespacedName) DeepCopyInto(out *CSVNamespacedName) {
	*out = *in
	if in.GroupVersionKind != nil {
		in, out := &in.GroupVersionKind, &out.GroupVersionKind
		*out = new(common.GroupVersionKind)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSVNamespacedName.
func (in *CSVNamespacedName) DeepCopy() *CSVNamespacedName {
	if in == nil {
		return nil
	}
	out := new(CSVNamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterDefinition) DeepCopyInto(out *MeterDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterDefinition.
func (in *MeterDefinition) DeepCopy() *MeterDefinition {
	if in == nil {
		return nil
	}
	out := new(MeterDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeterDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterDefinitionList) DeepCopyInto(out *MeterDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MeterDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterDefinitionList.
func (in *MeterDefinitionList) DeepCopy() *MeterDefinitionList {
	if in == nil {
		return nil
	}
	out := new(MeterDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeterDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterDefinitionSpec) DeepCopyInto(out *MeterDefinitionSpec) {
	*out = *in
	if in.InstalledBy != nil {
		in, out := &in.InstalledBy, &out.InstalledBy
		*out = new(common.NamespacedNameReference)
		(*in).DeepCopyInto(*out)
	}
	if in.VertexLabelSelector != nil {
		in, out := &in.VertexLabelSelector, &out.VertexLabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ServiceMeterLabels != nil {
		in, out := &in.ServiceMeterLabels, &out.ServiceMeterLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodMeterLabels != nil {
		in, out := &in.PodMeterLabels, &out.PodMeterLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterDefinitionSpec.
func (in *MeterDefinitionSpec) DeepCopy() *MeterDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(MeterDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterDefinitionStatus) DeepCopyInto(out *MeterDefinitionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkloadResources != nil {
		in, out := &in.WorkloadResources, &out.WorkloadResources
		*out = make([]WorkloadResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterDefinitionStatus.
func (in *MeterDefinitionStatus) DeepCopy() *MeterDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(MeterDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeterLabelQuery) DeepCopyInto(out *MeterLabelQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeterLabelQuery.
func (in *MeterLabelQuery) DeepCopy() *MeterLabelQuery {
	if in == nil {
		return nil
	}
	out := new(MeterLabelQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	if in.OwnerCRD != nil {
		in, out := &in.OwnerCRD, &out.OwnerCRD
		*out = new(common.GroupVersionKind)
		**out = **in
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AnnotationSelector != nil {
		in, out := &in.AnnotationSelector, &out.AnnotationSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricLabels != nil {
		in, out := &in.MetricLabels, &out.MetricLabels
		*out = make([]MeterLabelQuery, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadResource) DeepCopyInto(out *WorkloadResource) {
	*out = *in
	in.NamespacedNameReference.DeepCopyInto(&out.NamespacedNameReference)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadResource.
func (in *WorkloadResource) DeepCopy() *WorkloadResource {
	if in == nil {
		return nil
	}
	out := new(WorkloadResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	in.LastReadTime.DeepCopyInto(&out.LastReadTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}
