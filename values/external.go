package values

type S3 struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
}

type RabbitMQ struct {
	Username string
	Password string
}

type MongoDB struct {
	Username     string
	Password     string
	RootPassword string
	Database     string
	Endpoint     string
}

type Postgres struct {
	Username string
	Password string
	Database string
}

type Redis struct {
	Password string
}
