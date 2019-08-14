package values

import (
	"bufio"
	"fmt"
	"os"
)

type TokenService struct {
	URI           string
	RSASecretName string
}

type OAuth2 struct {
	ClientID     string
	ClientSecret string
}

type CustomOAuth2 struct {
	Name         string
	ClientID     string
	ClientSecret string
	IssuerURL    string
	Scope        string
}

type ConsentService struct {
	URI   string
	OAuth struct {
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
	URI            string
	ClientSecret   string
	CookieSecret   string
	FrontendSecret string
	ConsentSecret  string
	ConsentURL     string
	LoginURL       string
}

func configConsentService(root *Values) {
	vals := &root.ConsentService
	fmt.Println("ConsentService URI:")
	fmt.Print("[https://auth.lvh.me] ")
	fmt.Scanln(&vals.URI)

	if vals.URI == "" {
		vals.URI = "https://auth.lvh.me"
	}

	var useGoogle, useGitHub, useMicrosoft, useCustom1, useCustom2, useCustom3 bool
	fmt.Println("Configure Google OpenID provider?")
	fmt.Print("[y/N] ")
	fmt.Scanln(&useGoogle)
	if useGoogle {
		fmt.Println("Google OAuth2 ClientID:")
		fmt.Scanln(&vals.OAuth.Google.ClientID)
		fmt.Println("Google OAuth2 ClientSecret:")
		fmt.Scanln(&vals.OAuth.Google.ClientSecret)
	}
	fmt.Println("Configure GitHub OpenID provider?")
	fmt.Print("[y/N] ")
	fmt.Scanln(&useGitHub)
	if useGitHub {
		fmt.Println("GitHub OAuth2 ClientID:")
		fmt.Scanln(&vals.OAuth.GitHub.ClientID)
		fmt.Println("GitHub OAuth2 ClientSecret")
		fmt.Scanln(&vals.OAuth.GitHub.ClientSecret)
	}
	fmt.Println("Configure Microsoft OpenID provider?")
	fmt.Print("[y/N] ")
	fmt.Scanln(&useMicrosoft)
	if useMicrosoft {
		fmt.Println("Microsoft OAuth2 ClientID:")
		fmt.Scanln(&vals.OAuth.Microsoft.ClientID)
		fmt.Println("Microsoft OAuth2 ClientSecret:")
		fmt.Scanln(&vals.OAuth.Microsoft.ClientSecret)
	}
	fmt.Println("Configure a Custom OpenID provider?")
	fmt.Print("[y/N] ")
	fmt.Scanln(&useCustom1)

	parseCustomOAuth := func(c *CustomOAuth2) {
		fmt.Println("Provider Name:")
		fmt.Scanln(&c.Name)
		fmt.Println("Provider ClientID:")
		fmt.Scanln(&c.ClientID)
		fmt.Println("Provider ClientSecret:")
		fmt.Scanln(&c.ClientSecret)
		fmt.Println("Provider IssuerURL:")
		fmt.Scanln(&c.IssuerURL)
		fmt.Println("Provider Scope:")
		scanner := bufio.NewScanner(os.Stdin)
		scope := ""
		for scanner.Scan() {
			scope = scope + " " + scanner.Text()
		}
		c.Scope = scope
	}

	if useCustom1 {
		parseCustomOAuth(&vals.OAuth.Custom1)
		fmt.Println("Configure another Custom OpenID provider?")
		fmt.Print("[y/N] ")
		fmt.Scanln(&useCustom2)
		if useCustom2 {
			parseCustomOAuth(&vals.OAuth.Custom2)
			fmt.Println("Configure another Custom OpenID provider?")
			fmt.Print("[y/N] ")
			fmt.Scanln(&useCustom3)
			if useCustom3 {
				parseCustomOAuth(&vals.OAuth.Custom3)
			}
		}
	}
	fmt.Println("Username to coerce as Administrator:")
	fmt.Scanln(&vals.AdminUser)
}

func configHydra(root *Values) {
	vals := &root.Hydra
	fmt.Println("Hydra URI:")
	fmt.Print("[https://hydra.lvh.me] ")
	fmt.Scanln(&vals.URI)
	if vals.URI == "" {
		vals.URI = "https://hydra.lvh.me"
	}
	fmt.Println("Consent URL:")
	defaultConsentURL := root.ConsentService.URI + "/v1/consent"
	fmt.Printf("[%s] ", defaultConsentURL)
	fmt.Scanln(&vals.ConsentURL)
	if vals.ConsentURL == "" {
		vals.ConsentURL = defaultConsentURL
	}
	fmt.Println("Login URL:")
	defaultLoginURL := root.ConsentService.URI + "/v1/login"
	fmt.Scanln(&vals.LoginURL)
	if vals.LoginURL == "" {
		vals.LoginURL = defaultLoginURL
	}

	fmt.Println("Generating random CookieSecret...")
	vals.CookieSecret = randomSecret(64)
	fmt.Println("Generating random FrontendSecret...")
	vals.FrontendSecret = randomSecret(64)
	fmt.Println("Generating random ClientSecret...")
	vals.ClientSecret = randomSecret(64)
	fmt.Println("Generating random ConsentSecret...")
	vals.ConsentSecret = randomSecret(64)
}

func configTokenService(root *Values) {
	vals := &root.TokenService
	fmt.Println("TokenService URI:")
	defaultURI := "https://token.lvh.me"
	fmt.Printf("[%s] ", defaultURI)
	fmt.Scanln(&vals.URI)
	if vals.URI == "" {
		vals.URI = defaultURI
	}
	defaultRSASecretName := "singularity-enterprise-token"
	fmt.Println("Token Name:")
	fmt.Printf("[%s] ", defaultRSASecretName)
	if vals.RSASecretName == "" {
		vals.RSASecretName = defaultRSASecretName
	}
}

func ConfigAuthService(vals *Values) {
	configConsentService(vals)
	configHydra(vals)
	configTokenService(vals)
}
