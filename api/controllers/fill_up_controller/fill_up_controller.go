package fill_up_controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mpg-tracker/api/models"
	. "mpg-tracker/api/models/errors"
	. "mpg-tracker/api/repositories/fill_up_repository"
	"net/http"
	"strconv"
)

func handleError(err error, w http.ResponseWriter, r *http.Request, defaultStatusCode int) {
	if err != nil {
		log.Println(err)
		switch err.(type) {
		case *NotFoundError:
			http.NotFound(w, r)
		case *NotImplementedError:
			http.Error(w, err.Error(), http.StatusNotImplemented)
		default:
			http.Error(w, err.Error(), defaultStatusCode)
		}
	}
}

// FillUpController contains router handlers pertaining to fill ups
type FillUpController struct {
	Repository *FillUpRepository
}

func (c FillUpController) PostFillUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var body models.FillUpEntity
	err := json.NewDecoder(r.Body).Decode(&body)
	fillUps, err := c.Repository.Create(body)

	err = json.NewEncoder(w).Encode(fillUps)
	if err != nil {
		handleError(err, w, r, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c FillUpController) GetListFillUps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	fillUps, err := c.Repository.Get()

	err = json.NewEncoder(w).Encode(fillUps)
	if err != nil {
		handleError(err, w, r, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c FillUpController) GetFillUpById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, idParseError := strconv.ParseUint(vars["fillupId"], 10, 32)
	if idParseError != nil {
		handleError(idParseError, w, r, http.StatusBadRequest)
		return
	}

	fillUp, findError := c.Repository.GetById(uint(id))
	if findError != nil {
		handleError(findError, w, r, http.StatusInternalServerError)
		return
	}
	if encodeError := json.NewEncoder(w).Encode(fillUp); encodeError != nil {
		handleError(encodeError, w, r, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c FillUpController) PutFillUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, idParseError := strconv.ParseUint(mux.Vars(r)["fillupId"], 10, 32)
	if idParseError != nil {
		handleError(idParseError, w, r, http.StatusBadRequest)
		return
	}

	var body models.FillUpEntity
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleError(err, w, r, http.StatusBadRequest)
		return
	}

	fillUps, err := c.Repository.Put(uint(id), body)
	if err = json.NewEncoder(w).Encode(fillUps); err != nil {
		handleError(err, w, r, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c FillUpController) DeleteFillUpById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id, idParseError := strconv.ParseUint(mux.Vars(r)["fillupId"], 10, 32)
	if idParseError != nil {
		handleError(idParseError, w, r, http.StatusBadRequest)
		return
	}

	fillUp, deleteError := c.Repository.Delete(uint(id))
	if deleteError != nil {
		handleError(deleteError, w, r, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(fillUp); err != nil {
		handleError(err, w, r, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
