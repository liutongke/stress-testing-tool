package http_client

import (
	"flag"
	"fmt"
	"strings"
)

// HeaderList 用于存储多个HTTP头信息
type HeaderDict map[string]string

func (h *HeaderDict) String() string {
	return fmt.Sprint(*h)
}

func (h *HeaderDict) Set(value string) error {
	if *h == nil {
		*h = make(HeaderDict) // 初始化 HeaderDict，如果它是 nil
	}

	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid header format: %s", value)
	}

	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	(*h)[key] = val

	return nil
}

// ABConfig 是压测工具的配置参数结构体
type ABConfig struct {
	Requests      int
	Method        string
	Concurrency   int
	Timelimit     int
	Postfile      string
	ContentType   string
	Headers       HeaderDict
	Putfile       string
	Verbosity     int
	UseHTML       bool
	UseHEAD       bool
	Cookie        string
	Auth          string
	ProxyAuth     string
	Proxy         string
	PrintVersion  bool
	HTTPKeepAlive bool
	NoPercentiles bool
	NoErrorStats  bool
	QuietMode     bool
	AcceptVarLen  bool
	GnuplotFile   string
	CSVFile       string
	RequestURL    string
}

// go run .\main.go -n 100 -c 10 -T application/json -H "Accept: application/json" -H "Cache-Control: no-cache" -p /path/to/postfile -k -url http://example.com/
// NewFlagConfig 用于解析命令行参数并返回一个配置实例
func NewFlagConfig() *ABConfig {
	cfg := &ABConfig{}

	flag.StringVar(&cfg.RequestURL, "url", "http://example.com", "压测的目标地址。")
	flag.StringVar(&cfg.Method, "m", "GET", "HTTP请求方法")
	flag.IntVar(&cfg.Requests, "n", 1, "总请求次数")
	flag.IntVar(&cfg.Concurrency, "c", 1, "并发请求数")
	flag.IntVar(&cfg.Timelimit, "t", 0, "最长执行时间（秒）")
	flag.StringVar(&cfg.Postfile, "p", "", "包含要POST的数据的文件路径")
	flag.StringVar(&cfg.ContentType, "T", "application/x-www-form-urlencoded", "POST数据的Content-Type")
	flag.Var(&cfg.Headers, "H", "自定义HTTP头信息。例如：-H 'Accept: application/json' -H 'Cookie: mycookie'")
	flag.StringVar(&cfg.Putfile, "u", "", "包含要PUT的数据的文件路径")
	flag.IntVar(&cfg.Verbosity, "v", 0, "设置显示的信息的详细程度")
	flag.BoolVar(&cfg.UseHTML, "w", false, "输出HTML表格")
	flag.BoolVar(&cfg.UseHEAD, "i", false, "执行HEAD请求")
	flag.StringVar(&cfg.Cookie, "C", "", "添加Cookie信息，格式为：name=value")
	flag.StringVar(&cfg.Auth, "A", "", "添加基础认证，格式为：username:password")
	flag.StringVar(&cfg.ProxyAuth, "P", "", "添加代理服务器认证，格式为：username:password")
	flag.StringVar(&cfg.Proxy, "X", "", "设置代理服务器，格式为：proxy:port")
	flag.BoolVar(&cfg.PrintVersion, "V", false, "打印版本信息")
	flag.BoolVar(&cfg.HTTPKeepAlive, "k", false, "启用HTTP KeepAlive")
	flag.BoolVar(&cfg.NoPercentiles, "d", false, "不显示百分比满意度表")
	flag.BoolVar(&cfg.NoErrorStats, "S", false, "不显示错误统计信息")
	flag.BoolVar(&cfg.QuietMode, "q", false, "静默模式，不显示进度")
	flag.BoolVar(&cfg.AcceptVarLen, "l", false, "接受可变长度的返回包")
	flag.StringVar(&cfg.GnuplotFile, "g", "", "输出测试结果到'gnuplot'或'TSV'文件")
	flag.StringVar(&cfg.CSVFile, "e", "", "输出测试结果到CSV文件")

	flag.Parse()

	return cfg
}
