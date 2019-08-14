package values

import "fmt"

type CloudLibraryServer struct {
	URI string
}

func ConfigCloudLibrary(root *Values) {
	vals := &root.CloudLibraryServer
	fmt.Println("CloudLibraryServer URI:")
	defaultURI := "https://library.lvh.me"
	fmt.Printf("[%s] ", defaultURI)
	fmt.Scanln(&vals.URI)
	if vals.URI == "" {
		vals.URI = defaultURI
	}
}
