package helpers

import "testing"

//Link Hatalıysa
func TestReturnErrorWhenPrefixCorrupt(t *testing.T) {
	err := LinkIsValid("ty/?Page=Product&ContentId=151515&MerchantId=105064")
	if err == nil {
		t.Errorf("Error: Deeplink Prefix Corrupt")
	}

	err = LinkIsValid("https://www.mysite.com/sr?q=elbise")
	if err == nil {
		t.Errorf("Error: WebUrl Prefix Corrupt")
	}
}

//link Boşsa
func TestReturnErrorWhenLinkIsEmpty(t *testing.T) {
	err := LinkIsValid("")
	if err == nil {
		t.Errorf("Error: Link is empty")
	}
}
