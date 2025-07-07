# GEODNS - 全球DNS查询工具

一个高性能的**全球DNS查询工具**，支持多区域与批量DNS查询、多种记录类型和丰富的输出格式，提供比传统DNS工具更简洁、更高效的查询体验。

## 🎯 核心优势

### 🌍 **全球DNS查询**
- **13个全球区域**：支持亚太、欧洲、美洲、非洲的DNS查询
- **智能区域选择**：自动选择最佳查询区域，提高查询成功率
- **地理分布优化**：根据不同地理位置返回最优DNS结果

### 🎨 **智能输出设计**
- **自动去重**：智能去除重复记录，输出更清晰
- **彩色分类**：不同记录类型用不同颜色区分，一目了然
- **多种输出模式**：标准输出、JSON格式、仅响应值、文件输出
- **静默模式**：适合脚本集成和自动化工作流

### 🚀 **高性能特性**
- **高并发处理**：可自定义并发线程数，支持批量域名查询
- **智能错误处理**：优雅处理网络错误和DNS解析失败
- **内存优化**：高效的结果处理和内存管理
- **连接复用**：复用HTTP连接，减少连接开销

### 🔧 **多解析器支持**
- **AliDNS**：阿里DNS (223.5.5.5)
- **Google DNS**：Google DNS (8.8.8.8)
- **Cloudflare**：Cloudflare DNS (1.1.1.1)

## 📦 安装
## 方式一 : GO install 
```
go install github.com/JaveleyQAQ/geodns/cmd/geodns@latest
```


### 方式二: 从源码编译
```bash
git clone https://github.com/JaveleyQAQ/geodns.git
cd geodns
go mod tidy
go build -o geodns cmd/geodns/main.go
```

### 直接使用
```bash
./geodns -h
```

## 🚀 使用方法

### 命令行参数

#### 输入选项
- `-l string` - 子域名列表文件或标准输入
- `-d string` - 域名列表，文件/逗号分隔/标准输入

#### 查询类型
- `-a` - 查询A记录
- `-aaaa` - 查询AAAA记录
- `-cname` - 查询CNAME记录
- `-ns` - 查询NS记录
- `-txt` - 查询TXT记录
- `-srv` - 查询SRV记录
- `-ptr` - 查询PTR记录
- `-mx` - 查询MX记录
- `-soa` - 查询SOA记录
- `-any` - 查询ANY记录
- `-axfr` - 查询AXFR记录
- `-caa` - 查询CAA记录
- `-recon` - 查询所有类型

#### 输出控制
- `-re` - 显示响应
- `-ro` - 只输出响应值
- `-json` - 输出完整JSON格式
- `-o string` - 输出到指定文件
- `-silent` - 静默模式，不显示logo

#### 其他选项
- `-r string` - DNS解析器 (alidns/google/cloudflare) (默认: cloudflare)
- `-t int` - 并发线程数 (默认: 10)
- `-v` - 详细模式，显示调试信息

## 📝 使用示例

### 基本查询
```bash
# 查询单个域名的A记录
./geodns -d google.com

# 查询多个域名
./geodns -d "google.com,github.com,example.com"

# 从文件查询域名列表
./geodns -d domains.txt

# 从标准输入查询
echo "google.com" | ./geodns -d -
```

### 指定记录类型
```bash
# 查询AAAA记录
./geodns -d google.com -aaaa

# 查询TXT记录
./geodns -d google.com -txt

# 查询MX记录
./geodns -d google.com -mx

# 查询所有记录类型（侦察模式）
./geodns -d google.com -recon
```

### 输出格式控制
```bash
# 只显示响应值
./geodns -d google.com -ro

# JSON格式输出
./geodns -d google.com -json

# 输出到文件
./geodns -d google.com -o results.txt

# 静默模式（适合脚本）
./geodns -d google.com -silent -ro

# 组合使用：静默模式 + 文件输出 + JSON格式
./geodns -d domains.txt -recon -json -silent -o results.json
```

### 高级功能
```bash
# 使用Google DNS解析器
./geodns -d google.com -r google

# 设置50个并发线程
./geodns -d domains.txt -t 50

# 详细模式（显示调试信息）
./geodns -d google.com -mx -v

# 组合使用
./geodns -d domains.txt -recon -json -r alidns -t 20
```

## 🎨 输出格式

### 标准输出格式
```
域名 [记录类型] [值]
```


### JSON输出格式
```json
{
  "domain": "google.com",
  "results": [
    {
      "domain": "google.com",
      "region": "hnd1",
      "answers": [
        {
          "type": "A",
          "value": "142.250.197.110"
        }
      ]
    }
  ],
  "unique_answers": {
    "A": ["142.250.197.110", "142.250.197.174"]
  }
}
```

## 🌍 支持的全球区域

### Vercel模式（默认）
- **亚太地区**: hnd1(东京), kix1(大阪), sin1(新加坡), icn1(首尔), bom1(孟买), syd1(悉尼), hkg1(香港)
- **欧洲地区**: lhr1(伦敦), fra1(法兰克福), cdg1(巴黎), dub1(都柏林), arn1(斯德哥尔摩)
- **非洲地区**: cpt1(开普敦)

### Cloudflare模式
- **美洲**: ams, den, dfw, ewr, iad, jfk, lax, ord, sea, sfo, yul, yyz, mex
- **欧洲**: fra, gru, lhr, mad, man, otp, par, zag, zur
- **亚太**: hkg, nrt, sgp, sin, tpe

## 🔧 配置选项

### DNS解析器
- `alidns` - 阿里DNS (223.5.5.5)
- `google` - Google DNS (8.8.8.8)  
- `cloudflare` - Cloudflare DNS (1.1.1.1)

## 📁 输入文件格式

### 域名列表文件 (domains.txt)
```
google.com
github.com
example.com
microsoft.com
```


## 🐛 调试模式

使用 `-v` 参数启用详细模式，显示：
- 原始DNS响应长度
- 原始响应的十六进制数据
- DNS响应解析状态
- 记录类型处理信息

```bash
./geodns -d google.com -mx -v
```

## ⚠️ 注意事项

1. **输入参数限制**：不能同时使用 `-l` 和 `-d` 参数
2. **默认行为**：不指定查询类型时默认查询A记录
3. **并发控制**：建议根据网络环境调整线程数（默认10）
4. **DNS解析器**：不同解析器可能返回略有不同的结果
5. **网络环境**：某些区域可能因网络限制无法访问

## 🚀 性能优化

- **连接池复用**：复用HTTP连接，减少连接开销
- **并发控制**：可调节并发线程数，平衡性能和稳定性
- **智能超时**：设置合理的超时时间，避免长时间等待
- **内存优化**：高效的结果处理和内存管理
- **去重算法**：自动去除重复记录，减少输出冗余

## 🔍 故障排除

### 常见问题
1. **无输出结果**：检查域名格式和网络连接
2. **部分区域无响应**：可能是网络限制，尝试其他区域
3. **解析器问题**：尝试切换不同的DNS解析器
4. **并发过高**：降低线程数避免被限制

### 调试技巧
```bash
# 启用详细模式查看问题
./geodns -d google.com -v

# 使用不同解析器
./geodns -d google.com -r google

# 降低并发数
./geodns -d domains.txt -t 5

# 静默模式调试
./geodns -d google.com -silent -v
```

## 🤝 贡献
- 所有代码都来自我的助手`cursor`
- 欢迎提交Issue和Pull Request来改进这个项目！

## 📞 联系方式

- GitHub: https://github.com/JaveleyQAQ
- 项目地址: https://github.com/JaveleyQAQ/geodns
- 数据来源: https://dns.surf/
