package values

import "fmt"

type S3 struct {
	Enabled   bool
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
}

type Minio struct {
	Hostname string
}

type RabbitMQ struct {
	Username string
	Password string
}

type MongoDB struct {
	Enabled      bool
	Username     string
	Password     string
	RootPassword string
	Database     string
	Endpoint     string
}

type Postgres struct {
	Enabled  bool
	Username string
	Password string
	Database string
	Endpoint string
}

type Redis struct {
	Password string
}

func configS3(root *Values) error {
	vals := &root.S3

	if err := Ask("Use in-cluster Minio object store?", func() (err error) {
		vals.Enabled, err = ScanYesNo(true)
		return
	}); err != nil {
		return err
	}
	if !vals.Enabled {
		if err := Ask("S3 Endpoint:", func() (err error) {
			vals.Endpoint, err = ScanString("https://s3.example.com")
			return
		}); err != nil {
			return err
		}
		if err := Ask("S3 Bucket:", func() (err error) {
			vals.Bucket, err = ScanString("sylabs")
			return
		}); err != nil {
			return err
		}
		if err := Ask("S3 AccessKey:", func() (err error) {
			vals.AccessKey, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("S3 SecretKey:", func() (err error) {
			vals.SecretKey, err = ScanString("")
			return
		}); err != nil {
			return err
		}
	} else {
		if err := Ask("Minio Hostname:", func() (err error) {
			root.Minio.Hostname, err = ScanString("minio.lvh.me")
			return
		}); err != nil {
			return err
		}
		fmt.Println("S3 Endpoint:")
		fmt.Println("local")
		fmt.Println("S3 Bucket:")
		fmt.Println("sylabs")
		vals.Bucket = "sylabs"
		fmt.Println("Randomly generating AccessKey...")
		vals.AccessKey = randomSecret(32)
		fmt.Println("Randomly generating SecretKey...")
		vals.SecretKey = randomSecret(64)
	}
	return nil
}

func configRabbitMQ(root *Values) error {
	vals := &root.RabbitMQ
	vals.Username = "rabbitmq"
	fmt.Println("Generating random RabbitMQ password...")
	vals.Password = randomSecret(64)
	return nil
}

func configMongoDB(root *Values) error {
	vals := &root.MongoDB
	defaultDatabase := "sylabs"

	if err := Ask("Use in-cluster MongoDB?", func() (err error) {
		vals.Enabled, err = ScanYesNo(true)
		return
	}); err != nil {
		return err
	}

	if !vals.Enabled {
		if err := Ask("MongoDB Endpoint:", func() (err error) {
			vals.Endpoint, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("MongoDB Username:", func() (err error) {
			vals.Username, err = ScanString("")
			return
		}); err != nil {
			return err
		}
		if err := Ask("MongoDB Password:", func() (err error) {
			vals.Password, err = ScanString("")
			return
		}); err != nil {
			return err
		}
	} else {
		vals.Database = defaultDatabase
		fmt.Println("Randomly generating MongoDB root password....")
		vals.RootPassword = randomSecret(64)
		vals.Username = "mongodb"
		fmt.Println("Randomly generating MongoDB password...")
		vals.Password = randomSecret(64)
	}
	return nil
}

func configPostgres(root *Values) error {
	vals := &root.Postgres

	if err := Ask("Use in-cluster PostgreSQL?", func() (err error) {
		vals.Enabled, err = ScanYesNo(true)
		return
	}); err != nil {
		return err
	}

	if !vals.Enabled {
		if err := Ask("PostgreSQL Endpoint:", func() (err error) {
			vals.Endpoint, err = ScanString("")
			return
		}); err != nil {
			return err
		}

		if err := Ask("PostgreSQL Database:", func() (err error) {
			vals.Database, err = ScanString("")
			return
		}); err != nil {
			return err
		}

		if err := Ask("PostgreSQL Username:", func() (err error) {
			vals.Username, err = ScanString("")
			return
		}); err != nil {
			return err
		}

		if err := Ask("PostgreSQL Password:", func() (err error) {
			vals.Password, err = ScanString("")
			return
		}); err != nil {
			return err
		}
	} else {
		vals.Username = "postgres"
		fmt.Println("Generating random Postgres password...")
		vals.Password = randomSecret(64)
		vals.Database = "hydra"
	}

	return nil
}

func configRedis(root *Values) error {
	fmt.Println("Generating random Redis password...")
	root.Redis.Password = randomSecret(64)
	return nil
}

func ConfigExternal(root *Values) error {
	if err := configS3(root); err != nil {
		return err
	}
	if err := configRabbitMQ(root); err != nil {
		return err
	}
	if err := configMongoDB(root); err != nil {
		return err
	}
	if err := configPostgres(root); err != nil {
		return err
	}
	if err := configRedis(root); err != nil {
		return err
	}
	return nil
}
