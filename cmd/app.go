package cmd

import (
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
)

func Parse() {
	switch strings.ToLower(os.Getenv("TEST_MODE")) {
	case "density":
		ParseDestiny()
	case "load":
		ParseLoad()
	}
}

func sendResult(jsonData []byte) error {
	callbackUrl := os.Getenv("CALLBACK_URL")
	log.Printf("callback url is : %s,body is : %s", callbackUrl, string(jsonData))
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI(callbackUrl)
	req.Header.Set("Authorization", "Internal install cluster callback")
	req.Header.SetContentType("application/json")
	req.SetBodyString(string(jsonData))
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	client := fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("回调函数返回值：%s", string(resp.Body()))
	return nil
}
