package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Schattenbrot/nos-api/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createWeapon is the handler for the InsertWeapon method.
func (app *application) createWeapon(w http.ResponseWriter, r *http.Request) {
	var weapon models.Weapon

	err := json.NewDecoder(r.Body).Decode(&weapon)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	type jsonResp struct {
		OK bool                `json:"ok"`
		ID *primitive.ObjectID `json:"_id"`
	}

	id, err := app.models.DB.InsertWeapon(weapon)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
		ID: id,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// findAllWeapons is the handler for the FindAllWeapons method.
func (app *application) findAllWeapons(w http.ResponseWriter, r *http.Request) {
	weapons, err := app.models.DB.FindAllWeapons()
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, weapons, "weapons")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) findAllWeaponsByProfession(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	profession := params.ByName("profession")

	weapons, err := app.models.DB.FindAllWeaponsByProfession(profession)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, weapons, "weapons")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) findOneWeaponById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := primitive.ObjectIDFromHex(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	weapon, err := app.models.DB.FindOneWeaponById(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, weapon, "weapon")
	if err != nil {
		app.logger.Println(err)
	}
}
