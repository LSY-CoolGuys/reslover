package cmd

import (
	"fmt"
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
	state := os.Getenv("TEST_STATE")
	if state == "" || state == "failed" {
		state = "0"
	} else if state == "SUCCESS" {
		state = "1"
	}
	log.Printf("callback url is : %s,body is : %s", callbackUrl, string(jsonData))
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI(callbackUrl + fmt.Sprintf("?status=%s", state))
	req.Header.Set("Authorization", "Internal testing")
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
