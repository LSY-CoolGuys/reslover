package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetOutputJson(dm *DestinyMetrics) error {
	for k, v := range DestinyOutputPath {
		switch k {
		case "PodStartupLatency":
			err := parseJson(&dm.PodStartupLatency, v)
			if err != nil {
				return err
			}
		case "SaturationPodStartupLatency":
			err := parseJson(&dm.SaturationPodStartupLatency, v)
			if err != nil {
				return err
			}
		case "SchedulingThroughput":
			err := parseSch(&dm.SchedulingThroughput, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parseJson(dm *PerfData, path string) error {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return err
	}
	err = json.Unmarshal(jsonData, dm)
	if err != nil {
		fmt.Println("Error unmarshalling JSON file:", err)
		return err
	}
	return nil
}

func parseSch(sc *SchedulerThroughput, path string) error {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return err
	}
	err = json.Unmarshal(jsonData, sc)
	if err != nil {
		fmt.Println("Error unmarshalling JSON file:", err)
		return err
	}
	return nil
}
