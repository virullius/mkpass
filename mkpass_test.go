package mkpass

import (
  "testing"
  "strings"
)

func TestGenerateDefault(t *testing.T) {
  s, err := Generate(DefaultOptions)
  if err != nil {
    t.Fatal(err.Error())
  }
  if len(s) != DefaultOptions.Length {
    t.Errorf("Expecting length %d, got %d", DefaultOptions.Length, len(s))
  }
}

func TestGenerateEmptySet(t *testing.T) {
  _, err := Generate(Options{})
  if err == nil {
    t.Errorf("Expecting empty charset error, received nil error")
  }
}

func TestGenerateUpperOnly(t *testing.T) {
  s, err := Generate(Options{
    Upper: true,
  })
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
  s, err := Generate(Options{
    Lower: true,
  })
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
  s, err := Generate(Options{
    Number: true,
  })
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
  s, err := Generate(Options{
    Symbol: true,
  })
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