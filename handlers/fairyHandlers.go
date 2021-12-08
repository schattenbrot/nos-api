package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Schattenbrot/nos-api/config"
	"github.com/Schattenbrot/nos-api/models"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertFairy(w http.ResponseWriter, r *http.Request) {
	var fairy models.Fairy

	err := json.NewDecoder(r.Body).Decode(&fairy)
	if err != nil {
		errorJSON(w, err)
	}

	type jsonResp struct {
		OK bool                `json:"ok"`
		ID *primitive.ObjectID `json:"_id"`
	}

	id, err := config.App.Models.DB.InsertFairy(fairy)
	if err != nil {
		errorJSON(w, err)
		return
	}

	ok := jsonResp{OK: true, ID: id}

	err = writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		errorJSON(w, err)
		return
	}
}

// findAllFairies is the handler for the FindAllFairies method.
func FindAllFairies(w http.ResponseWriter, r *http.Request) {
	fairies, err := config.App.Models.DB.FindAllFairies()
	if err != nil {
		config.App.Logger.Println(err)
	}

	err = writeJSON(w, http.StatusOK, fairies, "fairies")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func FindAllFairiesByElement(w http.ResponseWriter, r *http.Request) {
	element := chi.URLParam(r, "element")

	fairies, err := config.App.Models.DB.FindAllFairiesByElement(element)
	if err != nil {
		config.App.Logger.Println(err)
	}

	err = writeJSON(w, http.StatusOK, fairies, "fairies")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func FindFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	fairy, err := config.App.Models.DB.FindFairyById(id)
	if err != nil {
		config.App.Logger.Println(err)
	}

	err = writeJSON(w, http.StatusOK, fairy, "fairy")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func UpdateFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	var updateFairy models.Fairy
	err = json.NewDecoder(r.Body).Decode(&updateFairy)
	if err != nil {
		errorJSON(w, err)
		return
	}

	result, err := config.App.Models.DB.UpdateFairyById(id, updateFairy)
	if err != nil {
		errorJSON(w, err)
		return
	}

	err = writeJSON(w, http.StatusOK, result, "updated")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func DeleteFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	result, err := config.App.Models.DB.DeleteFairyById(id)
	if err != nil {
		errorJSON(w, err)
		return
	}

	err = writeJSON(w, http.StatusOK, result, "deleted")
	if err != nil {
		config.App.Logger.Println(err)
	}
}