package values

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
