package handlers

import (
	"encoding/json"
	"fmt"
	"link-converter/models"
	"net/http"
)

type ResponseHandler func(http.ResponseWriter, *http.Request) (models.ResponseModel, *models.CustomError)

func (fn ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := fn(w, r)

	if err != nil {
		responseJson(w, err, http.StatusBadRequest)
		return
	}

	responseJson(w, response, http.StatusOK)
}

func responseJson(w http.ResponseWriter, response interface{}, httpStatus int) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(jsonResponse)
}
