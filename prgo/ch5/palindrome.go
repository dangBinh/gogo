package main

import (
  "fmt"
  "os"
  "path/filepath"
  "unicode/utf8"
)
var IsPalindrome func(string) bool // hold reference for function

func init() {
  if len(os.Args[1]) > 1 && (os.Args[1] == "-a" || os.Args[1] == "--ascii") {
    os.Args = append(os.Args[:1], os.Args[2:]...)
    IsPalindrome  = func (s string) bool {
      i := 0
      for i <= len(s) / 2 {
        if s[i] != s[len(s) - i - 1] {
          return false
        } else {
          i++
        }
      }
      return true
    }
  } else {
    IsPalindrome  = func (s string) bool {
      for len(s) > 0 {
        first, sizeOfFirst := utf8.DecodeRuneInString(s)
        if sizeOfFirst == len(s) { // one rune
          break
        }
        last, sizeOfLast := utf8.DecodeLastRuneInString(s)
        if first != last {
          return false
        }
        s = s[sizeOfFirst : len(s) - sizeOfLast]
      }
      return true
    }
  }
}

func main() {
    if len(os.Args) == 1 {
        fmt.Printf("usage: %s [-a|--ascii] word1 [word2 [... wordN]]\n",
            filepath.Base(os.Args[0]))
        os.Exit(1)
    }
    words := os.Args[1:]
    for _, word := range words {
        fmt.Printf("%5t %q\n", IsPalindrome(word), word)
    }
}
