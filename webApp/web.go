package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
  "html/template"
  "time"
  "crypto/md5"
  "strconv"
  "io"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
  r.ParseForm() // parse url
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key: ", k)
    fmt.Println("val: ", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello binh")
}

func login(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method", r.Method) // request Method post put get
  if r.Method == "GET" {
    curtime := time.Now().Unix()
    h := md5.New() // return hash md5
    io.WriteString(h, strconv.FormatInt(curtime, 10)) // return string with base 10 (write content)
    token := fmt.Sprintf("%x", h.Sum(nil))

    t, _ := template.ParseFiles("login.gtpl")
    t.Execute(w, token)
  } else {
    r.ParseForm()
    token := r.Form.Get("token")
    if token != "" {

    } else {

    }
    fmt.Println("Username: ", template.HTMLEscapeString(r.Form.Get("username")))
    fmt.Println("passwornd:", template.HTMLEscapeString(r.Form.Get("password")))
    template.HTMLEscape(w, []byte(r.Form.Get("username")))
  }
}

func main() {
  http.HandleFunc("/", sayHelloName)
  http.HandleFunc("/login", login)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServer:", err)
  }
}
