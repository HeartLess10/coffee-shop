package handlers

import (
	"context"
	"net/http"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager"
	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	ctx *context.Context
	dbm *dbManager.DbManager
}

func NewHandler(dbm *dbManager.DbManager) Handler {
	return &handler{dbm: dbm}
}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	user := models.User[primitive.ObjectID]{Name: "HHH", Age: 15, Email: "tEST@GMAIL", UserID: "????"}
	dbManager.GetDbCollection[models.User[primitive.ObjectID]](h.dbm).InsertOne(*h.ctx, user)
}
