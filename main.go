package main

import (
	"fmt"
	"github.com/kinecosystem/agora-common/kin"
	"strings"
)

func main() {
	fmt.Println("Enter search string:")

	var search string

	// Taking input from user
	fmt.Scanln(&search)

	found := false

	i := 0

	for {
		i++

		kp := generateKey(search)

		if kp != nil {
			found = true
			fmt.Println("Found a match in iteration %i", i)
		}

		if found {
			break
		}
		if i > 99999 {
			fmt.Println("None found...")
			break
		}
	}

}

func generateKey(prefix string) kin.PrivateKey {
	kp, _ := kin.NewPrivateKey()

	if strings.HasPrefix(strings.ToLower(kp.Public().Base58()), strings.ToLower(prefix)) {
		fmt.Println("Found a match:")
		fmt.Println("Secret: ", kp.StellarSeed())
		fmt.Println("Public Key: ", kp.Public().Base58())
		return kp
	}

	return nil
}
