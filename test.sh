#!/bin/sh

VERSION="0.2.0"

if [ "$#" -lt 1 ]; then
  echo "program argument missing" >/dev/stderr
  exit 1
fi

if [ ! -f "$1" ]; then
  echo "Program file not found" >/dev/stderr
  exit 1
fi

if [ ! -x "$1" ]; then
  echo "Program file is not executable" >/dev/stderr
  exit 1
fi

exe=$1

echo "Test v$VERSION on $exe"

function assert_run {
  $($1 &>/dev/null)
  if [ $? -gt 0 ]; then
    printf "FAIL: %s\n" "$2"
    return 1
  fi
  printf "PASS: %s\n" "$2"
}

function assert_length {
  str=$($1)
  if [ ${#str} -ne $2 ]; then
    printf "FAIL: Expected length %s, got %s\n" "$2" ${#str}
    return 1
  fi
  printf "PASS: %s\n" "$3"
}

assert_run "$exe -h" "-h option"
assert_run "$exe -version" "-version option"
assert_length "$exe" 16 "default length"
assert_length "$exe -l 1000" 1000 "1k Long string"
