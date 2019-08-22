package values

type RemoteBuildServer struct {
	URI string
}

type RemoteBuildManager struct {
	URI string
}

func configRemoteBuildServer(root *Values) error {
	vals := &root.RemoteBuildServer
	defaultServer := "https://build.lvh.me"
	return Ask("RemoteBuildServer URI:", func() (err error) {
		vals.URI, err = ScanString(defaultServer)
		return
	})
}

func configRemoteBuildManager(root *Values) error {
	vals := &root.RemoteBuildManager
	defaultManager := "https://manager.lvh.me"
	return Ask("RemoteBuildManager URI:", func() (err error) {
		vals.URI, err = ScanString(defaultManager)
		return
	})
}

func ConfigRemoteBuild(root *Values) error {
	if err := configRemoteBuildServer(root); err != nil {
		return err
	}
	if err := configRemoteBuildManager(root); err != nil {
		return err
	}
	return nil
}
