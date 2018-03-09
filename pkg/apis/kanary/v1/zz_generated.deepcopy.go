// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationStatus) DeepCopyInto(out *DestinationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationStatus.
func (in *DestinationStatus) DeepCopy() *DestinationStatus {
	if in == nil {
		return nil
	}
	out := new(DestinationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kanary) DeepCopyInto(out *Kanary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kanary.
func (in *Kanary) DeepCopy() *Kanary {
	if in == nil {
		return nil
	}
	out := new(Kanary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kanary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanaryList) DeepCopyInto(out *KanaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kanary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanaryList.
func (in *KanaryList) DeepCopy() *KanaryList {
	if in == nil {
		return nil
	}
	out := new(KanaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KanaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanaryRoute) DeepCopyInto(out *KanaryRoute) {
	*out = *in
	in.Backend.DeepCopyInto(&out.Backend)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanaryRoute.
func (in *KanaryRoute) DeepCopy() *KanaryRoute {
	if in == nil {
		return nil
	}
	out := new(KanaryRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanaryService) DeepCopyInto(out *KanaryService) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanaryService.
func (in *KanaryService) DeepCopy() *KanaryService {
	if in == nil {
		return nil
	}
	out := new(KanaryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanarySpec) DeepCopyInto(out *KanarySpec) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]KanaryRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanarySpec.
func (in *KanarySpec) DeepCopy() *KanarySpec {
	if in == nil {
		return nil
	}
	out := new(KanarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanaryStatus) DeepCopyInto(out *KanaryStatus) {
	*out = *in
	out.DestinationStatus = in.DestinationStatus
	if in.RoutesStatus != nil {
		in, out := &in.RoutesStatus, &out.RoutesStatus
		*out = make([]RouteStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanaryStatus.
func (in *KanaryStatus) DeepCopy() *KanaryStatus {
	if in == nil {
		return nil
	}
	out := new(KanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}
