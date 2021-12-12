package mkpass

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	VERSION   string = "0.2.0"
	MAX_TRIES int    = 10

	UPPER  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER  string = "abcdefghijklmnopqrstuvwxyz"
	NUMBER string = "0123456789"
	SYMBOL string = "`~!@#$%^&*()_+-=[]\\{}|;':\",./<>?"
)

var DefaultOptions = Options{
	Upper:  true,
	Lower:  true,
	Number: true,
	Symbol: true,
	Length: 16,
}
var opts = DefaultOptions
var charset string

type Options struct {
	Upper  bool
	Lower  bool
	Number bool
	Symbol bool
	Length int
}

func init() {
  Configure(DefaultOptions)
}

// Generate character set based on Options
func Configure(o Options) error {
  cs := ""
	if o.Upper {
		cs += UPPER
	}
	if o.Lower {
		cs += LOWER
	}
	if o.Number {
		cs += NUMBER
	}
	if o.Symbol {
		cs += SYMBOL
	}
	if o.Length == 0 {
		o.Length = DefaultOptions.Length
	}
	if len(cs) == 0 {
		return fmt.Errorf("No character classes choosen, characterset is empty")
	}
  opts = o
  charset = cs
  return nil
}

func Generate() (string, error) {
	tries := 0
  str := ""

  // loop that breaks when circuit breaker trips or suitable string is generated
	for {
		tries++
		// circuit breaker for uncaught errors that prevent generation of suitable strings
		if tries > MAX_TRIES {
			return "", fmt.Errorf("Failed to generate suitable random string in %d tries", MAX_TRIES)
		}

		// pick random characters from set one at a time
		l := big.NewInt(int64(len(charset)))
		for i := 0; i < opts.Length; i++ {
			n, err := rand.Int(rand.Reader, l)
			if err != nil {
				return "", fmt.Errorf("Reading random integer: %s", err.Error())
			}
			str += string(charset[n.Int64()])
		}

		// check if each used class is included in output
		if opts.Upper && !strings.ContainsAny(str, UPPER) {
			continue
		}
		if opts.Lower && !strings.ContainsAny(str, LOWER) {
			continue
		}
		if opts.Number && !strings.ContainsAny(str, NUMBER) {
			continue
		}
		if opts.Symbol && !strings.ContainsAny(str, SYMBOL) {
			continue
		}
		break
	}

	return str, nil
}
