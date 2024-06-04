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
	APIResponsiveness                    PerfData `json:"apiResponsiveness,omitempty"`
	PodStartupLatency                    PerfData `json:"podStartupLatency"`
	StatelessPodStartupLatency           PerfData `json:"statelessPodStartupLatency,omitempty"`
	StatelessSaturationPodStartupLatency PerfData `json:"statelessSaturationPodStartupLatency,omitempty"`
	SaturationPodStartupLatency          PerfData `json:"saturationPodStartupLatency"`
	SchedulingThroughput                 PerfData `json:"schedulingThroughput"`
	SchedulingMetrics                    PerfData `json:"schedulingMetrics,omitempty"`
}

type SchedulerThroughput struct {
	Perc50 float64 `json:"perc50"`
	Perc90 float64 `json:"perc90"`
	Perc99 float64 `json:"perc99"`
	Max    float64 `json:"max"`
}

// ServiceMetrics 包含创建和删除service的时间
type ServiceMetrics struct {
	CreateTimeStamp float64 `json:"create_time_stamp"`
	DeleteTimeStamp float64 `json:"delete_time_stamp"`
}

// LoadMetrics 解析load junit文件
type LoadMetrics struct {
	SchedulerThroughputParallel PerfData
	SequenceParallelism         PerfData
	Junit                       ServiceMetrics
}

// DestinyOutputPath 所有指标的输出路径
var DestinyOutputPath = map[string]string{
	"APIResponsiveness":                    "/tmp/test/destiny/APIResponsivenessPrometheus_density.json", //"/tmp/result/destiny/APIResponsivenessPrometheus_simple_density.json",
	"PodStartupLatency":                    "/tmp/result/destiny/PodStartupLatency_PodStartupLatency_density.json",
	"StatelessPodStartupLatency":           "/tmp/result/destiny/StatelessPodStartupLatency_PodStartupLatency_density.json",
	"StatelessSaturationPodStartupLatency": "/tmp/result/destiny/StatelessPodStartupLatency_SaturationPodStartupLatency_density.json",
	"SaturationPodStartupLatency":          "/tmp/result/destiny/PodStartupLatency_SaturationPodStartupLatency_density.json",
	"SchedulingThroughput":                 "/tmp/result/destiny/SchedulingThroughput_density.json",
}

// LoadOutputPath 所有指标的输出路径
var LoadOutputPath = []string{
	"/tmp/test/load/junit.txt",
}
