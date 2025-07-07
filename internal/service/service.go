package service

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/JaveleyQAQ/geodns/internal/client"
	"github.com/JaveleyQAQ/geodns/internal/config"
	"github.com/JaveleyQAQ/geodns/internal/formatter"
	"github.com/JaveleyQAQ/geodns/internal/input"
	"github.com/JaveleyQAQ/geodns/internal/processor"
	"github.com/JaveleyQAQ/geodns/internal/query"
	"github.com/JaveleyQAQ/geodns/internal/types"
)

type DNSQueryService struct {
	query      *query.DNSQuery
	client     *client.HTTPClient
	processor  *processor.DNSProcessor
	formatter  *formatter.OutputFormatter
	input      *input.InputProcessor
	threads    int
	outputFile string
}

func NewDNSQueryService(jsonOutput, responseOnly, showResponse bool, threads int, recordTypes []uint16, outputFile string) *DNSQueryService {
	return &DNSQueryService{
		query:      query.NewDNSQuery(),
		client:     client.NewHTTPClient(),
		processor:  processor.NewDNSProcessor(),
		formatter:  formatter.NewOutputFormatter(jsonOutput, responseOnly, showResponse, recordTypes, outputFile),
		input:      input.NewInputProcessor(),
		threads:    threads,
		outputFile: outputFile,
	}
}

func (s *DNSQueryService) Query(domain string, recordType uint16) {
	s.processor.Reset()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	encodedQuery := s.query.BuildQuery(domain, recordType)

	var wg sync.WaitGroup
	resultChan := make(chan types.RegionResult, len(config.Provider.Regions))

	for _, region := range config.Provider.Regions {
		wg.Add(1)
		go s.client.QueryRegion(ctx, domain, region, encodedQuery, recordType, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	s.processor.ProcessResults(resultChan)

	summary := s.processor.GetSummary(domain)
	s.formatter.FormatOutput(summary)
}

func (s *DNSQueryService) QueryMultiple(domains []string, recordTypes []uint16) {
	s.processor.Reset()

	semaphore := make(chan struct{}, s.threads)
	var wg sync.WaitGroup
	resultChan := make(chan types.RegionResult, len(domains)*len(recordTypes)*len(config.Provider.Regions))

	for _, domain := range domains {
		for _, recordType := range recordTypes {
			wg.Add(1)
			go func(d string, rt uint16) {
				defer wg.Done()
				semaphore <- struct{}{}
				defer func() { <-semaphore }()
				s.queryDomain(d, rt, resultChan)
			}(domain, recordType)
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	s.processor.ProcessResults(resultChan)

	// 检查是否为-ro模式（只输出响应值）
	if s.formatter.ResponseOnly {
		allValues := make(map[string]bool)
		for _, domain := range domains {
			summary := s.processor.GetSummary(domain)
			for _, rt := range s.formatter.RecordTypes {
				typeName := s.formatter.GetRecordTypeName(rt)
				if values, ok := summary.UniqueAnswers[typeName]; ok {
					for _, value := range values {
						allValues[value] = true
					}
				}
			}
		}
		uniqueValues := make([]string, 0, len(allValues))
		for value := range allValues {
			uniqueValues = append(uniqueValues, value)
		}
		sort.Strings(uniqueValues)
		for _, value := range uniqueValues {
			fmt.Println(value)
		}
		return
	}

	// 其他模式，逐域名输出
	for _, domain := range domains {
		summary := s.processor.GetSummary(domain)
		s.formatter.FormatOutput(summary)
	}
}

func (s *DNSQueryService) queryDomain(domain string, recordType uint16, resultChan chan<- types.RegionResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	encodedQuery := s.query.BuildQuery(domain, recordType)

	var wg sync.WaitGroup
	domainResultChan := make(chan types.RegionResult, len(config.Provider.Regions))

	for _, region := range config.Provider.Regions {
		wg.Add(1)
		go s.client.QueryRegion(ctx, domain, region, encodedQuery, recordType, &wg, domainResultChan)
	}

	go func() {
		wg.Wait()
		close(domainResultChan)
	}()

	for res := range domainResultChan {
		resultChan <- res
	}
}

func (s *DNSQueryService) Close() {
    if s.formatter != nil {
        s.formatter.Close()
    }
}

