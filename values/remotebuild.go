package values

import "fmt"

type RemoteBuildServer struct {
	URI string
}

type RemoteBuildManager struct {
	URI string
}

func configRemoteBuildServer(root *Values) {
	vals := &root.RemoteBuildServer
	fmt.Println("RemoteBuildServer URI:")
	defaultServer := "https://build.lvh.me"
	fmt.Printf("[%s] ", defaultServer)
	fmt.Scanln(&vals.URI)
	if vals.URI == "" {
		vals.URI = defaultServer
	}
}

func configRemoteBuildManager(root *Values) {
	vals := &root.RemoteBuildManager
	fmt.Println("RemoteBuildManager URI:")
	defaultManager := "https://manager.lvh.me"
	fmt.Printf("[%s] ", defaultManager)
	fmt.Scanln(&vals.URI)
	if vals.URI == "" {
		vals.URI = defaultManager
	}
}

func ConfigRemoteBuild(root *Values) {
	configRemoteBuildServer(root)
	configRemoteBuildManager(root)
}
