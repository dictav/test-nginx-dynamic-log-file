package main

import (
	"math/rand"
	"net/http"
	"time"
)

var versions = []string{
	"",
	"10",
	"11",
	"12",
	"13",
	"14",
}

func main() {
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		idx := rand.Perm(len(versions))[0]
		v := versions[idx]
		if v != "" {
			rw.Header().Set("X-Version", v)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

