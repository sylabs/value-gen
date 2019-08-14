package values

import "fmt"

type Frontend struct {
	URI       string
	RevokeURI string
}

func ConfigFrontend(root *Values) {
	vals := &root.Frontend
	fmt.Println("Frontend URI:")
	fmt.Print("[https://cloud.lvh.me] ")
	fmt.Scanln(&vals.URI)
	defaultRevoke := vals.URI + "/assets/html/auth-revoke.html"
	fmt.Println("Revoke URI:")
	fmt.Printf("[%s] ", defaultRevoke)
	fmt.Scanln(&vals.RevokeURI)
	if vals.RevokeURI == "" {
		vals.RevokeURI = defaultRevoke
	}
}
