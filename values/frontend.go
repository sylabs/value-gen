package values

type Frontend struct {
	Hostname  string
	RevokeURI string
}

func ConfigFrontend(root *Values) error {
	vals := &root.Frontend
	if err := Ask("Frontend Hostname:", func() (err error) {
		vals.Hostname, err = ScanString("cloud.lvh.me")
		return
	}); err != nil {
		return err
	}
	defaultRevoke := "https://" + vals.Hostname + "/assets/html/auth-revoke.html"
	if err := Ask("Revoke URI:", func() (err error) {
		vals.RevokeURI, err = ScanString(defaultRevoke)
		return
	}); err != nil {
		return err
	}
	return nil
}
