package proxypool

import (
	"net/http"
	"net/url"
)

// Your agent server
const ProxyServer = ""

// Your proxy info
const ProxyUser = ""
const ProxyPass = ""

type AbuyunProxy struct {
	AppID     string
	AppSecret string
}

func (p AbuyunProxy) ProxyClient() http.Client {
	proxyUrl, _ := url.Parse("http://" + p.AppID + ":" + p.AppSecret + "@" + ProxyServer)
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}
