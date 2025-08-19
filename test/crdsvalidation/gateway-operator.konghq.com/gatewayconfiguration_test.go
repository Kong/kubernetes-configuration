package crdsvalidation_test

import (
	"testing"

	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	operatorv1beta1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v1beta1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation/common"
)

func TestGatewayConfiguration(t *testing.T) {
	t.Run("extensions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv1beta1.GatewayConfiguration]{
			{
				Name: "no extensions",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec:       operatorv1beta1.GatewayConfigurationSpec{},
				},
			},
			{
				Name: "valid konnectExtension at the gatewayConfiguration level",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						Extensions: []commonv1alpha1.ExtensionRef{
							{
								Group: "konnect.konghq.com",
								Kind:  "KonnectExtension",
								NamespacedRef: commonv1alpha1.NamespacedRef{
									Name: "my-konnect-extension",
								},
							},
						},
					},
				},
			},
			{
				Name: "invalid konnectExtension",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						Extensions: []commonv1alpha1.ExtensionRef{
							{
								Group: "wrong.konghq.com",
								Kind:  "wrongExtension",
								NamespacedRef: commonv1alpha1.NamespacedRef{
									Name: "my-konnect-extension",
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Extension not allowed for GatewayConfiguration"),
			},
			{
				Name: "konnectExtension at the DataPlane level",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Extensions: []commonv1alpha1.ExtensionRef{
								{
									Group: "konnect.konghq.com",
									Kind:  "KonnectExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-konnect-extension",
									},
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("KonnectExtension must be set at the Gateway level"),
			},
			{
				Name: "konnectExtension at the ControlPlane level",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ControlPlaneOptions: &operatorv1beta1.ControlPlaneOptions{
							Extensions: []commonv1alpha1.ExtensionRef{
								{
									Group: "konnect.konghq.com",
									Kind:  "KonnectExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-konnect-extension",
									},
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("KonnectExtension must be set at the Gateway level"),
			},
		}.Run(t)
	})

	t.Run("DataPlaneOptions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv1beta1.GatewayConfiguration]{
			{
				Name: "no DataPlaneOptions",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: nil,
					},
				},
			},
			{
				Name: "specifying resources.PodDisruptionBudget",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Deployment: operatorv1beta1.DataPlaneDeploymentOptions{
								DeploymentOptions: operatorv1beta1.DeploymentOptions{
									Replicas: lo.ToPtr(int32(4)),
								},
							},
							Resources: &operatorv1beta1.GatewayConfigDataPlaneResources{
								PodDisruptionBudget: &operatorv1beta1.PodDisruptionBudget{
									Spec: operatorv1beta1.PodDisruptionBudgetSpec{
										MinAvailable:               lo.ToPtr(intstr.FromInt(1)),
										UnhealthyPodEvictionPolicy: lo.ToPtr(policyv1.IfHealthyBudget),
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "specifying resources.PodDisruptionBudget can only specify onf of maxUnavailable and minAvailable",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Deployment: operatorv1beta1.DataPlaneDeploymentOptions{
								DeploymentOptions: operatorv1beta1.DeploymentOptions{
									Replicas: lo.ToPtr(int32(4)),
								},
							},
							Resources: &operatorv1beta1.GatewayConfigDataPlaneResources{
								PodDisruptionBudget: &operatorv1beta1.PodDisruptionBudget{
									Spec: operatorv1beta1.PodDisruptionBudgetSpec{
										MinAvailable:   lo.ToPtr(intstr.FromInt(1)),
										MaxUnavailable: lo.ToPtr(intstr.FromInt(1)),
									},
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("You can specify only one of maxUnavailable and minAvailable in a single PodDisruptionBudgetSpec."),
			},
		}.Run(t)
	})

	t.Run("ControlPlaneOptions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv1beta1.GatewayConfiguration]{
			{
				Name: "no ControlPlaneOptions",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ControlPlaneOptions: nil,
					},
				},
			},
			{
				Name: "specifying watch namespaces, type=all",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ControlPlaneOptions: &operatorv1beta1.ControlPlaneOptions{
							WatchNamespaces: &operatorv1beta1.WatchNamespaces{
								Type: operatorv1beta1.WatchNamespacesTypeAll,
							},
						},
					},
				},
			},
			{
				Name: "specifying watch namespaces, type=own",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ControlPlaneOptions: &operatorv1beta1.ControlPlaneOptions{
							WatchNamespaces: &operatorv1beta1.WatchNamespaces{
								Type: operatorv1beta1.WatchNamespacesTypeOwn,
							},
						},
					},
				},
			},
			{
				Name: "specifying watch namespaces, type=list",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ControlPlaneOptions: &operatorv1beta1.ControlPlaneOptions{
							WatchNamespaces: &operatorv1beta1.WatchNamespaces{
								Type: operatorv1beta1.WatchNamespacesTypeList,
								List: []string{
									"namespace1",
									"namespace2",
								},
							},
						},
					},
				},
			},
		}.Run(t)
	})

	t.Run("ListenersOptions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv1beta1.GatewayConfiguration]{
			{
				Name: "specify nodeport for listeners with 'NodePort' dataplane ingress service",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Network: operatorv1beta1.GatewayConfigDataPlaneNetworkOptions{
								Services: &operatorv1beta1.GatewayConfigDataPlaneServices{
									Ingress: &operatorv1beta1.GatewayConfigServiceOptions{
										ServiceOptions: operatorv1beta1.ServiceOptions{
											Type: corev1.ServiceTypeNodePort,
										},
									},
								},
							},
						},
						ListenersOptions: []operatorv1beta1.GatewayConfigurationListenerOptions{
							{
								Name:     "http",
								NodePort: int32(30080),
							},
						},
					},
				},
			},
			{
				Name: "nodePort out of range",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Network: operatorv1beta1.GatewayConfigDataPlaneNetworkOptions{
								Services: &operatorv1beta1.GatewayConfigDataPlaneServices{
									Ingress: &operatorv1beta1.GatewayConfigServiceOptions{
										ServiceOptions: operatorv1beta1.ServiceOptions{
											Type: corev1.ServiceTypeNodePort,
										},
									},
								},
							},
						},
						ListenersOptions: []operatorv1beta1.GatewayConfigurationListenerOptions{
							{
								Name:     "http",
								NodePort: int32(0),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.listenersOptions[0].nodePort in body should be greater than or equal to 1"),
			},
			{
				Name: "Cannot specify nodeport for listeners with 'ClusterIP' dataplane ingress service",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Network: operatorv1beta1.GatewayConfigDataPlaneNetworkOptions{
								Services: &operatorv1beta1.GatewayConfigDataPlaneServices{
									Ingress: &operatorv1beta1.GatewayConfigServiceOptions{
										ServiceOptions: operatorv1beta1.ServiceOptions{
											Type: corev1.ServiceTypeClusterIP,
										},
									},
								},
							},
						},
						ListenersOptions: []operatorv1beta1.GatewayConfigurationListenerOptions{
							{
								Name:     "http",
								NodePort: int32(30080),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Can only specify listener's NodePort when the type of service for dataplane to receive ingress traffic ('spec.dataPlaneOptions.network.services.ingress') is NodePort or LoadBalancer"),
			},
			{
				Name: "Name must be unique in listener options",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ListenersOptions: []operatorv1beta1.GatewayConfigurationListenerOptions{
							{
								Name:     "http",
								NodePort: int32(30080),
							},
							{
								Name:     "http",
								NodePort: int32(30081),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Listener name must be unique within the Gateway"),
			},
			{
				Name: "Nodeport must be unique in listener options",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ListenersOptions: []operatorv1beta1.GatewayConfigurationListenerOptions{
							{
								Name:     "http",
								NodePort: int32(30080),
							},
							{
								Name:     "http-1",
								NodePort: int32(30080),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Nodeport must be unique within the Gateway if specified"),
			},
		}.Run(t)
	})
}
