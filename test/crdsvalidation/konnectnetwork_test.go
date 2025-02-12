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
		crdsvalidation.TestCasesGroup[*konnectv1alpha1.KonnectCloudGatewayNetwork]{
			{
				Name: "all required fields are set",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
			},
			{
				Name: "spec.name is immutable",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
				Update: func(n *konnectv1alpha1.KonnectCloudGatewayNetwork) {
					n.Spec.Name = "new-name"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Network name is immutable"),
			},
			{
				Name: "spec.cidr_block is immutable",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
				Update: func(n *konnectv1alpha1.KonnectCloudGatewayNetwork) {
					n.Spec.CidrBlock = "10.0.0.2/24"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Network CIDR block is immutable"),
			},
			{
				Name: "spec.region is immutable",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
				Update: func(n *konnectv1alpha1.KonnectCloudGatewayNetwork) {
					n.Spec.Region = "us-east"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Network region is immutable"),
			},
			{
				Name: "spec.region is immutable",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
				Update: func(n *konnectv1alpha1.KonnectCloudGatewayNetwork) {
					n.Spec.CloudGatewayProviderAccountID = "id-new-1234"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Network cloud gateway provider account ID is immutable"),
			},
			{
				Name: "spec.region is immutable",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
				Update: func(n *konnectv1alpha1.KonnectCloudGatewayNetwork) {
					n.Spec.AvailabilityZones = []string{
						"us-west-1b",
						"us-west-1c",
					}
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Network availability zones are immutable"),
			},
			{
				Name: "spec.state is immutable",
				TestObject: &konnectv1alpha1.KonnectCloudGatewayNetwork{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.KonnectCloudGatewayNetworkSpec{
						KonnectConfiguration: konnectv1alpha1.KonnectConfiguration{
							APIAuthConfigurationRef: konnectv1alpha1.KonnectAPIAuthConfigurationRef{
								Name: "test-konnect-api-auth-configuration",
							},
						},
						KonnectCloudGatewayNetworkAPISpec: konnectv1alpha1.KonnectCloudGatewayNetworkAPISpec{
							Name:   "test-network",
							Region: "us-west",
							AvailabilityZones: []string{
								"us-west-1a",
								"us-west-1b",
							},
							CidrBlock:                     "10.0.0.1/24",
							CloudGatewayProviderAccountID: "test-cloud-gateway-provider-account-id",
						},
					},
				},
				Update: func(n *konnectv1alpha1.KonnectCloudGatewayNetwork) {
					n.Spec.State = lo.ToPtr(sdkkonnectcomp.NetworkCreateStateOffline)
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Network state is immutable"),
			},
		}.Run(t)
	})
}
