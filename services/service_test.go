package services

import (
	. "link-converter/constants"
	"link-converter/models"
	"strings"
	"testing"
)

var successToWebUrlList = []struct {
	link          string
	convertedLink string
}{
	{"ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064", "https://www.mysite.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064"},
	{"ty://?Page=Product&ContentId=1925865", "https://www.mysite.com/brand/name-p-1925865"},
	{"ty://?Page=Search&Query=elbise", "https://www.mysite.com/sr?q=elbise"},
	{"ty://?Page=Home", "https://www.mysite.com/"},
}

func TestSuccessToWebUrlReturnConvertedLink(t *testing.T) {
	converterService := NewConverterService(nil)

	for _, _data := range successToWebUrlList {

		requestModel := models.RequestModel{Link: _data.link}

		response, err := converterService.ToWebUrl(requestModel)
		if err != nil || response.ConvertedLink != strings.ToLower(_data.convertedLink) {
			t.Errorf("Sonuçlar eşleşmedi: Link: %v, ConvertedLink: %v", _data.link, response.ConvertedLink)
		}
	}
}

//Query Parametresi Yoksa
func TestReturnNilIfNoQueryParameter(t *testing.T) {
	converterService := NewConverterService(nil)

	requestModel := models.RequestModel{Link: "https://www.mysite.com/sr?=elbise"}

	_, errDeepLink := converterService.ToDeepLink(requestModel)
	if errDeepLink == nil || errDeepLink.Message != MESSAGE_LINK_HAS_NOT_QUERY_PARAMETERS {
		t.Errorf("ToDeepLınk Query Parameter empty FAIL")
	}

	requestModel = models.RequestModel{Link: "ty://?Page=Search&=%C3%BCt%C3%BC"}

	_, errWeb := converterService.ToWebUrl(requestModel)
	if errWeb == nil || errWeb.Message != MESSAGE_LINK_HAS_NOT_QUERY_PARAMETERS {
		t.Errorf("ToWebUrl Query Parameter empty FAIL")
	}

}

//Contentid Yoksa Test Sonucu
func TestLinkHasNotContentid(t *testing.T) {
	converterService := NewConverterService(nil)

	requestModel := models.RequestModel{Link: "ty://?Page=Product&ContentId=&MerchantId=105064"}

	_, errWeb := converterService.ToWebUrl(requestModel)
	if errWeb == nil || errWeb.Message != MESSAGE_LINK_HAS_NOT_CONTENTID {
		t.Errorf("ToWebUrl Has Not ContentId empty FAIL")
	}

	requestModel = models.RequestModel{Link: "https://www.mysite.com/casio/erkek-kol-saati-p-?merchantId=105064"}

	_, errDeepLink := converterService.ToDeepLink(requestModel)
	if errWeb == nil || errDeepLink.Message != MESSAGE_LINK_HAS_NOT_CONTENTID {
		t.Errorf("ToDeepLink Has Not ContentId empty FAIL")
	}
}
