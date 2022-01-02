package services

import (
	. "link-converter/constants"
	"link-converter/models"
	"link-converter/repository"
	"strings"
	"time"
)

var (
	repo           *repository.RedisRepository
	productService ProductService
	searchService  SearchService
)

type ConverterService struct{}

func NewConverterService(couchbaseRepo *repository.RedisRepository) ConverterService {
	repo = couchbaseRepo
	productService = NewProductService()
	searchService = NewSearchService()

	return ConverterService{}
}

func (*ConverterService) ToDeepLink(model models.RequestModel) (models.ResponseModel, *models.CustomError) {
	var response models.ResponseModel
	var err *models.CustomError

	dataResponse := getCacheLink(model.Link)
	if dataResponse != "" {
		return models.ResponseModel{ConvertedLink: dataResponse}, nil
	}

	resourceType := model.GetWebUrlResourceType()

	if resourceType == RESOURCE_TYPE_PRODUCT {
		response.ConvertedLink, err = productService.ToDeepLink(model.Link)
	} else if resourceType == RESOURCE_TYPE_SEARCH {
		response.ConvertedLink, err = searchService.ToDeepLink(model.Link)
	} else {
		response.ConvertedLink = getOtherLink(EVENT_TYPE_WEBURL_TO_DEEPLINK)
	}

	if err == nil {
		setEvent(model, response, EVENT_TYPE_WEBURL_TO_DEEPLINK, resourceType)
	}

	return response, err
}

func (*ConverterService) ToWebUrl(model models.RequestModel) (models.ResponseModel, *models.CustomError) {
	var response models.ResponseModel
	var err *models.CustomError

	dataResponse := getCacheLink(model.Link)
	if dataResponse != "" {
		return models.ResponseModel{ConvertedLink: dataResponse}, nil
	}

	resourceType := model.GetDeepLinkResourceType()

	if resourceType == RESOURCE_TYPE_PRODUCT {
		response.ConvertedLink, err = productService.ToWebUrl(model.Link)
	} else if resourceType == RESOURCE_TYPE_SEARCH {
		response.ConvertedLink, err = searchService.ToWebUrl(model.Link)
	} else {
		response.ConvertedLink = getOtherLink(EVENT_TYPE_DEEPLINK_TO_WEBURL)
	}

	if err == nil {
		setEvent(model, response, EVENT_TYPE_DEEPLINK_TO_WEBURL, resourceType)
	}

	return response, err
}

func getOtherLink(eventType string) string {
	var sb strings.Builder
	if eventType == EVENT_TYPE_DEEPLINK_TO_WEBURL {
		sb.WriteString(WEBURL_PREFIX)
	} else {
		sb.WriteString(DEEPLIK_PREFIX)
		sb.WriteString(HOME_PAGE)
	}

	return sb.String()
}

func setEvent(requestModel models.RequestModel, responseModel models.ResponseModel, eventType string, resourceType string) {
	if repo != nil && requestModel.Link != "" {
		data := models.CacheModel{
			Date:          time.Now().UTC(),
			ConvertedLink: responseModel.ConvertedLink,
			EventType:     eventType,
			ResourceType:  resourceType,
		}

		repo.SetKey(requestModel.Link, data, 0)
	}
}

func getCacheLink(key string) string {
	if repo != nil && key != "" {
		cacheObj := &models.CacheModel{}

		err := repo.GetKey(key, cacheObj)
		if err == nil {
			return cacheObj.ConvertedLink
		}
	}

	return ""
}
