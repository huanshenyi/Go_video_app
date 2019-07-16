package main

import (
	"awesomeProject/govideo/api/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handlers.CreateUser)
	router.POST("/user/:user_name", handlers.Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8888",r)
}
