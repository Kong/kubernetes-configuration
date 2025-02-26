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
	"github.com/Kong/sdk-konnect-go/models/components"
	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSecret) DeepCopyInto(out *CertificateSecret) {
	*out = *in
	if in.Provisioning != nil {
		in, out := &in.Provisioning, &out.Provisioning
		*out = new(ProvisioningMethod)
		**out = **in
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSecret.
func (in *CertificateSecret) DeepCopy() *CertificateSecret {
	if in == nil {
		return nil
	}
	out := new(CertificateSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationDataPlaneGroupAutoscale) DeepCopyInto(out *ConfigurationDataPlaneGroupAutoscale) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(ConfigurationDataPlaneGroupAutoscaleStatic)
		**out = **in
	}
	if in.Autopilot != nil {
		in, out := &in.Autopilot, &out.Autopilot
		*out = new(ConfigurationDataPlaneGroupAutoscaleAutopilot)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationDataPlaneGroupAutoscale.
func (in *ConfigurationDataPlaneGroupAutoscale) DeepCopy() *ConfigurationDataPlaneGroupAutoscale {
	if in == nil {
		return nil
	}
	out := new(ConfigurationDataPlaneGroupAutoscale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationDataPlaneGroupAutoscaleAutopilot) DeepCopyInto(out *ConfigurationDataPlaneGroupAutoscaleAutopilot) {
	*out = *in
	if in.MaxRps != nil {
		in, out := &in.MaxRps, &out.MaxRps
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationDataPlaneGroupAutoscaleAutopilot.
func (in *ConfigurationDataPlaneGroupAutoscaleAutopilot) DeepCopy() *ConfigurationDataPlaneGroupAutoscaleAutopilot {
	if in == nil {
		return nil
	}
	out := new(ConfigurationDataPlaneGroupAutoscaleAutopilot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationDataPlaneGroupAutoscaleStatic) DeepCopyInto(out *ConfigurationDataPlaneGroupAutoscaleStatic) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationDataPlaneGroupAutoscaleStatic.
func (in *ConfigurationDataPlaneGroupAutoscaleStatic) DeepCopy() *ConfigurationDataPlaneGroupAutoscaleStatic {
	if in == nil {
		return nil
	}
	out := new(ConfigurationDataPlaneGroupAutoscaleStatic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationDataPlaneGroupEnvironmentField) DeepCopyInto(out *ConfigurationDataPlaneGroupEnvironmentField) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationDataPlaneGroupEnvironmentField.
func (in *ConfigurationDataPlaneGroupEnvironmentField) DeepCopy() *ConfigurationDataPlaneGroupEnvironmentField {
	if in == nil {
		return nil
	}
	out := new(ConfigurationDataPlaneGroupEnvironmentField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneClientAuth) DeepCopyInto(out *DataPlaneClientAuth) {
	*out = *in
	in.CertificateSecret.DeepCopyInto(&out.CertificateSecret)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneClientAuth.
func (in *DataPlaneClientAuth) DeepCopy() *DataPlaneClientAuth {
	if in == nil {
		return nil
	}
	out := new(DataPlaneClientAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneClientAuthStatus) DeepCopyInto(out *DataPlaneClientAuthStatus) {
	*out = *in
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneClientAuthStatus.
func (in *DataPlaneClientAuthStatus) DeepCopy() *DataPlaneClientAuthStatus {
	if in == nil {
		return nil
	}
	out := new(DataPlaneClientAuthStatus)
	in.DeepCopyInto(out)
	return out
}

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
func (in *KonnectCloudGatewayDataPlaneGroupConfiguration) DeepCopyInto(out *KonnectCloudGatewayDataPlaneGroupConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayDataPlaneGroupConfiguration.
func (in *KonnectCloudGatewayDataPlaneGroupConfiguration) DeepCopy() *KonnectCloudGatewayDataPlaneGroupConfiguration {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayDataPlaneGroupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectCloudGatewayDataPlaneGroupConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationList) DeepCopyInto(out *KonnectCloudGatewayDataPlaneGroupConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectCloudGatewayDataPlaneGroupConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayDataPlaneGroupConfigurationList.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationList) DeepCopy() *KonnectCloudGatewayDataPlaneGroupConfigurationList {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayDataPlaneGroupConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationSpec) DeepCopyInto(out *KonnectCloudGatewayDataPlaneGroupConfigurationSpec) {
	*out = *in
	if in.DataplaneGroups != nil {
		in, out := &in.DataplaneGroups, &out.DataplaneGroups
		*out = make([]KonnectConfigurationDataPlaneGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIAccess != nil {
		in, out := &in.APIAccess, &out.APIAccess
		*out = new(components.APIAccess)
		**out = **in
	}
	in.ControlPlaneRef.DeepCopyInto(&out.ControlPlaneRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayDataPlaneGroupConfigurationSpec.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationSpec) DeepCopy() *KonnectCloudGatewayDataPlaneGroupConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayDataPlaneGroupConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationStatus) DeepCopyInto(out *KonnectCloudGatewayDataPlaneGroupConfigurationStatus) {
	*out = *in
	out.KonnectEntityStatusWithControlPlaneRef = in.KonnectEntityStatusWithControlPlaneRef
	if in.DataPlaneGroups != nil {
		in, out := &in.DataPlaneGroups, &out.DataPlaneGroups
		*out = make([]KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayDataPlaneGroupConfigurationStatus.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationStatus) DeepCopy() *KonnectCloudGatewayDataPlaneGroupConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayDataPlaneGroupConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup) DeepCopyInto(out *KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup) {
	*out = *in
	if in.PrivateIPAddresses != nil {
		in, out := &in.PrivateIPAddresses, &out.PrivateIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EgressIPAddresses != nil {
		in, out := &in.EgressIPAddresses, &out.EgressIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup.
func (in *KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup) DeepCopy() *KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup)
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
func (in *KonnectCloudGatewayNetworkSpec) DeepCopyInto(out *KonnectCloudGatewayNetworkSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(components.NetworkCreateState)
		**out = **in
	}
	out.KonnectConfiguration = in.KonnectConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectCloudGatewayNetworkSpec.
func (in *KonnectCloudGatewayNetworkSpec) DeepCopy() *KonnectCloudGatewayNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectCloudGatewayNetworkSpec)
	in.DeepCopyInto(out)
	return out
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
func (in *KonnectConfigurationDataPlaneGroup) DeepCopyInto(out *KonnectConfigurationDataPlaneGroup) {
	*out = *in
	in.NetworkRef.DeepCopyInto(&out.NetworkRef)
	in.Autoscale.DeepCopyInto(&out.Autoscale)
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]ConfigurationDataPlaneGroupEnvironmentField, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectConfigurationDataPlaneGroup.
func (in *KonnectConfigurationDataPlaneGroup) DeepCopy() *KonnectConfigurationDataPlaneGroup {
	if in == nil {
		return nil
	}
	out := new(KonnectConfigurationDataPlaneGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectEndpoints) DeepCopyInto(out *KonnectEndpoints) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectEndpoints.
func (in *KonnectEndpoints) DeepCopy() *KonnectEndpoints {
	if in == nil {
		return nil
	}
	out := new(KonnectEndpoints)
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
func (in *KonnectExtension) DeepCopyInto(out *KonnectExtension) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectExtension.
func (in *KonnectExtension) DeepCopy() *KonnectExtension {
	if in == nil {
		return nil
	}
	out := new(KonnectExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectExtension) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectExtensionControlPlane) DeepCopyInto(out *KonnectExtensionControlPlane) {
	*out = *in
	in.ControlPlaneRef.DeepCopyInto(&out.ControlPlaneRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectExtensionControlPlane.
func (in *KonnectExtensionControlPlane) DeepCopy() *KonnectExtensionControlPlane {
	if in == nil {
		return nil
	}
	out := new(KonnectExtensionControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectExtensionControlPlaneStatus) DeepCopyInto(out *KonnectExtensionControlPlaneStatus) {
	*out = *in
	out.Endpoints = in.Endpoints
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectExtensionControlPlaneStatus.
func (in *KonnectExtensionControlPlaneStatus) DeepCopy() *KonnectExtensionControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectExtensionControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectExtensionList) DeepCopyInto(out *KonnectExtensionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectExtensionList.
func (in *KonnectExtensionList) DeepCopy() *KonnectExtensionList {
	if in == nil {
		return nil
	}
	out := new(KonnectExtensionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectExtensionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectExtensionSpec) DeepCopyInto(out *KonnectExtensionSpec) {
	*out = *in
	in.KonnectControlPlane.DeepCopyInto(&out.KonnectControlPlane)
	if in.DataPlaneClientAuth != nil {
		in, out := &in.DataPlaneClientAuth, &out.DataPlaneClientAuth
		*out = new(DataPlaneClientAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.KonnectConfiguration != nil {
		in, out := &in.KonnectConfiguration, &out.KonnectConfiguration
		*out = new(KonnectConfiguration)
		**out = **in
	}
	if in.DataPlaneLabels != nil {
		in, out := &in.DataPlaneLabels, &out.DataPlaneLabels
		*out = make(map[string]DataPlaneLabelValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectExtensionSpec.
func (in *KonnectExtensionSpec) DeepCopy() *KonnectExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectExtensionStatus) DeepCopyInto(out *KonnectExtensionStatus) {
	*out = *in
	if in.DataPlaneRefs != nil {
		in, out := &in.DataPlaneRefs, &out.DataPlaneRefs
		*out = make([]commonv1alpha1.NamespacedRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ControlPlaneRefs != nil {
		in, out := &in.ControlPlaneRefs, &out.ControlPlaneRefs
		*out = make([]commonv1alpha1.NamespacedRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataPlaneClientAuth != nil {
		in, out := &in.DataPlaneClientAuth, &out.DataPlaneClientAuth
		*out = new(DataPlaneClientAuthStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Konnect != nil {
		in, out := &in.Konnect, &out.Konnect
		*out = new(KonnectExtensionControlPlaneStatus)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectExtensionStatus.
func (in *KonnectExtensionStatus) DeepCopy() *KonnectExtensionStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectExtensionStatus)
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
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(KonnectEndpoints)
		**out = **in
	}
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
func (in *NetworkRef) DeepCopyInto(out *NetworkRef) {
	*out = *in
	if in.KonnectID != nil {
		in, out := &in.KonnectID, &out.KonnectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRef.
func (in *NetworkRef) DeepCopy() *NetworkRef {
	if in == nil {
		return nil
	}
	out := new(NetworkRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}
