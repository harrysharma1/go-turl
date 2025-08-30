package urlshortener

import (
	"crypto/sha512"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha512f(input string) []byte {
	algo := sha512.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encode, err := encoding.Encode(bytes)
	if err != nil {
		log.Fatalf("error ocurred whilst encoding: %v", err.Error())
		os.Exit(1)
	}
	return string(encode)
}

func ShortLinkGeneration(initLink string, userId string) string {
	urlHash := sha512f(initLink + userId)
	genNo := new(big.Int).SetBytes(urlHash).Uint64()
	finalString := base58encoded([]byte(fmt.Sprintf("%d", genNo)))
	fmt.Print(finalString)
	return finalString[:8]
}
