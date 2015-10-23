package main

import (
  "fmt"
)

type Optioner interface {
  Name() string // tag in Golang
  IsValid() bool
}

type OptionCommon struct {
  ShortName string
  LongName string
}

type IntOption struct {
  OptionCommon
  Value, Min, Max int
}

func (option IntOption) Name() string {
  return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool  {
  return option.Min <= option.Value && option.Value <= option.Max
}

func name(shortName, longName string) string {
  if longName == "" {
    return shortName
  }
  return longName
}

type StringOption struct {
  OptionCommon
  Value string
}

func (option StringOption) Name() string  {
  return name(option.ShortName, option.LongName)
}

func (option StringOption) IsValid() bool {
  return true
}

type FloatOption struct {
  OptionCommon // Anonymous field
  Value float64
}

func (option FloatOption) Name() string  {
  return name(option.ShortName, option.LongName)
}

func (option FloatOption) IsValid() bool {
  return true
}

type GenericOption struct {
  OptionCommon
}

func (option GenericOption) Name() string {
  return name(option.ShortName, option.LongName)
}

func (option GenericOption) IsValid() bool {
  return true
}

func main() {
  fileOption := StringOption{OptionCommon{"f", "file"}, "index.html"}
  topOption := IntOption{ OptionCommon: OptionCommon{"t", "top"}, Max: 100 }
  sizeOption := FloatOption{GenericOption {OptionCommon{"s", "size"}, 19.5}}
  for _, option := range []Optioner{topOption, fileOption, sizeOption} {
    fmt.Print("name=", option.Name(), " • valid=", option.IsValid())
    fmt.Print(" • value=")
    switch option := option.(type) { // shadow variable
      case IntOption:
        fmt.Print(option.Value, " • min=", option.Min, " • max=", option.Max, "\n")
      case StringOption:
        fmt.Println(option.Value)
      case FloatOption:
        fmt.Println(option.Value)
    }
  }
}
