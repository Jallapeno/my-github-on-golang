package main

import (
	"fmt"
	"log"
	"my-github-on-golang/app"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("My Github on GoLang")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro to load .env file:", err)
		return
	}

	app := app.App()

	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
