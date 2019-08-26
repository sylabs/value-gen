package values

import (
	"fmt"
)

type TokenService struct {
	Hostname      string
	RSASecretName string
}

type OAuth2 struct {
	Enabled      bool
	ClientID     string
	ClientSecret string
}

type CustomOAuth2 struct {
	Enabled      bool
	Name         string
	ClientID     string
	ClientSecret string
	IssuerURL    string
	Scope        string
}

type ConsentService struct {
	Hostname string
	OAuth    struct {
		Google    OAuth2
		GitHub    OAuth2
		Microsoft OAuth2
		Custom1   CustomOAuth2
		Custom2   CustomOAuth2
		Custom3   CustomOAuth2
	}
	AdminUser string
}

type Hydra struct {
	Hostname       string
	ClientSecret   string
	CookieSecret   string
	FrontendSecret string
	ConsentSecret  string
}

func configConsentService(root *Values) error {
	vals := &root.ConsentService
	if err := Ask("ConsentService Hostname:", func() (err error) {
		vals.Hostname, err = ScanString("auth.lvh.me")
		return
	}); err != nil {
		return err
	}

	if err := Ask("Configure Google OpenID provider?", func() (err error) {
		vals.OAuth.Google.Enabled, err = ScanYesNo(false)
		return
	}); err != nil {
		return err
	}

	if vals.OAuth.Google.Enabled {
		if err := Ask("Google OAauth2 ClientID:", func() (err error) {
			vals.OAuth.Google.ClientID, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("Google OAuth2 ClientSecret:", func() (err error) {
			vals.OAuth.Google.ClientSecret, err = ScanString("")
			return
		}); err != nil {
			return err
		}
	}

	if err := Ask("Configure GitHub OpenID provider?", func() (err error) {
		vals.OAuth.GitHub.Enabled, err = ScanYesNo(false)
		return
	}); err != nil {
		return err
	}
	if vals.OAuth.GitHub.Enabled {
		if err := Ask("GitHub OAuth2 ClientID:", func() (err error) {
			vals.OAuth.GitHub.ClientID, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("GitHub OAuth2 ClientSecret:", func() (err error) {
			vals.OAuth.GitHub.ClientSecret, err = ScanString("")
			return
		}); err != nil {
			return err
		}
	}

	if err := Ask("Configure Microsoft OpenID provider?", func() (err error) {
		vals.OAuth.Microsoft.Enabled, err = ScanYesNo(false)
		return
	}); err != nil {
		return err
	}

	if vals.OAuth.Microsoft.Enabled {
		if err := Ask("Microsoft OAuth2 ClientID:", func() (err error) {
			vals.OAuth.Microsoft.ClientID, err = ScanString("")
			return
		}); err != nil {
			return err
		}

		if err := Ask("Microsoft OAuth2 ClientSecret:", func() (err error) {
			vals.OAuth.Microsoft.ClientSecret, err = ScanString("")
			return
		}); err != nil {
			return err
		}
	}

	parseCustomOAuth := func(c *CustomOAuth2) error {
		if err := Ask("Provider Name:", func() (err error) {
			c.Name, err = ScanString("")
			return
		}); err != nil {
			return err
		}

		if err := Ask("Provider ClientID:", func() (err error) {
			c.ClientID, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("Provider ClientSecret:", func() (err error) {
			c.ClientSecret, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("Provider ClientSecret:", func() (err error) {
			c.IssuerURL, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("Provider Scope:", func() (err error) {
			c.Scope, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		return nil
	}

	if err := Ask("Configure a Custom OpenID provider?", func() (err error) {
		vals.OAuth.Custom1.Enabled, err = ScanYesNo(false)
		return
	}); err != nil {
		return err
	}

	if vals.OAuth.Custom1.Enabled {
		if err := parseCustomOAuth(&vals.OAuth.Custom1); err != nil {
			return err
		}

		if err := Ask("Configure a second Custom OpenID provider?", func() (err error) {
			vals.OAuth.Custom2.Enabled, err = ScanYesNo(false)
			return
		}); err != nil {
			return err
		}
		if vals.OAuth.Custom2.Enabled {
			if err := parseCustomOAuth(&vals.OAuth.Custom2); err != nil {
				return err
			}

			if err := Ask("Configure a third Custom OpenID provider?", func() (err error) {
				vals.OAuth.Custom3.Enabled, err = ScanYesNo(false)
				return
			}); err != nil {
				return err
			}
			if vals.OAuth.Custom3.Enabled {
				if err := parseCustomOAuth(&vals.OAuth.Custom3); err != nil {
					return err
				}
			}
		}
	}

	if err := Ask("Username to coerce as Admin:", func() (err error) {
		vals.AdminUser, err = ScanString("admin")
		return
	}); err != nil {
		return err
	}
	return nil
}

func configHydra(root *Values) error {
	vals := &root.Hydra

	if err := Ask("Hydra URI:", func() (err error) {
		vals.Hostname, err = ScanString("hydra.lvh.me")
		return
	}); err != nil {
		return err
	}

	fmt.Println("Generating random CookieSecret...")
	vals.CookieSecret = randomSecret(64)
	fmt.Println("Generating random FrontendSecret...")
	vals.FrontendSecret = randomSecret(64)
	fmt.Println("Generating random ClientSecret...")
	vals.ClientSecret = randomSecret(64)
	fmt.Println("Generating random ConsentSecret...")
	vals.ConsentSecret = randomSecret(64)

	return nil
}

func configTokenService(root *Values) error {
	vals := &root.TokenService

	if err := Ask("TokenService URI:", func() (err error) {
		vals.Hostname, err = ScanString("token.lvh.me")
		return
	}); err != nil {
		return err
	}

	defaultRSASecretName := "singularity-enterprise-token"
	vals.RSASecretName = defaultRSASecretName

	return nil
}

func ConfigAuthService(vals *Values) error {
	if err := configConsentService(vals); err != nil {
		return err
	}
	if err := configHydra(vals); err != nil {
		return err
	}
	if err := configTokenService(vals); err != nil {
		return err
	}
	return nil
}
