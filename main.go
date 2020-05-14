package main

import (
	"fmt"
	"go-echo-sample/router"
)

func main() {
	fmt.Println("Welcome to the server")
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
