# GeoDNS 项目结构

```
geodns/
├── cmd/geodns/              # 主程序入口
│   └── main.go             # 命令行参数解析和主流程
├── internal/               # 内部包（不对外暴露）
│   ├── config/            # 配置相关
│   │   ├── constants.go   # 常量定义（颜色、模式等）
│   │   └── provider.go    # DNS提供商配置
│   ├── types/             # 类型定义
│   │   └── types.go       # 数据结构定义
│   ├── query/             # DNS查询相关
│   │   └── query.go       # DNS查询构建
│   ├── client/            # HTTP客户端
│   │   └── http_client.go # HTTP请求处理
│   ├── processor/         # 结果处理
│   │   └── processor.go   # DNS结果处理
│   ├── formatter/         # 输出格式化
│   │   └── formatter.go   # 输出格式控制
│   ├── input/             # 输入处理
│   │   └── input.go       # 输入文件/参数处理
│   └── service/           # 服务层
│       └── service.go     # 业务逻辑服务
├── pkg/                   # 公共包（可对外暴露）
│   └── logo/              # Logo相关
│       └── logo.go        # ASCII艺术Logo
├── testdata/              # 测试数据
│   ├── test_domains.txt
│   ├── test_single.txt
│   └── test_both.txt
├── README.md              # 项目说明
├── PROJECT_STRUCTURE.md   # 项目结构说明
├── go.mod                 # Go模块定义
└── go.sum                 # 依赖校验
```

## 模块说明

### cmd/geodns/
- **main.go**: 程序入口点，负责命令行参数解析和调用各模块服务

### internal/
- **config/**: 配置管理
  - `constants.go`: 颜色常量、DNS模式常量
  - `provider.go`: DNS提供商配置（Vercel/Cloudflare）
- **types/**: 数据结构定义
  - `types.go`: DNSAnswer、RegionResult、ResultSummary等类型
- **query/**: DNS查询构建
  - `query.go`: DNS查询包的构建和编码
- **client/**: HTTP客户端
  - `http_client.go`: HTTP请求处理和DNS响应获取
- **processor/**: 结果处理
  - `processor.go`: DNS结果聚合和去重
- **formatter/**: 输出格式化
  - `formatter.go`: JSON、彩色输出等格式控制
- **input/**: 输入处理
  - `input.go`: 文件读取、标准输入、逗号分隔等输入处理
- **service/**: 服务层
  - `service.go`: 业务逻辑编排，协调各模块工作

### pkg/
- **logo/**: Logo显示
  - `logo.go`: ASCII艺术Logo和版本信息显示

### testdata/
- 测试用的域名文件和示例数据

## 编译和运行

```bash
# 编译
go build -o geodns ./cmd/geodns

# 运行
./geodns -d testdata/test_domains.txt -a -mx
```

## 优势

1. **清晰的职责分离**: 每个模块负责特定功能
2. **易于维护**: 修改某个功能只需关注对应模块
3. **可扩展性**: 新增功能可以独立模块开发
4. **可测试性**: 每个模块可以独立测试
5. **符合Go项目标准**: 遵循Go项目的最佳实践 