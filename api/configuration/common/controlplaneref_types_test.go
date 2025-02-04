package common_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	"github.com/kong/kubernetes-configuration/api/configuration/common"
)

func TestControlPlaneRefStringer(t *testing.T) {
	testCases := []struct {
		name     string
		ref      *common.ControlPlaneRef
		expected string
	}{
		{
			name: "unknown type - doesn't panic",
			ref: &common.ControlPlaneRef{
				Type: "notSupportedType",
			},
			expected: "<unknown:notSupportedType>",
		},
		{
			name:     "nil - doesn't panic",
			ref:      nil,
			expected: "<nil>",
		},
		{
			name: "konnectNamespacedRef with no namespace",
			ref: &common.ControlPlaneRef{
				Type: common.ControlPlaneRefKonnectNamespacedRef,
				KonnectNamespacedRef: &common.KonnectNamespacedRef{
					Name: "foo",
				},
			},
			expected: "<konnectNamespacedRef:foo>",
		},
		{
			name: "konnectNamespacedRef with namespace",
			ref: &common.ControlPlaneRef{
				Type: common.ControlPlaneRefKonnectNamespacedRef,
				KonnectNamespacedRef: &common.KonnectNamespacedRef{
					Namespace: "bar",
					Name:      "foo",
				},
			},
			expected: "<konnectNamespacedRef:bar/foo>",
		},
		{
			name: "konnectID without ID - doesn't panic",
			ref: &common.ControlPlaneRef{
				Type: common.ControlPlaneRefKonnectID,
			},
			expected: "<konnectID:nil>",
		},
		{
			name: "konnectID with ID",
			ref: &common.ControlPlaneRef{
				Type:      common.ControlPlaneRefKonnectID,
				KonnectID: lo.ToPtr("foo"),
			},
			expected: "<konnectID:foo>",
		},
		{
			name: "kic",
			ref: &common.ControlPlaneRef{
				Type: common.ControlPlaneRefKIC,
			},
			expected: "<kic>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.ref.String())
		})
	}
}
