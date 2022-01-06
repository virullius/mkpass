package mkpass

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	VERSION   string = "0.3.2"
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

type Options struct {
	Upper  bool
	Lower  bool
	Number bool
	Symbol bool
	Length int
}

type Generator struct {
	opts    Options
	charset string
}

func New(o Options) (Generator, error) {
	g := Generator{
		opts: o,
	}
	nc := 0
	if o.Upper {
		g.charset += UPPER
		nc++
	}
	if o.Lower {
		g.charset += LOWER
		nc++
	}
	if o.Number {
		g.charset += NUMBER
		nc++
	}
	if o.Symbol {
		g.charset += SYMBOL
		nc++
	}
	if o.Length == 0 {
		g.opts.Length = DefaultOptions.Length
	}
	if len(g.charset) == 0 {
		return g, fmt.Errorf("No character classes chosen, characterset is empty")
	}
	if nc > g.opts.Length {
		return g, fmt.Errorf("Length is shorter than number of character classes")
	}
	return g, nil
}

func (g Generator) Generate() (string, error) {
	tries := 0
	var str string

	// loop that breaks when circuit breaker trips or suitable string is generated
	for {
		tries++
		// circuit breaker for uncaught errors that prevent generation of suitable strings
		if tries > MAX_TRIES {
			return "", fmt.Errorf("Failed to generate suitable random string in %d tries", MAX_TRIES)
		}

		str = ""
		// pick random characters from set one at a time
		l := big.NewInt(int64(len(g.charset)))
		for i := 0; i < g.opts.Length; i++ {
			n, err := rand.Int(rand.Reader, l)
			if err != nil {
				return "", fmt.Errorf("Reading random integer: %s", err.Error())
			}
			str += string(g.charset[n.Int64()])
		}

		// check if each used class is included in output
		if g.opts.Upper && !strings.ContainsAny(str, UPPER) {
			continue
		}
		if g.opts.Lower && !strings.ContainsAny(str, LOWER) {
			continue
		}
		if g.opts.Number && !strings.ContainsAny(str, NUMBER) {
			continue
		}
		if g.opts.Symbol && !strings.ContainsAny(str, SYMBOL) {
			continue
		}
		break
	}

	return str, nil
}
