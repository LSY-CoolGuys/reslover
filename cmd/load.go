package cmd

import (
	"awesomeProject2/util"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ParseLoad() {
	createTime, deleteTime, err := util.GetLoadTestServicesCreateAndDeleteTime(os.Getenv("WATCH_DIR"))
	if err != nil {
		log.Fatal(err)
		return
	}

	if os.Getenv("IS_CALL_BACK") == "true" {
		fs, err := os.ReadFile("/var/load/firstData")
		if err != nil {
			log.Fatal(err)
			return
		}
		overCallBack := map[string]interface{}{
			os.Getenv("LOW_SCALE"): string(fs),
			os.Getenv("HIGH_SCALE"): map[string]float64{
				"createTime": createTime,
				"deleteTime": deleteTime,
			},
		}
		jsonData, err := json.Marshal(overCallBack)
		if err != nil {
			log.Fatal(err)
			return
		}
		//
		status := 1
		if createTime == 0 && deleteTime == 0 {
			status = 0
		}
		callBackUrl := url + fmt.Sprintf("?status=%d", status)
		// 回调
		if err = callbackBackend(jsonData, callBackUrl); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		callbackInfo := map[string]float64{
			"createTime": createTime,
			"deleteTime": deleteTime,
		}
		jsonData, err := json.Marshal(callbackInfo)
		if err != nil {
			log.Fatal(err)
			return
		}
		fs, err := os.OpenFile("/var/load/firstData", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fs.Close()

		n, err := fs.Write(jsonData)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("成功写入%d字节\n", n)
	}
}
