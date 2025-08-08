package v1alpha1_test

import (
	"reflect"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/v2/api/common/v1alpha1"
	v1alpha1 "github.com/kong/kubernetes-configuration/v2/api/konnect/v1alpha1"
	konnectv1alpha2 "github.com/kong/kubernetes-configuration/v2/api/konnect/v1alpha2"
)

func TestKonnectGatewayControlPlane_ConvertTo(t *testing.T) {
	name := "test-name"
	desc := "desc"
	clusterType := sdkkonnectcomp.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane
	authType := sdkkonnectcomp.AuthTypePkiClientCerts
	cloudGateway := true
	proxyUrls := []sdkkonnectcomp.ProxyURL{
		{Host: "host1", Port: 8080, Protocol: "http"},
		{Host: "host2", Port: 8443, Protocol: "https"},
	}
	labels := map[string]string{"foo": "bar"}
	sourceOrigin := commonv1alpha1.EntitySourceOrigin
	sourceMirror := commonv1alpha1.EntitySourceMirror
	members := []corev1.LocalObjectReference{{Name: "member1"}, {Name: "member2"}}
	konnectConfig := konnectv1alpha2.KonnectConfiguration{}

	cases := []struct {
		name             string
		spec             v1alpha1.KonnectGatewayControlPlaneSpec
		mirror           *v1alpha1.MirrorSpec
		expectsCreateReq bool
	}{
		{
			name: "Origin with all fields",
			spec: v1alpha1.KonnectGatewayControlPlaneSpec{
				CreateControlPlaneRequest: v1alpha1.CreateControlPlaneRequest{
					Name:         lo.ToPtr(name),
					Description:  lo.ToPtr(desc),
					ClusterType:  lo.ToPtr(clusterType),
					AuthType:     lo.ToPtr(authType),
					CloudGateway: lo.ToPtr(cloudGateway),
					ProxyUrls:    proxyUrls,
					Labels:       labels,
				},
				Source:               lo.ToPtr(sourceOrigin),
				Members:              members,
				KonnectConfiguration: konnectConfig,
			},
			mirror:           nil,
			expectsCreateReq: true,
		},
		{
			name: "Mirror with MirrorSpec",
			spec: v1alpha1.KonnectGatewayControlPlaneSpec{
				Source: lo.ToPtr(sourceMirror),
			},
			mirror:           &v1alpha1.MirrorSpec{Konnect: v1alpha1.MirrorKonnect{ID: commonv1alpha1.KonnectIDType("mirror-id")}},
			expectsCreateReq: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			obj := &v1alpha1.KonnectGatewayControlPlane{
				Spec: tc.spec,
			}
			obj.Spec.Mirror = tc.mirror
			dst := &konnectv1alpha2.KonnectGatewayControlPlane{}
			err := obj.ConvertTo(dst)
			assert.NoError(t, err)
			if tc.expectsCreateReq {
				assert.NotNil(t, dst.Spec.CreateControlPlaneRequest)
				assert.Equal(t, lo.FromPtr(tc.spec.Name), dst.Spec.CreateControlPlaneRequest.Name)
				assert.Equal(t, tc.spec.Description, dst.Spec.CreateControlPlaneRequest.Description)
				assert.Equal(t, tc.spec.ClusterType, dst.Spec.CreateControlPlaneRequest.ClusterType)
				assert.Equal(t, tc.spec.AuthType, dst.Spec.CreateControlPlaneRequest.AuthType)
				assert.Equal(t, tc.spec.CloudGateway, dst.Spec.CreateControlPlaneRequest.CloudGateway)
				assert.Equal(t, tc.spec.ProxyUrls, dst.Spec.CreateControlPlaneRequest.ProxyUrls)
				assert.Equal(t, tc.spec.Labels, dst.Spec.CreateControlPlaneRequest.Labels)
			} else {
				assert.Nil(t, dst.Spec.CreateControlPlaneRequest)
			}
			if tc.mirror != nil {
				assert.NotNil(t, dst.Spec.Mirror)
				assert.Equal(t, tc.mirror.Konnect.ID, dst.Spec.Mirror.Konnect.ID)
			} else {
				assert.Nil(t, dst.Spec.Mirror)
			}
			assert.Equal(t, tc.spec.Source, dst.Spec.Source)
			assert.Equal(t, tc.spec.Members, dst.Spec.Members)
			assert.Equal(t, tc.spec.KonnectConfiguration, dst.Spec.KonnectConfiguration)
		})
	}
}

func TestKonnectGatewayControlPlane_ConvertFrom(t *testing.T) {
	name := "test-name"
	desc := "desc"
	clusterType := sdkkonnectcomp.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane
	authType := sdkkonnectcomp.AuthTypePkiClientCerts
	cloudGateway := true
	proxyUrls := []sdkkonnectcomp.ProxyURL{
		{Host: "host1", Port: 8080, Protocol: "http"},
		{Host: "host2", Port: 8443, Protocol: "https"},
	}
	labels := map[string]string{"foo": "bar"}
	source := commonv1alpha1.EntitySourceOrigin
	members := []corev1.LocalObjectReference{{Name: "member1"}, {Name: "member2"}}
	konnectConfig := konnectv1alpha2.KonnectConfiguration{}

	cases := []struct {
		name             string
		src              konnectv1alpha2.KonnectGatewayControlPlaneSpec
		mirror           *konnectv1alpha2.MirrorSpec
		expectsCreateReq bool
	}{
		{
			name: "With CreateControlPlaneRequest and Mirror",
			src: konnectv1alpha2.KonnectGatewayControlPlaneSpec{
				CreateControlPlaneRequest: &sdkkonnectcomp.CreateControlPlaneRequest{
					Name:         name,
					Description:  &desc,
					ClusterType:  &clusterType,
					AuthType:     &authType,
					CloudGateway: &cloudGateway,
					ProxyUrls:    proxyUrls,
					Labels:       labels,
				},
				Source:               lo.ToPtr(source),
				Members:              members,
				KonnectConfiguration: konnectConfig,
			},
			mirror:           &konnectv1alpha2.MirrorSpec{Konnect: konnectv1alpha2.MirrorKonnect{ID: commonv1alpha1.KonnectIDType("mirror-id")}},
			expectsCreateReq: true,
		},
		{
			name: "No CreateControlPlaneRequest, no Mirror",
			src: konnectv1alpha2.KonnectGatewayControlPlaneSpec{
				Source: lo.ToPtr(source),
			},
			mirror:           nil,
			expectsCreateReq: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			obj := &v1alpha1.KonnectGatewayControlPlane{}
			src := &konnectv1alpha2.KonnectGatewayControlPlane{
				Spec: tc.src,
			}
			src.Spec.Mirror = tc.mirror
			require.NoError(t, obj.ConvertFrom(src))
			if tc.expectsCreateReq {
				assert.NotNil(t, obj.Spec.CreateControlPlaneRequest)
				assert.Equal(t, lo.ToPtr(tc.src.CreateControlPlaneRequest.Name), obj.Spec.Name)
				assert.Equal(t, tc.src.CreateControlPlaneRequest.Description, obj.Spec.Description)
				assert.Equal(t, tc.src.CreateControlPlaneRequest.ClusterType, obj.Spec.ClusterType)
				assert.Equal(t, tc.src.CreateControlPlaneRequest.AuthType, obj.Spec.AuthType)
				assert.Equal(t, tc.src.CreateControlPlaneRequest.CloudGateway, obj.Spec.CloudGateway)
				assert.Equal(t, tc.src.CreateControlPlaneRequest.ProxyUrls, obj.Spec.ProxyUrls)
				assert.Equal(t, tc.src.CreateControlPlaneRequest.Labels, obj.Spec.Labels)
			} else {
				// We are making sure that if the CreateControlPlaneRequest is not set, it defaults to an empty struct.
				assert.True(t, reflect.DeepEqual(obj.Spec.CreateControlPlaneRequest, v1alpha1.CreateControlPlaneRequest{}))
			}
			if tc.mirror != nil {
				assert.NotNil(t, obj.Spec.Mirror)
				assert.Equal(t, tc.mirror.Konnect.ID, obj.Spec.Mirror.Konnect.ID)
			} else {
				assert.Nil(t, obj.Spec.Mirror)
			}
			assert.Equal(t, tc.src.Source, obj.Spec.Source)
			assert.Equal(t, tc.src.Members, obj.Spec.Members)
			assert.Equal(t, tc.src.KonnectConfiguration, obj.Spec.KonnectConfiguration)
		})
	}
}
