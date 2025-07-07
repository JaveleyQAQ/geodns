package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JaveleyQAQ/geodns/internal/config"
	"github.com/JaveleyQAQ/geodns/internal/input"
	"github.com/JaveleyQAQ/geodns/internal/service"
	"github.com/JaveleyQAQ/geodns/pkg/logo"
)

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Println("  Input options:")
	fmt.Println("    -l string\t子域名列表文件或标准输入")
	fmt.Println("    -d string\t域名列表，文件/逗号分隔/标准输入")
	fmt.Println()
	fmt.Println("  Query options:")
	fmt.Println("    -a\t查询A记录")
	fmt.Println("    -aaaa\t查询AAAA记录")
	fmt.Println("    -cname\t查询CNAME记录")
	fmt.Println("    -ns\t查询NS记录")
	fmt.Println("    -txt\t查询TXT记录")
	fmt.Println("    -srv\t查询SRV记录")
	fmt.Println("    -ptr\t查询PTR记录")
	fmt.Println("    -mx\t查询MX记录")
	fmt.Println("    -soa\t查询SOA记录")
	fmt.Println("    -any\t查询ANY记录")
	fmt.Println("    -axfr\t查询AXFR记录")
	fmt.Println("    -caa\t查询CAA记录")
	fmt.Println("    -recon\t查询所有类型")
	fmt.Println()
	fmt.Println("  Filter options:")
	fmt.Println("    -re\t显示响应")
	fmt.Println("    -ro\t只输出响应值")
	fmt.Println("    -json\t输出完整JSON格式")
	fmt.Println()
	fmt.Println("  Output options:")
	fmt.Println("    -o string\t输出到指定文件")
	fmt.Println("    -silent\t静默模式，不显示logo")
	fmt.Println()
	fmt.Println("  Other options:")
	fmt.Println("    -r string\tDNS解析器 (alidns/google/cloudflare) (default cloudflare)")
	fmt.Println("    -t int\t并发线程数 (default 10)")
	fmt.Println("    -v\t\t详细模式，显示调试信息")
}

func main() {
	flag.Usage = printUsage

	// Input
	inputFile := flag.String("l", "", "子域名列表文件或标准输入")
	domainArg := flag.String("d", "", "域名列表，文件/逗号分隔/标准输入")

	// Query
	a := flag.Bool("a", false, "查询A记录")
	aaaa := flag.Bool("aaaa", false, "查询AAAA记录")
	cname := flag.Bool("cname", false, "查询CNAME记录")
	ns := flag.Bool("ns", false, "查询NS记录")
	txt := flag.Bool("txt", false, "查询TXT记录")
	srv := flag.Bool("srv", false, "查询SRV记录")
	ptr := flag.Bool("ptr", false, "查询PTR记录")
	mx := flag.Bool("mx", false, "查询MX记录")
	soa := flag.Bool("soa", false, "查询SOA记录")
	any := flag.Bool("any", false, "查询ANY记录")
	axfr := flag.Bool("axfr", false, "查询AXFR记录")
	caa := flag.Bool("caa", false, "查询CAA记录")
	recon := flag.Bool("recon", false, "查询所有类型")

	// Filter
	showResponse := flag.Bool("re", false, "显示响应")
	responseOnly := flag.Bool("ro", false, "只输出响应值")
	jsonOutput := flag.Bool("json", false, "输出完整JSON格式")

	// Output
	outputFile := flag.String("o", "", "输出到指定文件")
	silent := flag.Bool("silent", false, "静默模式，不显示logo")

	// Other
	resolver := flag.String("r", "cloudflare", "DNS解析器 (alidns/google/cloudflare)")
	threads := flag.Int("t", 10, "并发线程数")
	verbose := flag.Bool("v", false, "详细模式，显示调试信息")

	flag.Parse()

	// 显示 logo（除非静默模式）
	if !*silent {
		logo.PrintLogo()
	}

	// 设置调试模式
	config.SetVerbose(*verbose)

	// 设置解析器
	if err := config.SetResolver(*resolver); err != nil {
		fmt.Fprintf(os.Stderr, "解析器设置错误: %v\n", err)
		os.Exit(1)
	}

	// 输入参数校验
	if *inputFile != "" && *domainArg != "" {
		fmt.Fprintln(os.Stderr, "-l 和 -d 参数不能同时使用")
		os.Exit(1)
	}

	var domains []string
	var err error
	inputProcessor := input.NewInputProcessor()

	// 检查是否有标准输入
	stat, _ := os.Stdin.Stat()
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0

	if *inputFile != "" {
		domains, err = inputProcessor.ReadFromFile(*inputFile)
	} else if *domainArg != "" {
		domains, err = inputProcessor.GetDomains(*domainArg)
	} else if hasStdin {
		// 自动检测标准输入
		domains, err = inputProcessor.ReadFromStdin()
	} else {
		fmt.Fprintln(os.Stderr, "请使用 -l 或 -d 指定域名输入，或通过管道提供输入")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "输入处理错误: %v\n", err)
		os.Exit(1)
	}

	// 记录类型收集
	recordTypes := []uint16{}
	if *recon {
		recordTypes = []uint16{1, 28, 5, 2, 16, 33, 12, 15, 6, 257} // A,AAAA,CNAME,NS,TXT,SRV,PTR,MX,SOA,CAA
	} else {
		if *a {
			recordTypes = append(recordTypes, 1)
		}
		if *aaaa {
			recordTypes = append(recordTypes, 28)
		}
		if *cname {
			recordTypes = append(recordTypes, 5)
		}
		if *ns {
			recordTypes = append(recordTypes, 2)
		}
		if *txt {
			recordTypes = append(recordTypes, 16)
		}
		if *srv {
			recordTypes = append(recordTypes, 33)
		}
		if *ptr {
			recordTypes = append(recordTypes, 12)
		}
		if *mx {
			recordTypes = append(recordTypes, 15)
		}
		if *soa {
			recordTypes = append(recordTypes, 6)
		}
		if *any {
			recordTypes = append(recordTypes, 255)
		}
		if *axfr {
			recordTypes = append(recordTypes, 252)
		}
		if *caa {
			recordTypes = append(recordTypes, 257)
		}
	}
	if len(recordTypes) == 0 {
		recordTypes = append(recordTypes, 1) // 默认A记录
	}

	dnsService := service.NewDNSQueryService(*jsonOutput, *responseOnly, *showResponse, *threads, recordTypes, *outputFile)
	dnsService.QueryMultiple(domains, recordTypes)
	dnsService.Close()
}
