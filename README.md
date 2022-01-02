# Trenyol Link Converter


Başkalarının Trendyol.com bağlantılarını mobil ve web uygulamaları arasında dönüştürmesine yardımcı olan bir web hizmeti uygulamasıdır. Web uygulamaları URL'leri, mobil uygulamalar ise derin bağlantıları kullanır. Her iki uygulama da, uygulamaların içindeki belirli konumları yeniden yönlendirmek için bağlantıları kullanır. Uygulamalar arasında yönlendirme yapmak istediğinizde, URL'leri derin bağlantılara veya derin bağlantıları URL'lere dönüştürmelisiniz.

### Çalıştırma (Running)
---
`docker compose up`


### Teknik Detaylar
---
- Database olarak `redis` kullanılmıştır.
- Sistemde yoğunlugu azalmak için redis üzerinden `cache` oluşturulmuştur.

### Kullanım
---
Sistem 2 Endpoint üzerinden çalışır.

`POST /converter/toweblink  # Web Url to Deep Link` <br> 
`POST /converter/todeeplink # Deep Link to Web URL`



##### Request ToDeepLink Example
`   curl -X POST http://localhost:8080/converter/todeeplink
   -H 'Content-Type: application/json'
   -d '{"Link": "https://www.trendyol.com/sr?q=elbise"}' 
   `
###### Response

```
{
    "ConvertedLink": "ty://?Page=Search&Query=elbise"
} 
```

##### Request ToWebURL Example
`   curl -X POST http://localhost:8080/converter/toweburl
   -H 'Content-Type: application/json'
   -d '{"Link": "ty://?Page=Product&ContentId=1925865"}'
   `
###### Response
```
{
    "ConvertedLink": "https://www.trendyol.com/brand/name-p-1925865"
} 
```
