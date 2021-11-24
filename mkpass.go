package mkpass

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	VERSION   string = "0.1.0"
	MAX_TRIES int    = 10

	UPPER  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER  string = "abcdefghijklmnopqrstuvwxyz"
	NUMBER string = "0123456789"
	SYMBOL string = "`~!@#$%^&*()_+-=[]\\{}|;':\",./<>?"
)

var (
	DefaultOptions Options = Options{
		Upper:  true,
		Lower:  true,
		Number: true,
		Symbol: true,
		Length: 16,
	}
)

type Options struct {
	Upper  bool
	Lower  bool
	Number bool
	Symbol bool
	Length int
}

func Generate(o Options) (string, error) {
	var (
		charset = ""
		output  = ""
	)
	if o.Upper {
		charset += UPPER
	}
	if o.Lower {
		charset += LOWER
	}
	if o.Number {
		charset += NUMBER
	}
	if o.Symbol {
		charset += SYMBOL
	}
	if o.Length == 0 {
		o.Length = DefaultOptions.Length
	}
	if len(charset) == 0 {
		return "", fmt.Errorf("No character classes choosen, characterset is empty")
	}

	tries := 0

	for {
		tries++
		// circuit breaker for uncaught errors that prevent generation of suitable strings
		if tries > MAX_TRIES {
			return "", fmt.Errorf("Failed to generate suitable random string in %d tries", MAX_TRIES)
		}

		output = ""
		// pick random characters from set one at a time
		l := big.NewInt(int64(len(charset)))
		for i := 0; i < o.Length; i++ {
			n, err := rand.Int(rand.Reader, l)
			if err != nil {
				return "", fmt.Errorf("Reading random integer: %s", err.Error())
			}
			output += string(charset[n.Int64()])
		}

		// check if each used class is included in output
		if o.Upper && !strings.ContainsAny(output, UPPER) {
			continue
		}
		if o.Lower && !strings.ContainsAny(output, LOWER) {
			continue
		}
		if o.Number && !strings.ContainsAny(output, NUMBER) {
			continue
		}
		if o.Symbol && !strings.ContainsAny(output, SYMBOL) {
			continue
		}
		break
	}

	return output, nil
}
