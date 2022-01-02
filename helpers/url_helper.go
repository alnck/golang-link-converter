package helpers

import (
	. "link-converter/constants"
	"link-converter/errors"
	"link-converter/models"
	"log"
	"net/url"
	"strings"
)

func LinkIsValid(link string) *models.CustomError {
	if link == "" {
		return errors.NewCustomErr(MESSAGE_LINK_IS_EMPTY)
	}

	if !strings.HasPrefix(link, WEBURL_PREFIX) && !strings.HasPrefix(link, DEEPLIK_PREFIX) {
		return errors.NewCustomErr(MESSAGE_LINK_NOT_CONVERSION)
	}

	return nil
}

func WebUrlParse(weburl string) (map[string]string, *models.CustomError) {
	if err := LinkIsValid(weburl); err != nil {
		return nil, err
	}

	values := make(map[string]string)

	parseurl, err := url.Parse(weburl)
	if err != nil {
		return nil, errors.NewCustomErr(err.Error())
	}

	values["path"] = parseurl.Path

	parameters, err := url.ParseQuery(parseurl.RawQuery)
	if err != nil {
		return nil, errors.NewCustomErr(err.Error())
	}

	for key, value := range parameters {
		values[strings.ToLower(key)] = url.QueryEscape(value[0])
	}

	return values, nil
}

func DeepLinkParse(deeplink string) (map[string]string, *models.CustomError) {
	if err := LinkIsValid(deeplink); err != nil {
		return nil, err
	}

	parseDeep, err := url.Parse(deeplink)
	if err != nil {
		log.Fatal(err)
		return nil, errors.NewCustomErr(err.Error())
	}

	parameters, err := url.ParseQuery(parseDeep.RawQuery)
	if err != nil {
		log.Fatal(err)
		return nil, errors.NewCustomErr(err.Error())
	}

	values := make(map[string]string)

	for key, value := range parameters {
		values[strings.ToLower(key)] = url.QueryEscape(value[0])
	}

	return values, nil
}
