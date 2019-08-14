package values

import "fmt"

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

func ConfigInfrastructure(root *Values) {
	root.Ingress.Enabled = true

	// Not supporting routes at the moment
	root.Route.Enabled = false

	root.ServiceMonitor.Enabled = false

	root.PullCredentials.Name = "sylabs-regcred"
	fmt.Println("Sylabs Registry Username:")
	fmt.Print("[ex: john@doe.com] ")
	fmt.Scanln(&root.PullCredentials.Username)
	fmt.Println("Sylabs Registry Password:")
	fmt.Scanln(&root.PullCredentials.Password)
}
