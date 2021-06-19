package util

import (
	"io/ioutil"
	"net/http"
)

func GetPublicIP() (string, error) {
	ipService := GetEnv("IP_SERVICE", "http://139.9.170.194:8848/ip")

	resp, err := http.Get(ipService)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
