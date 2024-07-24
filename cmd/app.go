package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
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
	case "cri":
		ParseCri()
	}
}

func DirectCallBack() {
	body := map[string]interface{}{"values": "0"}
	jsonData, err := json.Marshal(body)
	if err != nil {
		logrus.Errorf("解析错误%s", err)
		return
	}

	Url := os.Getenv("CALLBACK_URL")
	callBackUrl := Url + fmt.Sprintf("?status=%d", 0)
	if err := callbackBackend(jsonData, callBackUrl); err != nil {
		logrus.Errorf("回调错误%s", err)
	}
}

// TODO: callBackurl
func callbackBackend(jsonData []byte, callBackUrl string) error {
	log.Printf("callback url is : %s,body is : %s", callBackUrl, string(jsonData))
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI(callBackUrl)
	req.Header.Set("Authorization", "Internal testing")
	req.Header.SetContentType("application/json")
	req.SetBody(jsonData)
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
