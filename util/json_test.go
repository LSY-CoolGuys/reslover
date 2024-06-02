package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// 测试parseJson函数
func TestParseJson(t *testing.T) {
	// 创建一个临时的JSON文件用于测试
	tempJSONFile := filepath.Join(os.TempDir(), "/tmp/test/destiny/APIResponsivenessPrometheus_density.json")
	defer os.Remove(tempJSONFile)

	// 创建一个PerfData实例并写入JSON文件
	pd := PerfData{
		// 根据实际情况填充字段
		Version: "1.0",
		DataItems: []DataItem{
			// 根据实际情况填充字段
			{
				Data: map[string]float64{
					"P99": 0.000998,
				},
				Unit: "ms",
				Labels: map[string]string{
					"k8s-app": "ww",
				},
			},
		},
	}

	jsonData, err := json.Marshal(pd)
	if err != nil {
		t.Fatalf("Failed to marshal PerfData: %v", err)
	}
	err = os.WriteFile(tempJSONFile, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write JSON file: %v", err)
	}

	// 调用parseJson函数进行测试
	err = parseJson(&pd, []string{tempJSONFile})
	if err != nil {
		t.Errorf("parseJson failed: %v", err)
	}

	fmt.Printf("%+v\n", pd)
	// 这里可以添加更多的断言来检查pd的属性是否符合预期
}
