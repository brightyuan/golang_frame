package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestApiRequest(t *testing.T) {

}

/**
Wed, 04 Nov 2020 06:16:27 GMT
*/
func TestGetDate(t *testing.T) {
	date := GetDate()
	fmt.Println(date)
}

func TestHmac_encode_body(t *testing.T) {

}

func TestSignature(t *testing.T) {

}

func TestString(t *testing.T) {
}

func TestPostMan(t *testing.T) {
	body := "{\n            \"lng\":\"116.23128\",\n            \"lat\":\"40.22077\",\n            \"ip\":\"36.112.75.34\",\n            \"token\":\"ad95e6983a197eb448de26d1f5db60b6\",\n            \"bu\":{\n                \"union_id\":1,\n                \"media_id\":1,\n                \"posid\":\"test4maidian\",\n                \"website\":\"www.baidu.com\",\n                \"pack_name\":\"ad.union.package\",\n                \"media_type\":1,\n                \"key_word\":\"玻尿酸\"\n            },\n            \"device\":{\n                \"device_id\":\"123123test123123\",\n                \"os\":\"iOS\",\n                \"osv\":\"10.3.1\",\n                \"connection_type\":1,\n                \"mac\":\"00-16-EA-AE-3C-40\",\n                \"imei\":\"123123test123123\",\n                \"idfa\":\"123123test123123\",\n                \"oaid\":\"123123test123123\",\n                \"androidid\":\"123123test123123\"\n            },\n            \"content\":{\n                \"page\":1,\n                \"limit\":5\n            },\n            \"api_version\":\"1.0\",\n            \"request_id\":\"\"\n        }"

	sk := "l4XItYeVzvvYpYdZAdDR9ZlRYttwtuK7"
	host := "open.sy.soyoung.com"
	path := "/union/union/getList"

	dates := GetDate()
	digest := GetDigest(body)
	strToSign := "host: " + host + "\ndate: " + dates + "\nPOST " + path + " HTTP/1.1\ndigest: " + digest
	signature := Signature(strToSign, sk)
	signature = "\"" + signature + "\""
	ak := "\"zYhNVeuGLbesGRFM\""

	url := "http://open.sy.soyoung.com:8000/union/union/getList"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.Host = "open.sy.soyoung.com"

	req.Header.Set("Content-Type", "application/raw")

	req.Header.Set("Date", dates)
	req.Header.Set("Digest", digest)
	authStr := "hmac username=" + ak + ", algorithm=\"hmac-sha256\", headers=\"host date request-line digest\", signature=" + signature
	req.Header.Set("Authorization", string(authStr))

	resp, _ := (&http.Client{}).Do(req)
	s, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}

func TestUrl(t *testing.T) {
	s := "http://open.sy.soyoung.com:8000/union/union/getList"
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.Host)
	fmt.Println(u.Port())

}
