#!/usr/bin/env tclsh

set VERSION "0.1.0"
set MAX_TRIES 10

set UPPER "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
set LOWER "abcdefghijklmnopqrstuvwxyz"
set NUMBER "0123456789"
set SYMBOL "`~!@#$%^&*()_+-=[]\\{}|;':\",./<>?"

set upper "y"
set lower "y"
set number "y"
set symbol "y"
set l 16
set n 1

for {set i 0} {$i < [llength $argv]} {incr i} {
  if {[string equal [lindex $argv $i] "-version"]} {
    puts $VERSION
    exit
  }
  if {[string equal [lindex $argv $i] "-xu"]} {
    set upper "n"
  }
  if {[string equal [lindex $argv $i] "-xl"]} {
    set lower "n"
  }
  if {[string equal [lindex $argv $i] "-xn"]} {
    set number "n"
  }
  if {[string equal [lindex $argv $i] "-xs"]} {
    set symbol "n"
  }
  if {[string equal [lindex $argv $i] "-l"]} {
    set len [lindex $argv [expr $i + 1]]
    if {[string is integer $len]} {
      set l $len
    } else {
      puts stderr "Length argument given but no intger value followed"
      exit 1
    }
  }
  if {[string equal [lindex $argv $i] "-n"]} {
    set num [lindex $argv [expr $i + 1]]
    if {[string is integer $num]} {
      set n $num
    } else {
      puts stderr "Number argument given but no intger value followed"
      exit 1
    }
  }
}

set charset ""
if {[string equal $upper "y"]} {
  append charset $UPPER
}
if {$lower eq "y"} {
  append charset $LOWER
}
if {$number eq "y"} {
  append charset $NUMBER
}
if {$symbol eq "y"} {
  append charset $SYMBOL
}

proc generate {charset len} {
  set str ""
  for {set i 0} {$i < $len} {incr i} {
    set x [expr int(rand() * [string len $charset])]
    append str [string index $charset $x]
  }
  # TODO check for character classes in str
  return $str
}

for {set i 0} {
puts [generate $charset $l]