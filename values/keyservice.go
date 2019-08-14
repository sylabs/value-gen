package values

import "fmt"

type KeyService struct {
	URI string
}

func ConfigKeyService(root *Values) {
	vals := &root.KeyService
	fmt.Println("KeyService URI:")
	defaultKeys := "https://keys.lvh.me"
	fmt.Printf("[%s] ", defaultKeys)
	fmt.Scanln(&vals.URI)
	if vals.URI == "" {
		vals.URI = defaultKeys
	}
}
