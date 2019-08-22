package values

type Frontend struct {
	URI       string
	RevokeURI string
}

func ConfigFrontend(root *Values) error {
	vals := &root.Frontend
	if err := Ask("Frontend URI:", func() (err error) {
		vals.URI, err = ScanString("https://cloud.lvh.me")
		return
	}); err != nil {
		return err
	}
	defaultRevoke := vals.URI + "/assets/html/auth-revoke.html"
	if err := Ask("Revoke URI:", func() (err error) {
		vals.RevokeURI, err = ScanString(defaultRevoke)
		return
	}); err != nil {
		return err
	}
	return nil
}
