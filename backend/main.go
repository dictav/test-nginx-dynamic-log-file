package main

import (
	"math/rand"
	"net/http"
	"time"
)

var files = []string{
	"file_a",
	"file_b",
	"file_c",
}

func main() {
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		idx := rand.Perm(len(files))[0]
		rw.Header().Set("X-Log-File", files[idx])
	})

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

