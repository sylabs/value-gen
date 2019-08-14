package values

import "fmt"

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

func configS3(root *Values) {
	vals := &root.S3
	var useMinio = true
	fmt.Println("Use in-cluster Minio object store?")
	fmt.Print("[Y/n] ")
	fmt.Scanln(&useMinio)
	if !useMinio {
		fmt.Println("S3 Endpoint:")
		fmt.Print("[ex: https://s3.us-east-1.amazon.com] ")
		fmt.Scanln(&vals.Endpoint)
		fmt.Println("S3 Bucket:")
		fmt.Print("[ex: my-bucket] ")
		fmt.Scanln(&vals.Bucket)
		fmt.Println("S3 AccessKey:")
		fmt.Scanln(&vals.AccessKey)
		fmt.Println("S3 SecretKey:")
		fmt.Scanln(&vals.SecretKey)
	} else {
		fmt.Println("S3 Endpoint:")
		fmt.Println("local")
		fmt.Println("S3 Bucket:")
		fmt.Println("sylabs")
		vals.Bucket = "sylabs"
		fmt.Println("Randomly generating AccessKey...")
		vals.AccessKey = randomSecret(32)
		vals.SecretKey = randomSecret(64)
	}
}

func configRabbitMQ(root *Values) {
	vals := &root.RabbitMQ
	vals.Username = "rabbitmq"
	fmt.Println("Generating random RabbitMQ password...")
	vals.Password = randomSecret(64)
}

func configMongoDB(root *Values) {
	vals := &root.MongoDB
	useInternal := true
	defaultDatabase := "sylabs"

	fmt.Println("Use in-cluster MongoDB?")
	fmt.Print("[Y/n] ")
	fmt.Scanln(&useInternal)

	if !useInternal {
		fmt.Println("MongoDB Endpoint:")
		fmt.Print("[ex: https://example.atlas.mongodb.com] ")
		fmt.Scanln(&vals.Endpoint)
		fmt.Println("MongoDB Username:")
		fmt.Scanln(&vals.Username)
		fmt.Println("MongoDB Password:")
		fmt.Scanln(&vals.Password)
	} else {
		vals.Database = defaultDatabase
		vals.RootPassword = randomSecret(64)
		vals.Username = "mongodb"
		vals.Password = randomSecret(64)
	}
}

func configPostgres(root *Values) {
	vals := &root.Postgres
	vals.Username = "postgres"
	fmt.Println("Generating random Postgres password...")
	vals.Password = randomSecret(64)
	vals.Database = "hydra"
}

func configRedis(root *Values) {
	fmt.Println("Generating random Redis password...")
	root.Redis.Password = randomSecret(64)
}

func ConfigExternal(root *Values) {
	configS3(root)
	configRabbitMQ(root)
	configMongoDB(root)
	configPostgres(root)
	configRedis(root)
}
