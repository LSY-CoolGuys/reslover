package cmd

import (
	"awesomeProject2/util"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func ParseLoad() {
	createTime, deleteTime, err := util.GetLoadTestServicesCreateAndDeleteTime(os.Getenv("WATCH_DIR") + "/junit.xml")
	if err != nil {
		logrus.Infof("解析junit文件错误，%v", err)
		return
	}
	var zero float64 = 0
	status := 1
	if createTime == zero && deleteTime == zero {
		status = 0
	}
	overCallBack := map[string]interface{}{
		"createTime": createTime,
		"deleteTime": deleteTime,
	}

	body := map[string]interface{}{
		"values": overCallBack,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
		return
	}
	//
	scale := os.Getenv("LOW_SCALE")
	callBackUrl := url + fmt.Sprintf("?scale=%s", scale)
	if os.Getenv("IS_CALL_BACK") == "true" {
		scale = os.Getenv("HIGH_SCALE")
		callBackUrl = url + fmt.Sprintf("?status=%d&scale=%s", status, scale)
	}

	// 回调
	if err = callbackBackend(jsonData, callBackUrl); err != nil {
		logrus.Infof("回调失败，%v", err)
		return
	}
}
