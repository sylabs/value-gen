package values

type KeyService struct {
	Hostname string
}

func ConfigKeyService(root *Values) error {
	vals := &root.KeyService
	defaultKeys := "keys." + root.DefaultDomain
	return Ask("KeyService Hostname:", func() (err error) {
		vals.Hostname, err = ScanString(defaultKeys)
		return
	})
}
