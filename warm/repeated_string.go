package warm

import (
	"strings"
)

func RepeatedString(s string, n int64) int64 {

	// check how much a in a string
	// split string to slice
	sSlice := strings.Split(s, "")

	aPerString := 0

	for i := range sSlice {
		if sSlice[i] == "a" {
			aPerString++
		}

	}

	perCount := int64(n / int64(len(s)))

	// mendapatkan semua banyak char a kecuali di a
	result := perCount * int64(aPerString)

	// check sisa menggunaakan module
	// jika tidak gunakan yang lain
	sisaCount := int(n) % len(sSlice)
	if sisaCount != 0 {
		sisaString := sSlice[:sisaCount]

		for i := range sisaString {
			if sisaString[i] == "a" {
				result += 1
			}

		}

	}

	return result

}
