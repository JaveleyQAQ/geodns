package formatter

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/JaveleyQAQ/geodns/internal/config"
	"github.com/JaveleyQAQ/geodns/internal/types"
)

type OutputFormatter struct {
	outputWriter *os.File
	jsonOutput   bool
	ResponseOnly bool
	showResponse bool
	RecordTypes  []uint16 // 记录类型过滤
	colorful     bool     // 是否彩色输出
}

// ... existing code ...
func NewOutputFormatter(jsonOutput, responseOnly, showResponse bool, recordTypes []uint16, outputFile string) *OutputFormatter {
	var writer *os.File
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err == nil {
			writer = f
		} else {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		}
	}
	colorful := !jsonOutput && !responseOnly && outputFile == ""
	return &OutputFormatter{
		outputWriter: writer,
		jsonOutput:   jsonOutput,
		ResponseOnly: responseOnly,
		showResponse: showResponse,
		RecordTypes:  recordTypes,
		colorful:     colorful,
	}
}

func (of *OutputFormatter) writeOutput(s string) {
	if of.outputWriter != nil {
		of.outputWriter.WriteString(s)
	} else {
		fmt.Print(s)
	}
}

func (of *OutputFormatter) writelnOutput(s string) {
	if of.outputWriter != nil {
		of.outputWriter.WriteString(s + "\n")
	} else {
		fmt.Println(s)
	}
}

func (of *OutputFormatter) FormatOutput(summary types.ResultSummary) {
	if of.jsonOutput {
		of.outputJSON(summary)
	} else if of.ResponseOnly {
		of.outputResponseOnly(summary)
	} else {
		of.outputResponse(summary)
	}
}

func (of *OutputFormatter) outputJSON(summary types.ResultSummary) {
	jsonData, _ := json.MarshalIndent(summary, "", "  ")
	of.writelnOutput(string(jsonData))
}

func (of *OutputFormatter) outputResponseOnly(summary types.ResultSummary) {
	// 收集所有记录类型的值并去重
	allValues := make(map[string]bool)

	// 从UniqueAnswers中获取所有值
	for _, values := range summary.UniqueAnswers {
		for _, value := range values {
			allValues[value] = true
		}
	}

	// 如果没有有效结果，直接返回，不显示任何内容
	if len(allValues) == 0 {
		return
	}

	// 转换为切片
	uniqueValues := make([]string, 0, len(allValues))
	for value := range allValues {
		uniqueValues = append(uniqueValues, value)
	}

	// 排序（从小到大）
	sort.Strings(uniqueValues)

	// 输出去重后的值
	for _, value := range uniqueValues {
		of.writelnOutput(value)
	}
}

func (of *OutputFormatter) outputResponse(summary types.ResultSummary) {
	// 过滤掉有错误的结果，只保留成功的结果
	validResults := make([]types.RegionResult, 0)
	for _, result := range summary.Results {
		if result.Error == "" {
			validResults = append(validResults, result)
		}
	}

	// 如果没有有效结果，直接返回，不显示任何内容
	if len(validResults) == 0 {
		return
	}

	// 重新构建有效的UniqueAnswers
	validUniqueAnswers := make(map[string]map[string]bool)
	for _, result := range validResults {
		for _, answer := range result.Answers {
			if validUniqueAnswers[answer.Type] == nil {
				validUniqueAnswers[answer.Type] = make(map[string]bool)
			}
			validUniqueAnswers[answer.Type][answer.Value] = true
		}
	}

	// 转换为字符串切片
	uniqueAnswersSlice := make(map[string][]string)
	for recordType, values := range validUniqueAnswers {
		uniqueAnswersSlice[recordType] = types.Keys(values)
	}

	// 使用过滤后的结果进行输出
	recordTypes := make([]string, 0, len(uniqueAnswersSlice))
	for recordType := range uniqueAnswersSlice {
		if of.shouldIncludeRecordType(recordType) {
			recordTypes = append(recordTypes, recordType)
		}
	}
	sort.Strings(recordTypes)
	for _, recordType := range recordTypes {
		values := uniqueAnswersSlice[recordType]
		for _, value := range values {
			if of.colorful {
				color := getColorForRecordType(recordType)
				line := fmt.Sprintf("%s [%s%s%s] [%s%s%s]", summary.Domain, color, recordType, config.ColorReset, config.ColorGreen, value, config.ColorReset)
				of.writelnOutput(line)
			} else {
				line := fmt.Sprintf("%s [%s] [%s]", summary.Domain, recordType, value)
				of.writelnOutput(line)
			}
		}
	}
}

func (of *OutputFormatter) shouldIncludeRecordType(recordType string) bool {
	if len(of.RecordTypes) == 0 {
		return true
	}
	for _, rt := range of.RecordTypes {
		if of.GetRecordTypeName(rt) == recordType {
			return true
		}
	}
	return false
}

func (of *OutputFormatter) GetRecordTypeName(recordType uint16) string {
	switch recordType {
	case 1:
		return "A"
	case 28:
		return "AAAA"
	case 5:
		return "CNAME"
	case 2:
		return "NS"
	case 16:
		return "TXT"
	case 33:
		return "SRV"
	case 12:
		return "PTR"
	case 15:
		return "MX"
	case 6:
		return "SOA"
	case 257:
		return "CAA"
	default:
		return "UNKNOWN"
	}
}

func getColorForRecordType(recordType string) string {
	switch recordType {
	case "A":
		return config.ColorPurple
	case "AAAA":
		return config.ColorBrightBlue
	case "CNAME":
		return config.ColorBrightCyan
	case "NS":
		return config.ColorBrightYellow
	case "TXT":
		return config.ColorBrightGreen
	case "SRV":
		return config.ColorBrightRed
	case "PTR":
		return config.ColorLightBlue
	case "MX":
		return config.ColorOrange
	case "SOA":
		return config.ColorBrightWhite
	case "CAA":
		return config.ColorLightYellow
	default:
		return config.ColorCyan
	}
}

func (of *OutputFormatter) Close() {
	if of.outputWriter != nil {
		of.outputWriter.Close()
	}
}
