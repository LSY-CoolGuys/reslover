package util

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetLoadTestServicesCreateAndDeleteTime(t *testing.T) {
	xmlContent := `  
<testsuite name="ClusterLoaderV2" tests="0" failures="0" errors="0" time="12.135">  
    <testcase name="load: [step: 02] Creating k8s services" classname="ClusterLoaderV2" time="0.257525709"></testcase>  
    <testcase name="load: [step: 07] Deleting k8s services" classname="ClusterLoaderV2" time="0.359431966"></testcase>  
    <!-- 其他 testcase ... -->  
</testsuite>  
`

	// 使用 bytes.Buffer 模拟文件读取
	xmlBuffer := bytes.NewBufferString(xmlContent)

	// 使用自定义的ReadFile函数，因为我们不想实际读取文件
	readFile := func(path string) ([]byte, error) {
		return ioutil.ReadAll(xmlBuffer)
	}

	// 调用 GetLoadTestServicesCreateAndDeleteTime 函数进行测试
	step02Time, step07Time, err := GetLoadTestServicesCreateAndDeleteTimeWithReadFile(readFile, "dummy_path.xml")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 验证返回的时间值
	expectedStep02Time := 0.257525709
	expectedStep07Time := 0.359431966
	if step02Time != expectedStep02Time {
		t.Errorf("Expected step 02 time to be %.6f, got %.6f", expectedStep02Time, step02Time)
	}
	if step07Time != expectedStep07Time {
		t.Errorf("Expected step 07 time to be %.6f, got %.6f", expectedStep07Time, step07Time)
	}
}

func GetLoadTestServicesCreateAndDeleteTimeWithReadFile(readFile func(string) ([]byte, error), path string) (float64, float64, error) {
	xmlData, err := readFile(path)
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

	return getStepTimes(&ts, "load: [step: 02]", "load: [step: 07]")
}
