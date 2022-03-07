package goamap

import (
	"io/ioutil"
	"net/http"
	"net/url"
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
func Ipv2(key, typ, ip string) string {
	ur := Amapapi + "v5/ip?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("type", typ)
	params.Add("ip", ip)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

// 地理/逆地理编码
func Georegeo(key, address string) string {
	ur := Amapapi + "v3/geocode/geo?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("address", address)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

// 路径规划
func Direction(key, origin, destination string) string {
	ur := Amapapi + "v3/direction/walking?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("origin", origin)
	params.Add("destination", destination)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

// 路径规划v2
func DirectionV2(key, origin, destination string) string {
	ur := Amapapi + "v5/direction/driving?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("origin", origin)
	params.Add("destination", destination)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

// 行政区域查询
func District(key, keywords string) string {
	ur := Amapapi + "v3/config/district?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("keywords", keywords)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

// 天气查询
func Weatherinfo(key, city string) string {
	ur := Amapapi + "v3/weather/weatherInfo?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("city", city)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

// 输入提示
func Inputtips(key, keywords string) string {
	ur := Amapapi + "v3/assistant/inputtips?parameters"
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("key", key)
	params.Add("keywords", keywords)
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(con)
}

//
