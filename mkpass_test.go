package mkpass

import (
	"strings"
	"testing"
)

func TestGenerateWithDefaultOptions(t *testing.T) {
	s, err := Generate()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(s) != DefaultOptions.Length {
		t.Errorf("Expecting length %d, got %d", DefaultOptions.Length, len(s))
	}
}

func TestConfigureWithEmptyOptions(t *testing.T) {
	err := Configure(Options{})
	if err == nil {
		t.Errorf("Expecting empty charset error, received nil error")
	}
}

func TestGenerateUpperOnly(t *testing.T) {
  err := Configure(Options{
    Upper: true,
  })
	if err != nil {
		t.Fatal(err.Error())
	}
  s, err := Generate()
	if err != nil {
		t.Fatal(err.Error())
	}
	if strings.ContainsAny(s, LOWER) {
		t.Errorf("Unexpected lower case character in %s", s)
	}
	if strings.ContainsAny(s, NUMBER) {
		t.Errorf("Unexpected number character in %s", s)
	}
	if strings.ContainsAny(s, SYMBOL) {
		t.Errorf("Unexpected symbol character in %s", s)
	}
}

func TestGenerateLowerOnly(t *testing.T) {
	err := Configure(Options{
		Lower: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := Generate()
	if err != nil {
		t.Fatal(err.Error())
	}
	if strings.ContainsAny(s, UPPER) {
		t.Errorf("Unexpected upper case character in %s", s)
	}
	if strings.ContainsAny(s, NUMBER) {
		t.Errorf("Unexpected number character in %s", s)
	}
	if strings.ContainsAny(s, SYMBOL) {
		t.Errorf("Unexpected symbol character in %s", s)
	}
}

func TestGenerateNumberOnly(t *testing.T) {
	err := Configure(Options{
		Number: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := Generate()
	if err != nil {
		t.Fatal(err.Error())
	}
	if strings.ContainsAny(s, UPPER) {
		t.Errorf("Unexpected upper case character in %s", s)
	}
	if strings.ContainsAny(s, LOWER) {
		t.Errorf("Unexpected lower case character in %s", s)
	}
	if strings.ContainsAny(s, SYMBOL) {
		t.Errorf("Unexpected symbol character in %s", s)
	}
}

func TestGenerateSymbolOnly(t *testing.T) {
	err := Configure(Options{
		Symbol: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := Generate()
	if err != nil {
		t.Fatal(err.Error())
	}
	if strings.ContainsAny(s, UPPER) {
		t.Errorf("Unexpected upper case character in %s", s)
	}
	if strings.ContainsAny(s, LOWER) {
		t.Errorf("Unexpected lower case character in %s", s)
	}
	if strings.ContainsAny(s, NUMBER) {
		t.Errorf("Unexpected number character in %s", s)
	}
}

func BenchmarkGenerateWithDefaultOptions(b *testing.B) {
  for i := 0; i < b.N; i++ {
	  Generate()
  }
}

func BenchmarkGenerate10kLong(b *testing.B) {
  opt := DefaultOptions
  opt.Length = 10000
  Configure(opt)
  for i := 0; i < b.N; i++ {
    Generate()
  }
}
