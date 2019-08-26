package values

import (
	"io"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Values struct {
	DefaultDomain string
	// Auth Service Structs
	TokenService   TokenService
	ConsentService ConsentService
	Hydra          Hydra

	// Key Service Structs
	KeyService KeyService

	// Cloud Library Structs
	CloudLibraryServer CloudLibraryServer

	// Frontend Structs
	Frontend Frontend

	// Remote Build Structs
	RemoteBuildManager RemoteBuildManager
	RemoteBuildServer  RemoteBuildServer

	// Externally Maintained Dependencies
	S3       S3
	RabbitMQ RabbitMQ
	MongoDB  MongoDB
	Postgres Postgres
	Redis    Redis
	Minio    Minio

	// Infrastructure Related Structs
	Ingress         Ingress
	Route           Route
	ServiceMonitor  ServiceMonitor
	PullCredentials PullCredentials
}

func ConfigValues(root *Values) error {
	if err := Ask("Default Domain:", func() (err error) {
		root.DefaultDomain, err = ScanString("lvh.me")
		return
	}); err != nil {
		return err
	}

	cs := [](func(*Values) error){ConfigInfrastructure, ConfigExternal, ConfigAuthService, ConfigKeyService, ConfigCloudLibrary, ConfigRemoteBuild, ConfigFrontend}
	for _, c := range cs {
		if err := c(root); err != nil {
			return err
		}
	}
	return nil
}

func (v *Values) Configure() (err error) {
	return ConfigValues(v)
}

func (v *Values) Render(w io.Writer) (err error) {
	t, err := template.New("values").Funcs(sprig.TxtFuncMap()).Parse(Template)
	if err != nil {
		panic(err) // bad hardcoded string, panic
	}
	return t.Execute(w, v)
}
