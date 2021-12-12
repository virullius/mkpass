package main

import (
  "flag"
  "mkpass"
  "fmt"
  "os"
)

var (
  printVersion bool
  xUpper, xLower, xNumber, xSymbol bool
  length, count int
)

func init() {
  flag.BoolVar(&printVersion, "version", false, "print version and quit")
  flag.BoolVar(&xUpper, "xu", false, "exclude upper case characters")
  flag.BoolVar(&xLower, "xl", false, "exclude lower case characters")
  flag.BoolVar(&xNumber, "xn", false, "exclude numeric characters")
  flag.BoolVar(&xSymbol, "xs", false, "exclude symbol characters")
  flag.IntVar(&length, "l", 16, "length of each generated string")
  flag.IntVar(&count, "n", 1, "number of strings to genereate")
}

func main() {
  flag.Parse()
  if printVersion {
    fmt.Println(mkpass.VERSION)
    os.Exit(0)
  }
  opt := mkpass.Options{
    Upper: !xUpper,
    Lower: !xLower,
    Number: !xNumber,
    Symbol: !xSymbol,
    Length: length,
  }
  mkpass.Configure(opt)
  for i := 0; i < count; i++ {
    s, err := mkpass.Generate()
    if err != nil {
      fmt.Fprintf(os.Stderr, err.Error())
      os.Exit(1)
    }
    fmt.Println(s)
  }
}
