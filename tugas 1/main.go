package main
import (
  "fmt"
  "strings" 
  "strconv"
)

func main() {
  fmt.Println("n :")
  var n string
  fmt.Scanln(&n)
  jumlah := 0
  _n, _ := strconv.Atoi(n)
  for i := 0; i< _n; i++ {
    var input string
    var _t1, _t2, _t3 string
    fmt.Scanln(&_t1, &_t2, &_t3)
    t1, err := strconv.Atoi(_t1)
    t2, err := strconv.Atoi(_t2)
    t3, err:= strconv.Atoi(_t3)
    if err != nil {
      fmt.Println(err)
      return
    }
    if t1 + t2 + t3 >= 2 {
      jumlah++
    }
    fmt.Println(jumlah)
  }
}