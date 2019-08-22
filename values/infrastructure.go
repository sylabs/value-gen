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

func ConfigInfrastructure(root *Values) error {
	root.Ingress.Enabled = true

	// Not supporting routes at the moment
	root.Route.Enabled = false

	root.ServiceMonitor.Enabled = false

	root.PullCredentials.Name = "sylabs-regcred"

	if err := Ask("Sylabs Registry Username (ex: joe@example.com):", func() (err error) {
		root.PullCredentials.Username, err = ScanString("")
		return
	}); err != nil {
		return err
	}
	if err := Ask("Sylabs Registry Password:", func() (err error) {
		root.PullCredentials.Password, err = ScanString("")
		return
	}); err != nil {
		return err
	}
	return nil
}
