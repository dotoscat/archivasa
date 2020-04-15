package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dotoscat/archivasa/content"
)

func main() {
	fmt.Println("Hola mundo")
	cwd, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Current working directory", cwd)
	content.Export("testing", cwd)
}
