package main

import (
	"net/http"

	"highload-architect/pkg/config"
	"highload-architect/pkg/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func handlerRequests() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/sign-in/", handlers.SignIn)
	http.HandleFunc("/sign-up/", handlers.SignUp)
	http.HandleFunc("/sign-out/", handlers.SignOut)
	http.HandleFunc("/profile/", handlers.Profile)
	http.HandleFunc("/friends/", handlers.Friends)

	http.ListenAndServe(":8080", nil)
}

func init() {
	if err := godotenv.Load(); err != nil {
		println("No .env file found")
	}
	
	println(config.GetDbConfig())
}

func main() {
	handlerRequests()
}
