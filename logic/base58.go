package logic

import (
	"math"
	"strings"
)

const alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

// EncodeB58 encode an int into a base 58 string
func EncodeB58(id int) string {
	hashString := ""

	for id > 0 {
		remain := id % 58
		id = id / 58

		hashString = hashString + alphabet[remain:remain+1]
	}

	return hashString
}

// DecodeB58 hash to an int
func DecodeB58(hash string) int {
	var power float64
	var decimal int

	for i := 0; i < len(hash); i++ {
		digit := hash[i : i+1]
		decimal += strings.Index(alphabet, digit) * int(math.Pow(58, power))

		power++
	}

	return decimal
}
