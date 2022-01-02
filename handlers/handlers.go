package handlers

import (
	"encoding/json"
	"link-converter/errors"
	"link-converter/repository"
	"link-converter/services"
	"net/http"

	"link-converter/helpers"
	"link-converter/models"

	"go.uber.org/zap"
)

func ToDeepLink(w http.ResponseWriter, r *http.Request) (models.ResponseModel, *models.CustomError) {
	var model, err = parseRequestModel(w, r)
	if err != nil {
		return models.ResponseModel{}, err
	}

	repo := getRepository()
	converterService := services.NewConverterService(repo)
	response, err := converterService.ToDeepLink(model)

	return response, err
}

func ToWebUrl(w http.ResponseWriter, r *http.Request) (models.ResponseModel, *models.CustomError) {
	var model, err = parseRequestModel(w, r)
	if err != nil {
		return models.ResponseModel{}, err
	}

	repo := getRepository()
	converterService := services.NewConverterService(repo)
	response, err := converterService.ToWebUrl(model)

	return response, err
}

func parseRequestModel(w http.ResponseWriter, r *http.Request) (models.RequestModel, *models.CustomError) {
	var model models.RequestModel

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		return model, errors.NewCustomErr(err.Error())
	}

	customErr := helpers.LinkIsValid(model.Link)
	if customErr != nil {
		return model, customErr
	}

	return model, nil
}

func getRepository() *repository.RedisRepository {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	redisRepo := repository.NewRedisRepository(logger)

	return redisRepo
}
