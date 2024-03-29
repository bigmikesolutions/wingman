package main

import (
	"net/http"

	svc "github.com/bigmikesolutions/wingman/service/http"
)

func main() {
	r, err := svc.NewRouter()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8000", r)
}
