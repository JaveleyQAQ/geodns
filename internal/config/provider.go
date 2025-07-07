package config

import (
	"fmt"
	"os"
)

// DNSProvider DNS服务提供商
type DNSProvider struct {
	BaseURL string
	Regions []string
}

// Resolver DNS解析器类型
type Resolver struct {
	Name string
	Code string
}

var Provider *DNSProvider
var Mode string
var CurrentResolver *Resolver
var Verbose bool // 添加调试模式变量

// 支持的解析器
var SupportedResolvers = map[string]*Resolver{
	"alidns":     {Name: "AliDNS", Code: "alidns"},
	"google":     {Name: "Google DNS", Code: "google"},
	"cloudflare": {Name: "Cloudflare", Code: "cloudflare"},
}

func init() {
	Mode = os.Getenv("DNS_MODE")
	if Mode == "" {
		Mode = ModeVercel
	}

	// 默认使用cloudflare解析器
	CurrentResolver = SupportedResolvers["cloudflare"]

	if Mode == ModeVercel {
		Provider = &DNSProvider{
			BaseURL: "https://vercel.dns.surf/api/region/%s?dns=%s&resolver=%s&region=%s&_=%.16f",
			Regions: []string{
				"hnd1", "kix1", "sin1", "icn1", "bom1", "syd1", "cpt1",
				"arn1", "dub1", "lhr1", "fra1", "cdg1", "hkg1",
			},
		}
	} else {
		Provider = &DNSProvider{
			BaseURL: "https://cloudflare.dns.surf/?dns=%s&region=%s&_=%.16f",
			Regions: []string{
				"ams", "arn", "bom", "cdg", "cle", "den", "dfw", "ewr", "fra", "gru", "hkg", "iad",
				"jfk", "lax", "lhr", "mad", "man", "nrt", "ord", "otp", "par", "sea", "sgp", "sin",
				"sfo", "syd", "tpe", "yul", "yyz", "zag", "zur", "mex", "maa", "del", "dac", "ccu",
				"khi", "isb", "tun", "jnb", "cpt", "los", "abuja", "kgl", "mpm", "dar", "nbo", "acc",
			},
		}
	}
}

// SetResolver 设置解析器
func SetResolver(resolverName string) error {
	if resolver, exists := SupportedResolvers[resolverName]; exists {
		CurrentResolver = resolver
		return nil
	}
	return fmt.Errorf("unsupported resolver: %s. Supported: alidns, google, cloudflare", resolverName)
}

// GetResolverCode 获取当前解析器代码
func GetResolverCode() string {
	return CurrentResolver.Code
}

// SetVerbose 设置调试模式
func SetVerbose(verbose bool) {
	Verbose = verbose
}

// IsVerbose 检查是否启用调试模式
func IsVerbose() bool {
	return Verbose
}
