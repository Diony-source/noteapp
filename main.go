package main

import (
	"fmt"
	"noteapp/handlers"
)

func main() {
	fmt.Println("Welcome to the Note Taking App")
	handlers.Start()
}