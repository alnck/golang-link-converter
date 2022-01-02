package models

import (
	"strings"

	. "link-converter/constants"
)

type RequestModel struct {
	Link string
}

func (model RequestModel) GetWebUrlResourceType() string {
	if strings.Contains(model.Link, "-p-") {
		return RESOURCE_TYPE_PRODUCT
	} else if strings.Contains(model.Link, "/sr") {
		return RESOURCE_TYPE_SEARCH
	} else {
		return RESOURCE_TYPE_OTHER
	}
}

func (model RequestModel) GetDeepLinkResourceType() string {
	if strings.Contains(model.Link, "Product") {
		return RESOURCE_TYPE_PRODUCT
	} else if strings.Contains(model.Link, "Search") {
		return RESOURCE_TYPE_SEARCH
	} else {
		return RESOURCE_TYPE_OTHER
	}
}
