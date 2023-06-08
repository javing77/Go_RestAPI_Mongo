package main

import (
	"fmt"
	"net/http"

	"github.com/javing77/Go_RestAPI_Mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb+srv://localhost:27017")
	// Connect mongo with user mongo and password mongo

	if err != nil {
		fmt.Println("Error connecting to mongo: ", err)
		panic(err)
	}
	return s
}
