package values

type KeyService struct {
	URI string
}

func ConfigKeyService(root *Values) error {
	vals := &root.KeyService
	defaultKeys := "https://keys.lvh.me"
	return Ask("KeyService URI:", func() (err error) {
		vals.URI, err = ScanString(defaultKeys)
		return
	})
}
