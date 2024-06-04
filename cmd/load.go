package cmd

import (
	"awesomeProject2/util"
	"encoding/json"
	"log"
	"os"
)

func ParseLoad() {
	load := &util.LoadMetrics{}
	createTime, deleteTime, err := util.GetLoadTestServicesCreateAndDeleteTime(os.Getenv("WATCH_DIR"))
	if err != nil {
		log.Fatal(err)
		return
	}
	load.Junit.CreateTimeStamp = createTime
	load.Junit.DeleteTimeStamp = deleteTime

	callbackInfo := map[string]float64{
		"createTime": createTime,
		"deleteTime": deleteTime,
	}
	jsonData, err := json.Marshal(callbackInfo)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 回调
	if err = sendResult(jsonData); err != nil {
		log.Fatal(err)
		return
	}
}
