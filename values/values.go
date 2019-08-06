package values

type Values struct {
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

	// Infrastructure Related Structs
	Ingress         Ingress
	Route           Route
	ServiceMonitor  ServiceMonitor
	PullCredentials PullCredentials
}
