package contracts

type Converter interface {
	ToDeepLink(url string) string
	ToWebUrl(deeplink string) string
}
