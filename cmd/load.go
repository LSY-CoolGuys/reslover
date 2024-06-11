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

	if os.Getenv("IS_CALL_BACK") == "true" {
		fs, err := os.ReadFile("/var/load/firstData")
		if err != nil {
			logrus.Infof("获取第一次结果错误，%v", err)
			return
		}
		overCallBack := map[string]interface{}{
			os.Getenv("LOW_SCALE"): string(fs),
			os.Getenv("HIGH_SCALE"): map[string]float64{
				"createTime": createTime,
				"deleteTime": deleteTime,
			},
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
		var zero float64 = 0
		status := 1
		if createTime == zero && deleteTime == zero {
			status = 0
		}
		callBackUrl := url + fmt.Sprintf("?status=%d", status)
		// 回调
		if err = callbackBackend(jsonData, callBackUrl); err != nil {
			logrus.Infof("回调失败，%v", err)
			return
		}
	} else {
		callbackInfo := map[string]float64{
			"createTime": createTime,
			"deleteTime": deleteTime,
		}
		jsonData, err := json.Marshal(callbackInfo)
		if err != nil {
			logrus.Infof("序列化第一次测试结果失败，%v", err)
			return
		}
		fs, err := os.Open("/var/load/firstData")
		if err != nil {
			logrus.Infof("打开configmap挂载文件失败，err=%v", err)
			return
		}
		defer fs.Close()

		n, err := fs.Write(jsonData)
		if err != nil {
			logrus.Infof("写入文件失败，err=%v", err)
			return
		}
		logrus.Infof("成功写入%d字节", n)
	}
}
