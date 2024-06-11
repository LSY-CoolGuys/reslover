package cmd

import (
	"awesomeProject2/util"
	"encoding/json"
	"log"
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
		log.Fatal(err)
		return
	}
	// status现在无法获取 默认为1吧
	callBackUrl := url + "?/status=1"
	if err = callbackBackend(jsonData, callBackUrl); err != nil {
		log.Fatal(err)
		return
	}
}
