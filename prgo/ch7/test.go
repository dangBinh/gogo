package main

import (
  "fmt"
)

func main() {
  jobList := []string{"abc", "ef"}

  jobs := make(chan []string)
  done := make(chan bool, len(jobList))
  go func() {
    for _, job := range jobList {
      jobs <- job
    }
    close(jobs)
  }()

  go func() {
    for job := range jobs {
      fmt.Println(job)
      done <- true
    }
  }()

  for i := 0; i < len(jobList); i++ {
    <-done
  }
}
