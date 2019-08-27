package values

type RemoteBuildServer struct {
	Hostname string
}

type RemoteBuildManager struct {
	Hostname string
}

func configRemoteBuildServer(root *Values) error {
	vals := &root.RemoteBuildServer
	defaultServer := "build." + root.DefaultDomain
	return Ask("RemoteBuildServer URI:", func() (err error) {
		vals.Hostname, err = ScanString(defaultServer)
		return
	})
}

func configRemoteBuildManager(root *Values) error {
	vals := &root.RemoteBuildManager
	defaultManager := "manager." + root.DefaultDomain
	return Ask("RemoteBuildManager URI:", func() (err error) {
		vals.Hostname, err = ScanString(defaultManager)
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
