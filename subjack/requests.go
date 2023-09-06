package subjack

import (
	"crypto/tls"
	"time"

	"github.com/valyala/fasthttp"
)

func get(url string, ssl bool, timeout int) (body []byte) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(site(url, ssl))
	req.Header.Add("Connection", "close")
	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		}}

	client.DoTimeout(req, resp, time.Duration(timeout)*time.Second)

	return resp.Body()
}

func site(url string, ssl bool) (site string) {
	site = "http://" + url
	if ssl {
		site = "https://" + url
	}

	return site
}
