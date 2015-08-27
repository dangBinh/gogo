package main

import (
  "fmt"
  "strings"
  "log"
  "sort"
)

func main() {
  irregularMatrix := [][]int{{1, 2, 3, 4},
                        {5, 6, 7, 8},
                        {9, 10, 11},
                        {12, 13, 14, 15},
                        {16, 17, 18, 19, 20}}
  slice := Flatten(irregularMatrix)
  fmt.Printf("1x%d: %v\n", len(slice), slice)
  my2d := Make2D(slice, 3)
  fmt.Printf("%v", my2d)
  var array = []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
  unique := UniqueInts(array)
  fmt.Println(unique)
  iniData := []string{
    "; Cut down copy of Mozilla application.ini file",
    "",
    "[App]",
    "Vendor=Mozilla",
    "Name=Iceweasel",
    "Profile=mozilla/firefox",
    "Version=3.5.16",
    "[Gecko]",
    "MinVersion=1.9.1",
    "MaxVersion=1.9.1.*",
    "[XRE]",
    "EnableProfileMigrator=0",
    "EnableExtensionManager=1",
    }
    ini := ParseIni(iniData)
    PrintIni(ini)
}

func Make2D(slice []int, columnCount int) [][]int  {
  matrix := make([][]int, len(slice))
  for i, x := range slice {
    row := i / columnCount
    column := i % columnCount
    if matrix[row] == nil {
      matrix[row] = make([]int, columnCount)
    }
    matrix[row][column] = x
  }
  return matrix
}

func Flatten(slice [][]int) []int {
  // flatArr := []int{}
  flatArr := make([]int, 0, len(slice) + len(slice[0]))
  for _, x := range slice {
    for _, y := range x {
      flatArr = append(flatArr, y)
    }
  }
  return flatArr
}

func UniqueInts(slice []int) []int {
  seen := map[int]bool{}
  unique := []int{}
  for _, x := range slice {
    if _, found := seen[x]; !found {
      unique = append(unique, x)
      seen[x] = true
    }
  }
  return unique
}

func ParseIni(lines []string) map[string]map[string]string {
  const separator = "="
  ini := make(map[string]map[string]string)
  group := "General"
  for _, line := range lines {
    if line == "" || strings.HasPrefix(line, ";") {
      continue
    }
    if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
      group = line[1 : len(line) - 1]
    } else if i := strings.Index(line, separator); i > -1 {
      key := line[:i]
      value := line[i+len(separator):]
      if _, found := ini[group]; !found {
        ini[group] = make(map[string]string)
      }
      ini[group][key] = value
    } else {
      log.Print("Failed")
    }
  }
  return ini
}

func PrintIni (ini map[string]map[string]string) {
  groups := make([]string, 0, len(ini))
  for group := range ini {
    groups = append(groups, group)
  }
  sort.Strings(groups)
  for i, group := range groups {
    fmt.Printf("[%s]\n", group)
    keys := make([]string, 0, len(ini[group]))
    for key := range ini[group] {
      keys = append(keys, key)
    }
    sort.Strings(keys)
    for _, key := range keys {
      fmt.Printf("%s=%s\n", key, ini[group][key])
    }
    if i+1 < len(group) {
      fmt.Println()
    }
  }
}
