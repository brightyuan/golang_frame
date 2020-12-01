package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Bu struct {
	Union_id   string `json:"union_id"`
	Media_id   string `json:"media_id"`
	Posid      string `json:"posid"`
	Pack_name  string `json:"pack_name"`
	Website    string `json:"website"`
	Media_type string `json:"media_type"`
	Key_word   string `json:"key_word"`
}

type Device struct {
	Device_id       string `json:"device_id"`
	Os              string `json:"os"`
	Osv             string `json:"osv"`
	Connection_type int    `json:"connection_type"`
	Mac             string `json:"mac"`
	Imei            string `json:"imei"`
	Idfa            string `json:"idfa"`
	Oaid            string `json:"oaid"`
	Androidid       string `json:"androidid"`
}

type Content struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type Body struct {
	Lng         string  `json:"lng"`
	Lat         string  `json:"lat"`
	Ip          string  `json:"ip"`
	Token       string  `json:"token"`
	Bu          Bu      `json:"bu"`
	Device      Device  `json:"device"`
	Content     Content `json:"content"`
	Api_version string  `json:"api_version"`
	Request_id  string  `json:"request_id"`
}

func GetDigest(strBody string) string {
	h := sha256.New()
	h.Write([]byte(strBody))
	return "SHA-256=" + base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Signature(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GetDate() string {
	location, _ := time.LoadLocation("Etc/GMT")
	t := time.Now().In(location).Format(http.TimeFormat)
	return t
}

func ApiRequest(urls string, ak string, sk string, body Body) string {
	//1. headers
	var headers map[string]string /*创建集合 */
	headers = make(map[string]string)
	headers["Content-Type"] = "application/raw"
	date := GetDate()
	headers["Date"] = date

	strBody := GetBodyStr(body)
	digest := GetDigest(strBody)
	headers["Digest"] = digest

	//解析URL
	host, path := parseUrl(urls)

	strToHeader := "host date request-line digest"
	strToSign := "host: " + host + "\ndate: " + date + "\nPOST " + path + " HTTP/1.1\ndigest: " + digest
	signature := Signature(strToSign, sk)

	headers["Authorization"] = fmt.Sprintf("hmac username=\"%s\", algorithm=\"hmac-sha256\", headers=\"%s\", signature=\"%s\"", ak, strToHeader, signature)

	//3.create http reqeust
	req, _ := http.NewRequest(http.MethodPost, urls, strings.NewReader(strBody))
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Host = host

	//4.execute http post
	response, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Print("query cluster failed", err.Error())
	}
	defer response.Body.Close()
	s, _ := ioutil.ReadAll(response.Body)
	return string(s)
}

func parseUrl(urls string) (string, string) {
	u, err := url.Parse(urls)
	if err != nil {
		fmt.Println(err)
	}
	host := u.Host
	path := u.Path
	return host, path
}

func GetBodyStr(body Body) string {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Print(err)
	}
	strBody := string(jsonBody)
	return strBody
}

func main() {
	//测试demo
	url := "http://open.sy.soyoung.com:8000/union/union/getList"
	//url := "https://open.soyoung.com/union/api/feeds/get"
	ak := "zYhNVeuGLbesGRFM"
	sk := "l4XItYeVzvvYpYdZAdDR9ZlRYttwtuK7"
	bu := Bu{
		Union_id:   "1",
		Media_id:   "1",
		Posid:      "test4maidian",
		Website:    "www.baidu.com",
		Pack_name:  "ad.union.package",
		Media_type: "1",
		Key_word:   "玻尿酸",
	}

	device := Device{
		Device_id:       "123123test123123",
		Os:              "iOS",
		Osv:             "10.3.1",
		Connection_type: 1,
		Mac:             "00-16-EA-AE-3C-40",
		Imei:            "123123test123123",
		Idfa:            "123123test123123",
		Oaid:            "123123test123123",
		Androidid:       "123123test123123",
	}

	content := Content{
		Page:  1,
		Limit: 5,
	}

	body := Body{
		Lng:         "116.23128",
		Lat:         "40.22077",
		Ip:          "36.112.75.34",
		Token:       "ad95e6983a197eb448de26d1f5db60b6",
		Bu:          bu,
		Device:      device,
		Content:     content,
		Api_version: "1.0",
		Request_id:  "",
	}

	resp := ApiRequest(url, ak, sk, body)
	fmt.Println(resp)
}
