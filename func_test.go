package genplate

import (
	"testing"
)

func TestDepunct(t *testing.T) {
	candidates := []struct {
		ident        string
		capDepuncted string
		lowDepuncted string
	}{
		{
			ident:        "provider_id",
			capDepuncted: "ProviderID",
			lowDepuncted: "providerID",
		},
		{
			ident:        "app-identity",
			capDepuncted: "AppIdentity",
			lowDepuncted: "appIdentity",
		},
		{
			ident:        "uuid",
			capDepuncted: "UUID",
			lowDepuncted: "uuid",
		},
		{
			ident:        "oauth-client",
			capDepuncted: "OAuthClient",
			lowDepuncted: "oauthClient",
		},
		{
			ident:        "Dyno all",
			capDepuncted: "DynoAll",
			lowDepuncted: "dynoAll",
		},
	}

	for i, c := range candidates {
		depuncted := depunct(c.ident, true)
		if depuncted != c.capDepuncted {
			t.Errorf("%d: wants %v, got %v", i, c.capDepuncted, depuncted)
		}
	}

	for i, c := range candidates {
		depuncted := depunct(c.ident, false)
		if depuncted != c.lowDepuncted {
			t.Errorf("%d: wants %v, got %v", i, c.lowDepuncted, depuncted)
		}
	}
}
