package main

import (
	"fmt"
	"golang-api-template/app"
	"golang-api-template/helper"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	fmt.Println(port)
	fmt.Println("====================================")
	db := app.ConnectDatabase()

	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Println(fmt.Sprintf("Server is running on port %s", port))

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
