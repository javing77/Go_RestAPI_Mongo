package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/javing77/Go_RestAPI_Mongo/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mongo.Collection
}

func NewUserController(s *mongo.Collection) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	// find the id in mongo using mongo-driver
	if err := uc.session.FindOne(r.Context(), bson.M{"_id": oid}).Decode(&u); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	js, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(js)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	// insert the user in mongo using mongo-driver
	uc.session.InsertOne(r.Context(), u)
	// uc.session.DB("mong-golang").C("users").Insert(u)

	js, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	w.Write(js)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	// delete the user in mongo using mongo-driver
	if err := uc.session.FindOneAndDelete(r.Context(), bson.M{"_id": oid}).Decode(&models.User{}); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte("Deleted\n"))
}
