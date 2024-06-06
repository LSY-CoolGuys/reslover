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

	jsonData, err := json.Marshal(callBackInfo)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = sendResult(jsonData); err != nil {
		log.Fatal(err)
		return
	}
}
