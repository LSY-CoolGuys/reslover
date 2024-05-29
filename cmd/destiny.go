package cmd

import "awesomeProject2/util"

func ParseDestiny() {
	destiny := &util.DestinyMetrics{}
	util.GetOutputJson(destiny)
}
