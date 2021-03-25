package main

import (
	"github.com/ZupIT/horusec-platform/api/config/providers"
)

func main() {
	router, err := providers.Initialize("8000")
	if err != nil {
		panic(err)
	}

	router.ListenAndServe()
}