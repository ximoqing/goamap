package goamap

import (
	"io/ioutil"
	"net/http"
)

var Amapapi = "https://restapi.amap.com/"

// Ip ip定位
func IpJson(key, ip string) string {
	url := Amapapi + "v3/ip?ip=" + ip + "&output=json&key=" + key
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	bo := string(body)
	return bo
}

func IpXml(key, ip string) string {
	url := Amapapi + "v3/ip?ip=" + ip + "&output=xml&key=" + key
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	bo := string(body)
	return bo
}

// Ipv2 ip定位v2
func Ipv2(key, typ, ip, sig string) string {
	url := Amapapi + "v5/ip?ip=" + ip + "&output=xml&key=" + key
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	bo := string(body)
	return bo
}
