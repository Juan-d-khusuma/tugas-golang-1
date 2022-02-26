package main

import (
  "fmt"
  "strconv"
)

func hitungVolume(r int, t int) float32 {
  const PI = 3.14
  return (float32(r*r))*PI*float32(t)
}

func main() {
  var _r string
  var _t string
  fmt.Scan(&_r, &_t)
  r, err := strconv.Atoi(_r)
  t, err := strconv.Atoi(_t)
  if err != nil {
    return 
  }
  fmt.Println(hitungVolume(r, t))
}