package shortner

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	n, err := algorithm.Write([]byte(input))
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Print(n)
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
