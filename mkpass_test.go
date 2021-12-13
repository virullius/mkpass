package mkpass

import (
	"strings"
	"testing"
)

func TestGenerateWithDefaultOptions(t *testing.T) {
	g, err := New(DefaultOptions)
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := g.Generate()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(s) != DefaultOptions.Length {
		t.Errorf("Expecting length %d, got %d", DefaultOptions.Length, len(s))
	}
}

func TestGenerateLengths(t *testing.T) {
	o := DefaultOptions
	for i := 10; i < 30; i++ {
		o.Length = i
		g, err := New(o)
		if err != nil {
			t.Fatal(err.Error())
		}
		s, err := g.Generate()
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(s) != i {
			t.Errorf("Expected length %d, got %d", i, len(s))
		}
	}
}

func TestConfigureWithEmptyOptions(t *testing.T) {
	_, err := New(Options{})
	if err == nil {
		t.Errorf("Expecting empty charset error, received nil error")
	}
}

func TestGenerateUpperOnly(t *testing.T) {
	g, err := New(Options{
		Upper: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := g.Generate()
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
	g, err := New(Options{
		Lower: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := g.Generate()
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
	g, err := New(Options{
		Number: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := g.Generate()
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
	g, err := New(Options{
		Symbol: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := g.Generate()
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
	g, err := New(DefaultOptions)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		g.Generate()
	}
}

func BenchmarkGenerate10kLong(b *testing.B) {
	o := DefaultOptions
	o.Length = 10000
	g, err := New(o)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		g.Generate()
	}
}
