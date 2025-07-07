package types

import "sort"

// DNS记录类型
// type DNSRecordType struct {
// 	Name string
// 	Type uint16
// }

// DNSAnswer DNS答案
type DNSAnswer struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// RegionResult 区域查询结果
type RegionResult struct {
	Domain   string      `json:"domain"`
	Region   string      `json:"region"`
	Answers  []DNSAnswer `json:"answers"`
	IPs      []string    `json:"-"`
	CNAMEs   []string    `json:"-"`
	RawBytes []byte      `json:"-"`
	Error    string      `json:"error,omitempty"`
}

// ResultSummary 结果汇总
type ResultSummary struct {
	Domain        string              `json:"domain"`
	Results       []RegionResult      `json:"results"`
	UniqueAnswers map[string][]string `json:"unique_answers"`
}

// Keys 获取map的键并排序
func Keys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
