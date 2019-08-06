package values

type Ingress struct {
	Enabled bool
}

type Route struct {
	Enabled bool
}

type ServiceMonitor struct {
	Enabled bool
}

type PullCredentials struct {
	Name     string
	Username string
	Password string
}
