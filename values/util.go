package values

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomSecret(l int) string {
	var s strings.Builder
	s.Grow(l)

	for i := 0; i < l; i++ {
		r := rand.Intn(len(charset))
		s.WriteByte(charset[r])
	}
	return s.String()
}

func Ask(prompt string, scanFn func() error) error {
	tries := 3
	for i := 0; i < tries; i++ {
		fmt.Println(prompt)
		err := scanFn()
		if err != nil {
			fmt.Println(err)
		} else {
			return nil
		}
	}
	return errors.New("Aborting")
}

func ScanString(defaultChoice string) (string, error) {
	if defaultChoice != "" {
		fmt.Printf("[%s] ", defaultChoice)
	}
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		text = defaultChoice
	}
	return text, err
}

func ScanYesNo(defaultChoice bool) (bool, error) {
	var s string
	if defaultChoice {
		fmt.Print("[Y/n] ")
	} else {
		fmt.Print("[y/N] ")
	}
	b := bufio.NewReader(os.Stdin)
	s, err := b.ReadString('\n')
	if err != nil {
		return false, err
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return defaultChoice, nil
	}

	if s == "Y" || s == "y" {
		return true, nil
	}
	if s == "N" || s == "n" {
		return false, nil
	}
	return false, errors.New("Unrecognized input, expecting Y/N")
}
