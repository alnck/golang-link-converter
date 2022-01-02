package services

import (
	. "link-converter/constants"
	"link-converter/errors"
	"link-converter/helpers"
	"link-converter/models"
	"strings"
)

const ()

type SearchService struct{}

func NewSearchService() SearchService {
	return SearchService{}
}

func (*SearchService) ToDeepLink(weburl string) (string, *models.CustomError) {

	parameters, err := helpers.WebUrlParse(weburl)
	if err != nil {
		return "", err
	}

	err = searchUrlIsValid(parameters)
	if err != nil {
		return "", err
	}

	var deeplink strings.Builder
	deeplink.WriteString(DEEPLIK_PREFIX)
	deeplink.WriteString(SEARCH_DEEPLINK_PREFIX)
	deeplink.WriteString(parameters[SEARCH_KEY_WEBURL_QUERY])

	return deeplink.String(), err
}

func (*SearchService) ToWebUrl(deeplink string) (string, *models.CustomError) {

	parameters, err := helpers.DeepLinkParse(deeplink)
	if err != nil {
		return "", err
	}

	err = searchUrlIsValid(parameters)
	if err != nil {
		return "", err
	}

	var weburl strings.Builder
	weburl.WriteString(WEBURL_PREFIX)
	weburl.WriteString(SEARCH_WEBURL_PREFIX)

	weburl.WriteString(parameters[SEARCH_KEY_DEEPLINK_QUERY])

	return weburl.String(), err
}

func searchUrlIsValid(parameters map[string]string) *models.CustomError {
	if parameters[SEARCH_KEY_WEBURL_QUERY] == "" && parameters[SEARCH_KEY_DEEPLINK_QUERY] == "" {
		return errors.NewCustomErr(MESSAGE_LINK_HAS_NOT_QUERY_PARAMETERS)
	}

	return nil
}
