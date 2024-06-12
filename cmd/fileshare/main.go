package main

import (
	"fileShare/internal"
)

func main() {
	err := internal.NewRouter()

	if err != nil {
		panic(err)
	}
}
