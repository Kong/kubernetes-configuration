package crdsvalidation_test

import (
	"testing"

	"github.com/samber/lo"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"

	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation"
)

func TestKonnectNetwork(t *testing.T) {
	t.Run("spec", func(t *testing.T) {
		crdsvalidation.TestCasesGroup[*konnectv1alpha1.KonnectNetwork]{
			{
				Name: "all required fields are set",
				TestObject: &konnectv1alpha1.KonnectNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						CreateNetworkRequest: sdkkonnectcomp.CreateNetworkRequest{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
							// NOTE: this is required as of now because we embded the sdk type in the CRD
							// and we do not have controler over the field being optional or required
							// without any additional machinery in place.
							// The API itself https://docs.konghq.com/konnect/api/cloud-gateways/latest/#/Networks/create-network
							// does not require this field.
							State: lo.ToPtr(sdkkonnectcomp.NetworkCreateStateInitializing),
						},
					},
				},
			},
		}.Run(t)
	})
}
