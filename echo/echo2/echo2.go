// Echo2 prints its command-line arguments
package main

import (
    "fmt"
    "os"
//    "strings"
)

//func main() {
//    s, seq, newLine := "", "", "\n"
//    for index, arg := range os.Args[0:] {
//        s = seq + arg + newLine
//        seq = " "
//        fmt.Println(fmt.Sprintf("%d %s", index, s))
//    }
//}

//func main() {
//    fmt.Println(strings.Join(os.Args[1:], " "))
//    fmt.Println(os.Args[0:])
//}

func main() {
    for k, arg := range os.Args[0:] {
        fmt.Println(k, arg)
    }
}
