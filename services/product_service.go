package services

import (
	. "link-converter/constants"
	"link-converter/errors"
	"link-converter/helpers"
	"link-converter/models"
	"net/url"
	"strconv"
	"strings"
)

type ProductService struct{}

func NewProductService() ProductService {
	return ProductService{}
}

func (*ProductService) ToDeepLink(weburl string) (string, *models.CustomError) {
	parameters, err := helpers.WebUrlParse(weburl)
	if err != nil {
		return "", err
	}

	err = productUrlIsValid(parameters)
	if err != nil {
		return "", err
	}

	values := url.Values{}

	values.Add(PRODUCT_KEY_CONTENT_ID, getProductContentId(parameters["path"]))

	if parameters[PRODUCT_KEY_BOUTIQUE_ID] != "" {
		values.Add(PRODUCT_KEY_CAMPAIGN_ID, parameters[PRODUCT_KEY_BOUTIQUE_ID])
	}
	if parameters[PRODUCT_KEY_MERCHANT_ID] != "" {
		values.Add(PRODUCT_KEY_MERCHANT_ID, parameters[PRODUCT_KEY_MERCHANT_ID])
	}

	var deeplink strings.Builder
	deeplink.WriteString(DEEPLIK_PREFIX)
	deeplink.WriteString(PRODUCT_DEEPLINK_PAGE)
	deeplink.WriteString(AND_MARK)
	deeplink.WriteString(values.Encode())

	return deeplink.String(), err
}

func getProductContentId(urlPath string) string {

	if strings.Contains(urlPath, PRODUCT_PAGE_CONSTANT) {
		path := strings.Split(urlPath, PRODUCT_PAGE_CONSTANT)
		return path[1]
	}

	return ""
}

func productUrlIsValid(parameters map[string]string) *models.CustomError {
	if parameters[PRODUCT_KEY_CONTENT_ID] == "" {
		_, err := strconv.Atoi(getProductContentId(parameters["path"]))
		if err != nil {
			return errors.NewCustomErr(MESSAGE_LINK_HAS_NOT_CONTENTID)
		}
	}

	return nil
}

func (*ProductService) ToWebUrl(deeplink string) (string, *models.CustomError) {

	parameters, err := helpers.DeepLinkParse(deeplink)
	if err != nil {
		return "", err
	}

	err = productUrlIsValid(parameters)
	if err != nil {
		return "", err
	}

	values := url.Values{}

	var weburl strings.Builder
	weburl.WriteString(WEBURL_PREFIX)
	weburl.WriteString(PRODUCT_WEBURL_CONTENT_ID_PREFIX)
	weburl.WriteString(parameters[PRODUCT_KEY_CONTENT_ID])

	if parameters[PRODUCT_KEY_CAMPAIGN_ID] != "" {
		values.Add(PRODUCT_KEY_BOUTIQUE_ID, parameters[PRODUCT_KEY_CAMPAIGN_ID])
	}
	if parameters[PRODUCT_KEY_MERCHANT_ID] != "" {
		values.Add(PRODUCT_KEY_MERCHANT_ID, parameters[PRODUCT_KEY_MERCHANT_ID])
	}

	if len(values) > 0 {
		weburl.WriteString(QUESTION_MARK)
	}
	weburl.WriteString(values.Encode())

	return weburl.String(), err
}
