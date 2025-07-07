package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/JaveleyQAQ/geodns/internal/config"
	"github.com/JaveleyQAQ/geodns/internal/query"
	"github.com/JaveleyQAQ/geodns/internal/types"
	"github.com/miekg/dns"
)

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// HTTPClient HTTP客户端配置
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient 创建新的HTTP客户端
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 15 * time.Second,
			},
			Timeout: 20 * time.Second,
		},
	}
}

// QueryRegion 查询指定区域的DNS记录
func (h *HTTPClient) QueryRegion(ctx context.Context, domain, region, encodedQuery string, recordType uint16, wg *sync.WaitGroup, resultChan chan<- types.RegionResult) {
	defer wg.Done()

	var url string
	if config.Mode == config.ModeVercel {
		resolverCode := config.GetResolverCode()
		url = fmt.Sprintf(config.Provider.BaseURL, region, encodedQuery, resolverCode, region, query.RandFloat())
	} else {
		url = fmt.Sprintf(config.Provider.BaseURL, encodedQuery, region, query.RandFloat())
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		if config.IsVerbose() {
			log.Printf("[%s] Error creating request: %v\n", region, err)
		}
		return
	}

	resp, err := h.client.Do(req)
	if err != nil {
		if config.IsVerbose() {
			log.Printf("[%s] Error fetching response: %v\n", region, err)
		}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// 添加调试信息（仅在详细模式下显示）
	if config.IsVerbose() {
		log.Printf("[%s] Raw response length: %d bytes", region, len(body))
		if len(body) < 100 {
			log.Printf("[%s] Raw response (hex): %x", region, body[:min(len(body), 100)])
		}
	}

	var answers []types.DNSAnswer
	var aRecords, cnameRecords []string

	msg := new(dns.Msg)
	err = msg.Unpack(body)
	if err != nil {
		if config.IsVerbose() {
			log.Printf("[%s] Failed to unpack DNS response: %v", region, err)
			log.Printf("[%s] Response body (first 100 bytes): %x", region, body[:min(len(body), 100)])
		}
		return
	}

	if config.IsVerbose() {
		log.Printf("[%s] DNS response unpacked successfully, %d answers", region, len(msg.Answer))
	}

	// 检查DNS响应状态
	if msg.Rcode != dns.RcodeSuccess {
		statusText := getDNSStatusText(msg.Rcode)
		if config.IsVerbose() {
			log.Printf("[%s] DNS query failed with status: %s (code: %d)", region, statusText, msg.Rcode)
		}

		// 对于NXDOMAIN等错误状态，仍然返回结果但标记为错误
		resultChan <- types.RegionResult{
			Domain:   domain,
			Region:   region,
			Answers:  []types.DNSAnswer{},
			IPs:      []string{},
			CNAMEs:   []string{},
			RawBytes: body,
			Error:    fmt.Sprintf("DNS %s", statusText),
		}
		return
	}

	for _, ans := range msg.Answer {
		if config.IsVerbose() {
			log.Printf("[%s] Processing record type: %d", region, ans.Header().Rrtype)
		}
		switch v := ans.(type) {
		case *dns.A:
			answers = append(answers, types.DNSAnswer{Type: "A", Value: v.A.String()})
			aRecords = append(aRecords, v.A.String())
		case *dns.AAAA:
			answers = append(answers, types.DNSAnswer{Type: "AAAA", Value: v.AAAA.String()})
		case *dns.CNAME:
			answers = append(answers, types.DNSAnswer{Type: "CNAME", Value: v.Target})
			cnameRecords = append(cnameRecords, v.Target)
		case *dns.NS:
			answers = append(answers, types.DNSAnswer{Type: "NS", Value: v.Ns})
		case *dns.TXT:
			answers = append(answers, types.DNSAnswer{Type: "TXT", Value: strings.Join(v.Txt, " ")})
		case *dns.SRV:
			answers = append(answers, types.DNSAnswer{Type: "SRV", Value: v.Target})
		case *dns.PTR:
			answers = append(answers, types.DNSAnswer{Type: "PTR", Value: v.Ptr})
		case *dns.MX:
			answers = append(answers, types.DNSAnswer{Type: "MX", Value: v.Mx})
		case *dns.SOA:
			answers = append(answers, types.DNSAnswer{Type: "SOA", Value: v.Ns})
		case *dns.CAA:
			answers = append(answers, types.DNSAnswer{Type: "CAA", Value: v.Value})
		default:
			if config.IsVerbose() {
				log.Printf("[%s] Unknown record type: %d", region, ans.Header().Rrtype)
			}
		}
	}

	resultChan <- types.RegionResult{
		Domain:   domain,
		Region:   region,
		Answers:  answers,
		IPs:      aRecords,
		CNAMEs:   cnameRecords,
		RawBytes: body,
	}
}

// getDNSStatusText 获取DNS状态码的文本描述
func getDNSStatusText(rcode int) string {
	switch rcode {
	case dns.RcodeSuccess:
		return "Success"
	case dns.RcodeFormatError:
		return "Format Error"
	case dns.RcodeServerFailure:
		return "Server Failure"
	case dns.RcodeNameError:
		return "Name Error (NXDOMAIN)"
	case dns.RcodeNotImplemented:
		return "Not Implemented"
	case dns.RcodeRefused:
		return "Refused"
	case dns.RcodeYXDomain:
		return "YX Domain"
	case dns.RcodeYXRrset:
		return "YX RRset"
	case dns.RcodeNXRrset:
		return "NX RRset"
	case dns.RcodeNotAuth:
		return "Not Auth"
	case dns.RcodeNotZone:
		return "Not Zone"
	default:
		return fmt.Sprintf("Unknown Error (%d)", rcode)
	}
}
