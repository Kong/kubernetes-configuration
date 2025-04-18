package v1

// ConditionType is a type of condition associated with an object.
// This type should be used with the object's Status.Conditions field.
// +apireference:kgo:include
// +apireference:kic:include
type ConditionType string

// ConditionReason defines the set of reasons that explain why a particular
// condition type has been raised.
// +apireference:kgo:include
// +apireference:kic:include
type ConditionReason string

const (
	// ConditionProgrammed indicates whether the controller has generated Kong configuration
	// and has successfully applied it to Kong.
	//
	// Resources that support this condition are:
	//
	// * KongPlugin
	// * KongClusterPlugin
	// * KongConsumer
	// * KongConsumerGroup
	//
	// It is a positive-polarity summary condition, and so should always be
	// present on the resource with ObservedGeneration set.
	//
	// It should be set to Unknown if the controller performs updates to the
	// status before it has all the information it needs to be able to determine
	// if the condition is true.
	//
	// Possible reasons for this condition to be True are:
	//
	// * "Programmed"
	//
	// Possible reasons for this condition to be False are:
	//
	// * "Invalid"
	// * "Pending"
	//
	// Possible reasons for this condition to be Unknown are:
	//
	// * "Pending".
	//
	ConditionProgrammed ConditionType = "Programmed"

	// ReasonProgrammed is used with the ConditionProgrammed condition when the condition is
	// true.
	ReasonProgrammed ConditionReason = "Programmed"

	// ReasonInvalid is used with the ConditionProgrammed condition when the object fails to be
	// translated into Kong configuration or when Kong rejects the configuration.
	ReasonInvalid ConditionReason = "Invalid"

	// ReasonPending is used with the ConditionProgrammed when the status is "Unknown".
	ReasonPending ConditionReason = "Pending"
)

const (
	// ConditionKongConsumerCredentialSecretRefsValid is the condition type used on KongConsumer
	// to indicate whether the credential secret references are valid.
	ConditionKongConsumerCredentialSecretRefsValid = "CredentialSecretRefsValid"

	// ReasonKongConsumerCredentialSecretRefsValid is the condition reason used on KongConsumer
	// with CredentialSecretRefsValid condition when the condition is true.
	ReasonKongConsumerCredentialSecretRefsValid = "Valid"

	// ReasonKongConsumerCredentialSecretRefInvalid is the condition reason used on KongConsumer
	// with CredentialSecretRefsValid condition when the condition is false.
	// This can happen when the secret reference is invalid or the secret does not exist.
	ReasonKongConsumerCredentialSecretRefInvalid = "Invalid"
)
