package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
)

// Code generated by scripts/apitypes-funcs/main.go; DO NOT EDIT.

func (obj *KongKey) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndKeySetRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongKey status.
func (obj *KongKey) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongKey status.
func (obj *KongKey) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongKey status.
func (obj *KongKey) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongKey status.
func (obj *KongKey) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongKey status.
func (obj *KongKey) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongKey Kind name
func (obj KongKey) GetTypeName() string {
	return "KongKey"
}

// GetConditions returns the Status Conditions
func (obj *KongKey) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongKey) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongKey) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongKey) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongKeySet) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongKeySet status.
func (obj *KongKeySet) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongKeySet status.
func (obj *KongKeySet) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongKeySet status.
func (obj *KongKeySet) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongKeySet status.
func (obj *KongKeySet) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongKeySet status.
func (obj *KongKeySet) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongKeySet Kind name
func (obj KongKeySet) GetTypeName() string {
	return "KongKeySet"
}

// GetConditions returns the Status Conditions
func (obj *KongKeySet) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongKeySet) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongKeySet) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongKeySet) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongCredentialBasicAuth) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCredentialBasicAuth status.
func (obj *KongCredentialBasicAuth) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCredentialBasicAuth status.
func (obj *KongCredentialBasicAuth) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCredentialBasicAuth status.
func (obj *KongCredentialBasicAuth) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCredentialBasicAuth status.
func (obj *KongCredentialBasicAuth) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCredentialBasicAuth status.
func (obj *KongCredentialBasicAuth) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCredentialBasicAuth Kind name
func (obj KongCredentialBasicAuth) GetTypeName() string {
	return "KongCredentialBasicAuth"
}

// GetConditions returns the Status Conditions
func (obj *KongCredentialBasicAuth) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCredentialBasicAuth) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCredentialBasicAuth) SetKonnectConsumerIDInStatus(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ConsumerID = id
}

func (obj *KongCredentialBasicAuth) GetConsumerRefName() string {
	return obj.Spec.ConsumerRef.Name
}

func (obj *KongCredentialAPIKey) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCredentialAPIKey status.
func (obj *KongCredentialAPIKey) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCredentialAPIKey status.
func (obj *KongCredentialAPIKey) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCredentialAPIKey status.
func (obj *KongCredentialAPIKey) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCredentialAPIKey status.
func (obj *KongCredentialAPIKey) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCredentialAPIKey status.
func (obj *KongCredentialAPIKey) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCredentialAPIKey Kind name
func (obj KongCredentialAPIKey) GetTypeName() string {
	return "KongCredentialAPIKey"
}

// GetConditions returns the Status Conditions
func (obj *KongCredentialAPIKey) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCredentialAPIKey) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCredentialAPIKey) SetKonnectConsumerIDInStatus(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ConsumerID = id
}

func (obj *KongCredentialAPIKey) GetConsumerRefName() string {
	return obj.Spec.ConsumerRef.Name
}

func (obj *KongCredentialJWT) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCredentialJWT status.
func (obj *KongCredentialJWT) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCredentialJWT status.
func (obj *KongCredentialJWT) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCredentialJWT status.
func (obj *KongCredentialJWT) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCredentialJWT status.
func (obj *KongCredentialJWT) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCredentialJWT status.
func (obj *KongCredentialJWT) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCredentialJWT Kind name
func (obj KongCredentialJWT) GetTypeName() string {
	return "KongCredentialJWT"
}

// GetConditions returns the Status Conditions
func (obj *KongCredentialJWT) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCredentialJWT) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCredentialJWT) SetKonnectConsumerIDInStatus(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ConsumerID = id
}

func (obj *KongCredentialJWT) GetConsumerRefName() string {
	return obj.Spec.ConsumerRef.Name
}

func (obj *KongCredentialACL) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCredentialACL status.
func (obj *KongCredentialACL) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCredentialACL status.
func (obj *KongCredentialACL) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCredentialACL status.
func (obj *KongCredentialACL) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCredentialACL status.
func (obj *KongCredentialACL) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCredentialACL status.
func (obj *KongCredentialACL) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCredentialACL Kind name
func (obj KongCredentialACL) GetTypeName() string {
	return "KongCredentialACL"
}

// GetConditions returns the Status Conditions
func (obj *KongCredentialACL) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCredentialACL) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCredentialACL) SetKonnectConsumerIDInStatus(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ConsumerID = id
}

func (obj *KongCredentialACL) GetConsumerRefName() string {
	return obj.Spec.ConsumerRef.Name
}

func (obj *KongCredentialHMAC) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCredentialHMAC status.
func (obj *KongCredentialHMAC) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCredentialHMAC status.
func (obj *KongCredentialHMAC) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCredentialHMAC status.
func (obj *KongCredentialHMAC) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCredentialHMAC status.
func (obj *KongCredentialHMAC) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCredentialHMAC status.
func (obj *KongCredentialHMAC) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCredentialHMAC Kind name
func (obj KongCredentialHMAC) GetTypeName() string {
	return "KongCredentialHMAC"
}

// GetConditions returns the Status Conditions
func (obj *KongCredentialHMAC) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCredentialHMAC) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCredentialHMAC) SetKonnectConsumerIDInStatus(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ConsumerID = id
}

func (obj *KongCredentialHMAC) GetConsumerRefName() string {
	return obj.Spec.ConsumerRef.Name
}

func (obj *KongCACertificate) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCACertificate status.
func (obj *KongCACertificate) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCACertificate status.
func (obj *KongCACertificate) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCACertificate status.
func (obj *KongCACertificate) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCACertificate status.
func (obj *KongCACertificate) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCACertificate status.
func (obj *KongCACertificate) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCACertificate Kind name
func (obj KongCACertificate) GetTypeName() string {
	return "KongCACertificate"
}

// GetConditions returns the Status Conditions
func (obj *KongCACertificate) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCACertificate) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCACertificate) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongCACertificate) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongCertificate) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongCertificate status.
func (obj *KongCertificate) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongCertificate status.
func (obj *KongCertificate) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongCertificate status.
func (obj *KongCertificate) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongCertificate status.
func (obj *KongCertificate) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongCertificate status.
func (obj *KongCertificate) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongCertificate Kind name
func (obj KongCertificate) GetTypeName() string {
	return "KongCertificate"
}

// GetConditions returns the Status Conditions
func (obj *KongCertificate) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongCertificate) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongCertificate) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongCertificate) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongPluginBinding) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongPluginBinding status.
func (obj *KongPluginBinding) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongPluginBinding status.
func (obj *KongPluginBinding) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongPluginBinding status.
func (obj *KongPluginBinding) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongPluginBinding status.
func (obj *KongPluginBinding) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongPluginBinding status.
func (obj *KongPluginBinding) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongPluginBinding Kind name
func (obj KongPluginBinding) GetTypeName() string {
	return "KongPluginBinding"
}

// GetConditions returns the Status Conditions
func (obj *KongPluginBinding) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongPluginBinding) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongPluginBinding) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	if ref == nil {
		obj.Spec.ControlPlaneRef = commonv1alpha1.ControlPlaneRef{}
		return
	}
	obj.Spec.ControlPlaneRef = *ref
}

func (obj *KongPluginBinding) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return &obj.Spec.ControlPlaneRef
}

func (obj *KongService) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongService status.
func (obj *KongService) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongService status.
func (obj *KongService) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongService status.
func (obj *KongService) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongService status.
func (obj *KongService) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongService status.
func (obj *KongService) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongService Kind name
func (obj KongService) GetTypeName() string {
	return "KongService"
}

// GetConditions returns the Status Conditions
func (obj *KongService) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongService) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongService) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongService) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongRoute) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndServiceRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongRoute status.
func (obj *KongRoute) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongRoute status.
func (obj *KongRoute) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongRoute status.
func (obj *KongRoute) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongRoute status.
func (obj *KongRoute) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongRoute status.
func (obj *KongRoute) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongRoute Kind name
func (obj KongRoute) GetTypeName() string {
	return "KongRoute"
}

// GetConditions returns the Status Conditions
func (obj *KongRoute) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongRoute) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongRoute) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongRoute) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongUpstream) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongUpstream status.
func (obj *KongUpstream) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongUpstream status.
func (obj *KongUpstream) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongUpstream status.
func (obj *KongUpstream) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongUpstream status.
func (obj *KongUpstream) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongUpstream status.
func (obj *KongUpstream) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongUpstream Kind name
func (obj KongUpstream) GetTypeName() string {
	return "KongUpstream"
}

// GetConditions returns the Status Conditions
func (obj *KongUpstream) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongUpstream) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongUpstream) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongUpstream) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongTarget) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndUpstreamRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongTarget status.
func (obj *KongTarget) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongTarget status.
func (obj *KongTarget) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongTarget status.
func (obj *KongTarget) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongTarget status.
func (obj *KongTarget) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongTarget status.
func (obj *KongTarget) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongTarget Kind name
func (obj KongTarget) GetTypeName() string {
	return "KongTarget"
}

// GetConditions returns the Status Conditions
func (obj *KongTarget) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongTarget) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongVault) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongVault status.
func (obj *KongVault) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongVault status.
func (obj *KongVault) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongVault status.
func (obj *KongVault) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongVault status.
func (obj *KongVault) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongVault status.
func (obj *KongVault) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongVault Kind name
func (obj KongVault) GetTypeName() string {
	return "KongVault"
}

// GetConditions returns the Status Conditions
func (obj *KongVault) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongVault) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongVault) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongVault) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}

func (obj *KongSNI) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndCertificateRefs{}
}

// GetKonnectStatus returns the Konnect status contained in the KongSNI status.
func (obj *KongSNI) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongSNI status.
func (obj *KongSNI) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongSNI status.
func (obj *KongSNI) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongSNI status.
func (obj *KongSNI) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongSNI status.
func (obj *KongSNI) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongSNI Kind name
func (obj KongSNI) GetTypeName() string {
	return "KongSNI"
}

// GetConditions returns the Status Conditions
func (obj *KongSNI) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongSNI) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongDataPlaneClientCertificate) initKonnectStatus() {
	obj.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongDataPlaneClientCertificate status.
func (obj *KongDataPlaneClientCertificate) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongDataPlaneClientCertificate status.
func (obj *KongDataPlaneClientCertificate) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongDataPlaneClientCertificate status.
func (obj *KongDataPlaneClientCertificate) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongDataPlaneClientCertificate status.
func (obj *KongDataPlaneClientCertificate) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongDataPlaneClientCertificate status.
func (obj *KongDataPlaneClientCertificate) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongDataPlaneClientCertificate Kind name
func (obj KongDataPlaneClientCertificate) GetTypeName() string {
	return "KongDataPlaneClientCertificate"
}

// GetConditions returns the Status Conditions
func (obj *KongDataPlaneClientCertificate) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *KongDataPlaneClientCertificate) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

func (obj *KongDataPlaneClientCertificate) SetControlPlaneRef(ref *commonv1alpha1.ControlPlaneRef) {
	obj.Spec.ControlPlaneRef = ref
}

func (obj *KongDataPlaneClientCertificate) GetControlPlaneRef() *commonv1alpha1.ControlPlaneRef {
	return obj.Spec.ControlPlaneRef
}
