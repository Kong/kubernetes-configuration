package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	configurationv1 "github.com/kong/kubernetes-configuration/api/configuration/v1"
)

func TestKongConsumer(t *testing.T) {
	c := &configurationv1.KongConsumer{}

	require.Nil(t, c.GetKonnectStatus())
	require.Empty(t, c.GetKonnectStatus().GetKonnectID())
	require.Empty(t, c.GetKonnectStatus().GetOrgID())
	require.Empty(t, c.GetKonnectStatus().GetServerURL())

	require.Equal(t, "", c.GetControlPlaneID())
	c.SetControlPlaneID("123")
	require.Equal(t, "123", c.GetControlPlaneID())
	require.Equal(t, "123", c.Status.Konnect.ControlPlaneID)
}