package cmd

import (
	"awesomeProject2/util"
	"log"
)

const loadJunitPath = "/tmp/test/load/junit.xml"

func ParseLoad() {
	load := &util.LoadMetrics{}
	createTime, deleteTime, err := util.GetLoadTestServicesCreateAndDeleteTime(loadJunitPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	load.Junit.CreateTimeStamp = createTime
	load.Junit.DeleteTimeStamp = deleteTime
}
