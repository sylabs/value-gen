package values

type CloudLibraryServer struct {
	URI string
}

func ConfigCloudLibrary(root *Values) error {
	vals := &root.CloudLibraryServer
	defaultURI := "https://library.lvh.me"
	return Ask("CloudLibraryServer URI:", func() (err error) {
		vals.URI, err = ScanString(defaultURI)
		return
	})
}
