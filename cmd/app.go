package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

var testMode string
var podName string
var url string

func init() {
	testMode = os.Getenv("TEST_MODE")
	podName = os.Getenv("POD_NAME")
	url = os.Getenv("CALLBACK_URL")
}

func Parse() {
	switch strings.ToLower(os.Getenv("TEST_MODE")) {
	case "density":
		ParseDestiny()
	case "load":
		ParseLoad()
	}
}

// TODO: callBackurl
func callbackBackend(jsonData []byte, callBackUrl string) error {
	//callbackUrl := os.Getenv("CALLBACK_URL")
	//state := os.Getenv("TEST_STATE")
	//if state == "" || state == "failed" {
	//	state = "0"
	//} else if state == "SUCCESS" {
	//	state = "1"
	//}
	body := map[string][]byte{"values": jsonData}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}
	log.Printf("callback url is : %s,body is : %s", callBackUrl, string(jsonData))
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI(callBackUrl)
	req.Header.Set("Authorization", "Internal testing")
	req.Header.SetContentType("application/json")
	//req.SetBodyString(string(jsonData))
	req.SetBody(bodyJson)
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
