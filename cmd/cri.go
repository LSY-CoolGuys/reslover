package cmd

import (
	"awesomeProject2/util"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func ParseCri() {
	// junit file /tmp/result/cri/junit.xml
	data, err := util.GetCRITestBenchmark("/tmp/result/cri/junit.xml")
	if err != nil {
		logrus.Errorf("Get cri benchmark failed: %v", err)
	}
	status := "1"
	i := 0
	for _, v := range data {
		for _, subV := range v {
			if subV != float64(0) {
				i++
			}
		}
	}
	if i < 7 {
		status = "0"
	}
	callBackInfo := map[string]interface{}{"values": data}
	if os.Getenv("CALLBACK_URL") == "" {
		logrus.Errorf("CALLBACK_URL is empty")
		return
	}
	criUrl := os.Getenv("CALLBACK_URL")
	callBackUrl := criUrl + fmt.Sprintf("?status=%s", status)
	jsonData, err := json.Marshal(callBackInfo)
	if err != nil {
		logrus.Errorf("Marshal cri benchmark failed: %v", err)
	}
	if err := callbackBackend(jsonData, callBackUrl); err != nil {
		logrus.Errorf("callback Backend failed: %v", err)
	}
}
