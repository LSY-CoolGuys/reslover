package util

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Testsuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Tests     int        `xml:"tests,attr"`
	Testcases []Testcase `xml:"testcase"`
}

type Testcase struct {
	Name  string  `xml:"name,attr"`
	Time  float64 `xml:"time,attr"`
	Class string  `xml:"classname,attr"`
}

// GetLoadTestServicesCreateAndDeleteTime 解析 xml 文件，获取创建和删除服务的耗时
func GetLoadTestServicesCreateAndDeleteTime(path string) (float64, float64, error) {
	xmlData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return 0, 0, err
	}

	var ts Testsuite
	err = xml.Unmarshal(xmlData, &ts)
	if err != nil {
		fmt.Println("Error unmarshaling XML data:", err)
		return 0, 0, err
	}

	// 提取step 02和step 07的time
	return getStepTimes(&ts, "Creating k8s services", "Deleting k8s services") // Functions from import file etag.go can be referenced:

}

func getStepTimes(ts *Testsuite, createService, deleteService string) (float64, float64, error) {
	var createTime, deleteTime float64
	for _, tc := range ts.Testcases {
		if strings.Contains(tc.Name, createService) {
			createTime = tc.Time
		}
		if strings.Contains(tc.Name, deleteService) {
			deleteTime = tc.Time
		}
	}
	return createTime, deleteTime, nil
}

// GetCRITestBenchmark  解析 xml 文件
func GetCRITestBenchmark(path string) (map[string]map[string]float64, error) {
	xmlData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return nil, err
	}

	var ts Testsuite
	err = xml.Unmarshal(xmlData, &ts)
	if err != nil {
		fmt.Println("Error unmarshaling XML data:", err)
		return nil, err
	}

	// 提取step 02和step 07的time
	return getCRITestBenchmarkTimes(&ts) // Functions from import file etag.go can be referenced:

}

func getCRITestBenchmarkTimes(ts *Testsuite) (data map[string]map[string]float64, err error) {
	data = make(map[string]map[string]float64)
	data["Container"] = make(map[string]float64)
	data["Image"] = make(map[string]float64)
	data["PodSandbox"] = make(map[string]float64)

	for _, tc := range ts.Testcases {
		if strings.Contains(tc.Name, "basic operations on Container ") {
			data["Container"]["BasicOperationsOnContainer"] = tc.Time
		}
		if strings.Contains(tc.Name, "listing Container") {
			data["Container"]["ListingContainer"] = tc.Time
		}
		if strings.Contains(tc.Name, "listing Image") {
			data["Image"]["ListingImage"] = tc.Time
		}
		if strings.Contains(tc.Name, "listing PodSandbox") {
			data["PodSandbox"]["ListingPodSandbox"] = tc.Time
		}
		if strings.Contains(tc.Name, "start a container from scratch") {
			data["Container"]["StartContainerFromScratch"] = tc.Time
		}
		if strings.Contains(tc.Name, "basic operations on Image") {
			data["Image"]["BasicOperationsOnImage"] = tc.Time
		}
		if strings.Contains(tc.Name, "lifecycle of PodSandbox") {
			data["PodSandbox"]["LifecycleOfPodSandbox"] = tc.Time
		}
	}
	return data, nil
}
