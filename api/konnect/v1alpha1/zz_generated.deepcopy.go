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

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfiguration) DeepCopyInto(out *KonnectAPIAuthConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfiguration.
func (in *KonnectAPIAuthConfiguration) DeepCopy() *KonnectAPIAuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectAPIAuthConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationList) DeepCopyInto(out *KonnectAPIAuthConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectAPIAuthConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationList.
func (in *KonnectAPIAuthConfigurationList) DeepCopy() *KonnectAPIAuthConfigurationList {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectAPIAuthConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationRef) DeepCopyInto(out *KonnectAPIAuthConfigurationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationRef.
func (in *KonnectAPIAuthConfigurationRef) DeepCopy() *KonnectAPIAuthConfigurationRef {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationSpec) DeepCopyInto(out *KonnectAPIAuthConfigurationSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationSpec.
func (in *KonnectAPIAuthConfigurationSpec) DeepCopy() *KonnectAPIAuthConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationStatus) DeepCopyInto(out *KonnectAPIAuthConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationStatus.
func (in *KonnectAPIAuthConfigurationStatus) DeepCopy() *KonnectAPIAuthConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayNetwork) DeepCopyInto(out *KonnectCloudGatewayNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayNetwork.
func (in *KonnectCloudGatewayNetwork) DeepCopy() *KonnectCloudGatewayNetwork {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectCloudGatewayNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayNetworkList) DeepCopyInto(out *KonnectCloudGatewayNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectCloudGatewayNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayNetworkList.
func (in *KonnectCloudGatewayNetworkList) DeepCopy() *KonnectCloudGatewayNetworkList {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectCloudGatewayNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayNetworkStatus) DeepCopyInto(out *KonnectCloudGatewayNetworkStatus) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayNetworkStatus.
func (in *KonnectCloudGatewayNetworkStatus) DeepCopy() *KonnectCloudGatewayNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatus) DeepCopyInto(out *KonnectEntityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatus.
func (in *KonnectEntityStatus) DeepCopy() *KonnectEntityStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatusWithControlPlaneAndCertificateRefs) DeepCopyInto(out *KonnectEntityStatusWithControlPlaneAndCertificateRefs) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatusWithControlPlaneAndCertificateRefs.
func (in *KonnectEntityStatusWithControlPlaneAndCertificateRefs) DeepCopy() *KonnectEntityStatusWithControlPlaneAndCertificateRefs {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatusWithControlPlaneAndCertificateRefs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatusWithControlPlaneAndConsumerRefs) DeepCopyInto(out *KonnectEntityStatusWithControlPlaneAndConsumerRefs) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatusWithControlPlaneAndConsumerRefs.
func (in *KonnectEntityStatusWithControlPlaneAndConsumerRefs) DeepCopy() *KonnectEntityStatusWithControlPlaneAndConsumerRefs {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatusWithControlPlaneAndConsumerRefs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatusWithControlPlaneAndKeySetRef) DeepCopyInto(out *KonnectEntityStatusWithControlPlaneAndKeySetRef) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatusWithControlPlaneAndKeySetRef.
func (in *KonnectEntityStatusWithControlPlaneAndKeySetRef) DeepCopy() *KonnectEntityStatusWithControlPlaneAndKeySetRef {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatusWithControlPlaneAndKeySetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatusWithControlPlaneAndServiceRefs) DeepCopyInto(out *KonnectEntityStatusWithControlPlaneAndServiceRefs) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatusWithControlPlaneAndServiceRefs.
func (in *KonnectEntityStatusWithControlPlaneAndServiceRefs) DeepCopy() *KonnectEntityStatusWithControlPlaneAndServiceRefs {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatusWithControlPlaneAndServiceRefs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatusWithControlPlaneAndUpstreamRefs) DeepCopyInto(out *KonnectEntityStatusWithControlPlaneAndUpstreamRefs) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatusWithControlPlaneAndUpstreamRefs.
func (in *KonnectEntityStatusWithControlPlaneAndUpstreamRefs) DeepCopy() *KonnectEntityStatusWithControlPlaneAndUpstreamRefs {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatusWithControlPlaneAndUpstreamRefs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEntityStatusWithControlPlaneRef) DeepCopyInto(out *KonnectEntityStatusWithControlPlaneRef) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEntityStatusWithControlPlaneRef.
func (in *KonnectEntityStatusWithControlPlaneRef) DeepCopy() *KonnectEntityStatusWithControlPlaneRef {
	if in == nil {
		return nil
	}
	out := new(KonnectEntityStatusWithControlPlaneRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectGatewayControlPlane) DeepCopyInto(out *KonnectGatewayControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectGatewayControlPlane.
func (in *KonnectGatewayControlPlane) DeepCopy() *KonnectGatewayControlPlane {
	if in == nil {
		return nil
	}
	out := new(KonnectGatewayControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectGatewayControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectGatewayControlPlaneList) DeepCopyInto(out *KonnectGatewayControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectGatewayControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectGatewayControlPlaneList.
func (in *KonnectGatewayControlPlaneList) DeepCopy() *KonnectGatewayControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(KonnectGatewayControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectGatewayControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectGatewayControlPlaneSpec) DeepCopyInto(out *KonnectGatewayControlPlaneSpec) {
	*out = *in
	in.CreateControlPlaneRequest.DeepCopyInto(&out.CreateControlPlaneRequest)
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	out.KonnectConfiguration = in.KonnectConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectGatewayControlPlaneSpec.
func (in *KonnectGatewayControlPlaneSpec) DeepCopy() *KonnectGatewayControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectGatewayControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectGatewayControlPlaneStatus) DeepCopyInto(out *KonnectGatewayControlPlaneStatus) {
	*out = *in
	out.KonnectEntityStatus = in.KonnectEntityStatus
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectGatewayControlPlaneStatus.
func (in *KonnectGatewayControlPlaneStatus) DeepCopy() *KonnectGatewayControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectGatewayControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectNetworkSpec) DeepCopyInto(out *KonnectNetworkSpec) {
	*out = *in
	in.CreateNetworkRequest.DeepCopyInto(&out.CreateNetworkRequest)
	out.KonnectConfiguration = in.KonnectConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectNetworkSpec.
func (in *KonnectNetworkSpec) DeepCopy() *KonnectNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectNetworkSpec)
	in.DeepCopyInto(out)
	return out
}
