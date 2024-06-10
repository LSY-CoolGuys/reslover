package main

import (
	"awesomeProject2/cmd"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var junitFile string

func init() {
	junitFile = "/tmp/result/%s/junit.xml"
}

func main() {
	kind := os.Getenv("TEST_MODE")
	monitorFile := fmt.Sprintf(junitFile, kind)
	logrus.WithFields(logrus.Fields{
		"waitfile": monitorFile,
	}).Info("开始监测junit.xml文件")
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	//if err := cmd.CallbackPodName(); err != nil {
	//logrus.Fatal(err)
	//}
	for {
		select {
		case <-ticker.C:
			if _, err := os.Stat(monitorFile); err != nil {
				if !os.IsNotExist(err) {
					logrus.Tracef("监测到junit.xml文件生成，但发生错误:%s", err)
				}
				continue
			}
			logrus.Info("junit.xml文件生成，开始执行回调函数")
			cmd.Parse()
			logrus.Info("解析结束，容器关闭")
		}
	}
}
