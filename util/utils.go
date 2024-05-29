package util

// DataItem is the data point.
type DataItem struct {
	// Data is a map from bucket to real data point (e.g. "Perc90" -> 23.5). Notice
	// that all data items with the same label combination should have the same buckets.
	Data map[string]float64 `json:"data"`
	// Unit is the data unit. Notice that all data items with the same label combination
	// should have the same unit.
	Unit string `json:"unit"`
	// Labels is the labels of the data item.
	Labels map[string]string `json:"labels,omitempty"`
}

// PerfData contains all data items generated in current test.
type PerfData struct {
	// Version is the version of the metrics. The metrics consumer could use the version
	// to detect metrics version change and decide what version to support.
	Version   string     `json:"version"`
	DataItems []DataItem `json:"dataItems"`
	// Labels is the labels of the dataset.
	Labels map[string]string `json:"labels,omitempty"`
}

// DestinyMetrics  包含4种
// APIResponsiveness、PodStartupLatency、SaturationPodStartupLatency、SchedulingThroughput.
type DestinyMetrics struct {
	APIResponsiveness           PerfData
	PodStartupLatency           PerfData
	SaturationPodStartupLatency PerfData
	SchedulingThroughput        PerfData
	SchedulingMetrics           PerfData
}

// ServiceMetrics 包含创建和删除service的时间
type ServiceMetrics struct {
	CreateTimeStamp float64
	DeleteTimeStamp float64
}

// LoadMetrics 解析load junit文件
type LoadMetrics struct {
	SchedulerThroughputParallel PerfData
	SequenceParallelism         PerfData
	Junit                       ServiceMetrics
}

// DestinyOutputPath 所有指标的输出路径
var DestinyOutputPath = map[string]string{
	"APIResponsiveness":           "/tmp/test/destiny/ar.txt",
	"PodStartupLatency":           "/tmp/test/destiny/psl.txt",
	"SaturationPodStartupLatency": "/tmp/test/destiny/spsl.txt",
	"SchedulingThroughput":        "/tmp/test/destiny/st.txt",
}

// LoadOutputPath 所有指标的输出路径
var LoadOutputPath = []string{
	"/tmp/test/load/junit.txt",
	"/tmp/test/load/stp.txt",
	"/tmp/test/load/sp.txt",
}
