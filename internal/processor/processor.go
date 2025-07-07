package processor

import (
	"github.com/JaveleyQAQ/geodns/internal/types"
)

// DNSProcessor DNS结果处理器
type DNSProcessor struct {
	results []types.RegionResult
}

// Reset 清空历史结果
func (dp *DNSProcessor) Reset() {
	dp.results = make([]types.RegionResult, 0)
}

// NewDNSProcessor 创建新的DNS处理器
func NewDNSProcessor() *DNSProcessor {
	return &DNSProcessor{
		results: make([]types.RegionResult, 0),
	}
}

// ProcessResults 处理查询结果
func (dp *DNSProcessor) ProcessResults(resultChan <-chan types.RegionResult) {
	for res := range resultChan {
		dp.results = append(dp.results, res)
	}
}

// GetSummary 获取结果汇总
func (dp *DNSProcessor) GetSummary(domain string) types.ResultSummary {
	uniqueAnswers := make(map[string]map[string]bool)

	for _, res := range dp.results {
		// 只处理指定域名的结果
		if res.Domain == domain {
			for _, answer := range res.Answers {
				if uniqueAnswers[answer.Type] == nil {
					uniqueAnswers[answer.Type] = make(map[string]bool)
				}
				uniqueAnswers[answer.Type][answer.Value] = true
			}
		}
	}

	// 转换为字符串切片
	uniqueAnswersSlice := make(map[string][]string)
	for recordType, values := range uniqueAnswers {
		uniqueAnswersSlice[recordType] = types.Keys(values)
	}

	return types.ResultSummary{
		Domain:        domain,
		Results:       dp.results,
		UniqueAnswers: uniqueAnswersSlice,
	}
}
