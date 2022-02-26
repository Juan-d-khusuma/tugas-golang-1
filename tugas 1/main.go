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
    fmt.Scan(&input)
    splitted_input := strings.Fields(input)
    t1, err := strconv.Atoi(splitted_input[0])
    t2, err := strconv.Atoi(splitted_input[1])
    t3, err:= strconv.Atoi(splitted_input[2])
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