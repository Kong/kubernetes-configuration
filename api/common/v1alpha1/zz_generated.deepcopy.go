//go:build !ignore_autogenerated

/*
Copyright 2021 Kong, Inc.

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

package v1alpha1

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneRef) DeepCopyInto(out *ControlPlaneRef) {
	*out = *in
	if in.KonnectID != nil {
		in, out := &in.KonnectID, &out.KonnectID
		*out = new(KonnectIDType)
		**out = **in
	}
	if in.KonnectNamespacedRef != nil {
		in, out := &in.KonnectNamespacedRef, &out.KonnectNamespacedRef
		*out = new(KonnectNamespacedRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneRef.
func (in *ControlPlaneRef) DeepCopy() *ControlPlaneRef {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionRef) DeepCopyInto(out *ExtensionRef) {
	*out = *in
	in.NamespacedRef.DeepCopyInto(&out.NamespacedRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionRef.
func (in *ExtensionRef) DeepCopy() *ExtensionRef {
	if in == nil {
		return nil
	}
	out := new(ExtensionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectNamespacedRef) DeepCopyInto(out *KonnectNamespacedRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectNamespacedRef.
func (in *KonnectNamespacedRef) DeepCopy() *KonnectNamespacedRef {
	if in == nil {
		return nil
	}
	out := new(KonnectNamespacedRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameRef) DeepCopyInto(out *NameRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameRef.
func (in *NameRef) DeepCopy() *NameRef {
	if in == nil {
		return nil
	}
	out := new(NameRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedRef) DeepCopyInto(out *NamespacedRef) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedRef.
func (in *NamespacedRef) DeepCopy() *NamespacedRef {
	if in == nil {
		return nil
	}
	out := new(NamespacedRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
	if in.KonnectID != nil {
		in, out := &in.KonnectID, &out.KonnectID
		*out = new(string)
		**out = **in
	}
	if in.NamespacedRef != nil {
		in, out := &in.NamespacedRef, &out.NamespacedRef
		*out = new(NamespacedRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Tags) DeepCopyInto(out *Tags) {
	{
		in := &in
		*out = make(Tags, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in Tags) DeepCopy() Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return *out
}
