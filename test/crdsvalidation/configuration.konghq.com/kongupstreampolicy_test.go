package configuration_test

import (
	"testing"

	"github.com/samber/lo"

	configurationv1beta1 "github.com/kong/kubernetes-configuration/api/configuration/v1beta1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation/common"
)

func TestKongUpstreamPolicy(t *testing.T) {
	t.Run("consistent-hashing", func(t *testing.T) {
		common.TestCasesGroup[*configurationv1beta1.KongUpstreamPolicy]{
			{
				Name: "hash on cookie with valid input",
				TestObject: &configurationv1beta1.KongUpstreamPolicy{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1beta1.KongUpstreamPolicySpec{
						Algorithm: lo.ToPtr("consistent-hashing"),
						HashOn: &configurationv1beta1.KongUpstreamHash{
							Cookie:     lo.ToPtr("session-cookie-name"),
							CookiePath: lo.ToPtr("/cookie-path"),
						},
					},
				},
			},
			{
				Name: "hash on cookie requires cookiePath field to be set",
				TestObject: &configurationv1beta1.KongUpstreamPolicy{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1beta1.KongUpstreamPolicySpec{
						Algorithm: lo.ToPtr("consistent-hashing"),
						HashOn: &configurationv1beta1.KongUpstreamHash{
							Cookie: lo.ToPtr("session-cookie-name"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("When spec.hashOn.cookie is set, spec.hashOn.cookiePath is required."),
			},
			{
				Name: "hash on cookiePath requires cookie field to be set",
				TestObject: &configurationv1beta1.KongUpstreamPolicy{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1beta1.KongUpstreamPolicySpec{
						Algorithm: lo.ToPtr("consistent-hashing"),
						HashOn: &configurationv1beta1.KongUpstreamHash{
							CookiePath: lo.ToPtr("/cookie-path"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("When spec.hashOn.cookiePath is set, spec.hashOn.cookie is required."),
			},
			{
				Name: "hash on header with valid input",
				TestObject: &configurationv1beta1.KongUpstreamPolicy{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1beta1.KongUpstreamPolicySpec{
						Algorithm: lo.ToPtr("consistent-hashing"),
						HashOn: &configurationv1beta1.KongUpstreamHash{
							Header: lo.ToPtr("X-Custom-Header"),
						},
					},
				},
			},
			{
				Name: "hash on header, hash on fallback cookie with valid input",
				TestObject: &configurationv1beta1.KongUpstreamPolicy{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1beta1.KongUpstreamPolicySpec{
						Algorithm: lo.ToPtr("consistent-hashing"),
						HashOn: &configurationv1beta1.KongUpstreamHash{
							Header: lo.ToPtr("X-Custom-Header"),
						},
						HashOnFallback: &configurationv1beta1.KongUpstreamHash{
							Cookie:     lo.ToPtr("fallback-cookie"),
							CookiePath: lo.ToPtr("/fallback-cookie-path"),
						},
					},
				},
			},
			{
				// NOTE: Per https://developer.konghq.com/gateway/entities/upstream/#consistent-hashing
				// > The hash_fallback setting is invalid and can’t be used if cookie is the primary hashing mechanism.
				Name: "hash on fallback (cookie) cannot be set when hash on cookie is set",
				TestObject: &configurationv1beta1.KongUpstreamPolicy{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1beta1.KongUpstreamPolicySpec{
						Algorithm: lo.ToPtr("consistent-hashing"),
						HashOn: &configurationv1beta1.KongUpstreamHash{
							Cookie:     lo.ToPtr("cookie"),
							CookiePath: lo.ToPtr("/cookie-path"),
						},
						HashOnFallback: &configurationv1beta1.KongUpstreamHash{
							Cookie:     lo.ToPtr("fallback-cookie"),
							CookiePath: lo.ToPtr("/fallback-cookie-path"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.hashOnFallback must not be set when spec.hashOn.cookie is set."),
			},
		}.Run(t)
	})
}
