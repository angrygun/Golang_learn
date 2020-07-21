// bench mark
package main

import (
    "strings"
    "testing"
)

//func string2Join(a []string) string {
//    return strings.Join(a, " ")
//}
//
//func string2Plus(a []string) string {
//    var s, seq string
//    for i := 0; i < len(a); i++ {
//        s += seq + a[i]
//        seq = " "
//    }
//    return s
//}
//
//func BenchmarkString2Join(b *testing.B) {
//    for i := 0; i< b.N; i++ {
//        string2Join([]string{"Welcome", "To", "China"})
//    }
//}
//
//func BenchmarkString2Plus(b *testing.B) {
//    for i := 0; i < b.N; i++ {
//        string2Plus([]string{"Welcome", "To", "China"})
//    }
//}

func BenchmarkString2Join(b *testing.B) {
  for i := 0; i < b.N; i++ {
      input := []string{"Welcome", "To", "China"}
      result := strings.Join(input, " ")
      if result != "Welcome To China" {
          b.Error("Unexcepted result:" + result)
      }
  }
}

func BenchmarkString2Plus(b *testing.B) {
    for i := 0; i < b.N; i++ {
        input := []string{"Welcome", "To", "China"}
        var s, seq string
        for j := 0; j < len(input); j++ {
            s += seq + input[j]
            seq = " "
        }
        if s != "Welcome To China" {
            b.Error("Unexcepted result:" + s)
        }
    }
}
