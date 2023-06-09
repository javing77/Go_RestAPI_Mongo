package main

import (
	"net/http"

	"github.com/javing77/Go_RestAPI_Mongo/configs"
	"github.com/javing77/Go_RestAPI_Mongo/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(configs.GetCollection(configs.DB, "users"))

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}
