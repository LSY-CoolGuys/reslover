package cmd

import (
	"awesomeProject2/util"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func ParseDestiny() {
	destiny := &util.DestinyMetrics{}
	if err := util.GetOutputJson(destiny); err != nil {
		log.Fatal(err)
	}

	callBackInfo := map[string]interface{}{
		"PodStartupLatency":           destiny.PodStartupLatency,
		"SaturationPodStartupLatency": destiny.SaturationPodStartupLatency,
		"SchedulingThroughput":        destiny.SchedulingThroughput,
	}
	body := map[string]interface{}{"values": callBackInfo}
	//bodyJson, err := json.Marshal(body)
	//if err != nil {
	//	return err
	//}
	jsonData, err := json.Marshal(body)
	if err != nil {
		logrus.Errorf("解析错误%s", err)
		return
	}
	status := "?status=1"
	_, err = os.Stat("/tmp/result/status.yaml")
	if err != nil && os.IsNotExist(err) {
		status = "?status=0"
	}
	// status现在无法获取 默认为1吧
	if err = callbackBackend(jsonData, url+status); err != nil {
		log.Fatal(err)
		return
	}
}
