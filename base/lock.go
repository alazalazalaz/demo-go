package main
import (
    "fmt"
    "time"
)
func main() {
    var a = 0
    for i := 0; i < 1000; i++ {
        go func(idx int) {
            a += 1
            fmt.Println(a)
        }(i)
    }
    time.Sleep(time.Second)
}