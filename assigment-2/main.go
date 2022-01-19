package main

import (
	"assignment-7/router"

	_ "github.com/lib/pq"
)

func main() {
	router.StartRouter()
}
