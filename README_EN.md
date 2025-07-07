 # GEODNS - Global DNS Query Tool

A high-performance **global DNS query tool** that supports multi-region and batch DNS queries, multiple record types, and rich output formats, providing a more concise and efficient query experience than traditional DNS tools.

## 🎯 Core Advantages

### 🌍 **Global DNS Query**
- **13 Global Regions**: Supports DNS queries across Asia-Pacific, Europe, Americas, and Africa
- **Smart Region Selection**: Automatically selects optimal query regions for higher success rates
- **Geographic Distribution Optimization**: Returns optimal DNS results based on different geographic locations

### 🎨 **Intelligent Output Design**
- **Automatic Deduplication**: Intelligently removes duplicate records for cleaner output
- **Color Classification**: Different record types distinguished by different colors for easy identification
- **Multiple Output Modes**: Standard output, JSON format, response-only, file output
- **Silent Mode**: Perfect for script integration and automated workflows

### 🚀 **High-Performance Features**
- **High Concurrency Processing**: Customizable concurrent threads supporting batch domain queries
- **Smart Error Handling**: Gracefully handles network errors and DNS resolution failures
- **Memory Optimization**: Efficient result processing and memory management
- **Connection Reuse**: Reuses HTTP connections to reduce connection overhead

### 🔧 **Multi-Resolver Support**
- **AliDNS**: Alibaba DNS (223.5.5.5)
- **Google DNS**: Google DNS (8.8.8.8)
- **Cloudflare**: Cloudflare DNS (1.1.1.1)

## 📦 Installation

### Method 1: GO install
```bash
go install github.com/JaveleyQAQ/geodns/cmd/geodns@latest
```

### Method 2: Build from Source
```bash
git clone https://github.com/JaveleyQAQ/geodns.git
cd geodns
go mod tidy
go build -o geodns cmd/geodns/main.go
```

### Direct Usage
```bash
geodns -h
```

## 🚀 Usage

### Command Line Arguments

#### Input Options
- `-l string` - Subdomain list file or pipe input
- `-d string` - Domain list, file/comma-separated/standard input

#### Query Types
- `-a` - Query A records
- `-aaaa` - Query AAAA records
- `-cname` - Query CNAME records
- `-ns` - Query NS records
- `-txt` - Query TXT records
- `-srv` - Query SRV records
- `-ptr` - Query PTR records
- `-mx` - Query MX records
- `-soa` - Query SOA records
- `-any` - Query ANY records
- `-axfr` - Query AXFR records
- `-caa` - Query CAA records
- `-recon` - Query all record types

#### Output Control
- `-re` - Show response
- `-ro` - Output response values only
- `-json` - Output complete JSON format
- `-o string` - Output to specified file
- `-silent` - Silent mode, hide logo

#### Other Options
- `-r string` - DNS resolver (alidns/google/cloudflare) (default: cloudflare)
- `-t int` - Concurrent threads (default: 10)
- `-v` - Verbose mode, show debug information

## 📝 Usage Examples

### Basic Queries
- Query A records for a single domain
```bash
geodns -d google.com

google.com [A] [142.250.197.46]
google.com [A] [142.250.198.110]
google.com [A] [142.250.199.78]
google.com [A] [142.250.76.14]
google.com [A] [142.251.10.100]
google.com [A] [142.251.10.101]
google.com [A] [142.251.10.102]
google.com [A] [142.251.10.113]
google.com [A] [142.251.10.138]
google.com [A] [142.251.10.139]
google.com [A] [142.251.111.100]
google.com [A] [142.251.111.101]
google.com [A] [142.251.111.102]
google.com [A] [142.251.111.113]
google.com [A] [142.251.111.138]
google.com [A] [142.251.111.139]
google.com [A] [142.251.12.100]
google.com [A] [142.251.12.101]
google.com [A] [142.251.12.102]
google.com [A] [142.251.12.113]
google.com [A] [142.251.12.138]
google.com [A] [142.251.12.139]
google.com [A] [172.253.118.100]
google.com [A] [172.253.118.101]
google.com [A] [172.253.118.102]
google.com [A] [172.253.118.113]
google.com [A] [172.253.118.138]
google.com [A] [172.253.118.139]
google.com [A] [172.253.122.100]
google.com [A] [172.253.122.101]
google.com [A] [172.253.122.102]
google.com [A] [172.253.122.113]
google.com [A] [172.253.122.138]
google.com [A] [172.253.122.139]
google.com [A] [64.233.170.100]
google.com [A] [64.233.170.101]
google.com [A] [64.233.170.102]
google.com [A] [64.233.170.113]
google.com [A] [64.233.170.138]
google.com [A] [64.233.170.139]
google.com [A] [74.125.24.100]
google.com [A] [74.125.24.101]
google.com [A] [74.125.24.102]
google.com [A] [74.125.24.113]
google.com [A] [74.125.24.138]
google.com [A] [74.125.24.139]
google.com [A] [74.125.68.100]
google.com [A] [74.125.68.101]
google.com [A] [74.125.68.102]
google.com [A] [74.125.68.113]
google.com [A] [74.125.68.138]
google.com [A] [74.125.68.139]
```

- Query multiple domains AAAA records and show results only
```bash
geodns -d "google.com,github.com,example.com" -ro -silent

142.250.197.174
142.250.198.110
142.250.199.78
142.250.71.238
142.250.76.238
142.251.12.100
142.251.12.101
142.251.12.102
142.251.12.113
142.251.12.138
142.251.12.139
172.217.194.100
172.217.194.101
172.217.194.102
172.217.194.113
172.217.194.138
172.217.194.139
20.205.243.166
23.192.228.80
23.192.228.84
23.215.0.136
23.215.0.138
64.233.170.100
64.233.170.101
64.233.170.102
64.233.170.113
64.233.170.138
64.233.170.139
74.125.130.100
74.125.130.101
74.125.130.102
74.125.130.113
74.125.130.138
74.125.130.139
74.125.68.100
74.125.68.101
74.125.68.102
74.125.68.113
74.125.68.138
74.125.68.139
96.7.128.175
96.7.128.198
```

- Query domain list from file
```bash
geodns -d domains.txt
```

- Query from standard input
```bash
echo "google.com" | geodns 
cat url.txt | geodns
```

- Query all record types (reconnaissance mode)
```bash
geodns -d google.com -recon
```

- Output format control
```bash
# Show response values only
geodns -d google.com -ro

# JSON format output
geodns -d google.com -json

# Output to file
geodns -d google.com -o results.txt

# Silent mode (suitable for scripts)
geodns -d google.com -silent -ro

# Combined usage: silent mode + file output + JSON format
./geodns -d domains.txt -recon -json -silent -o results.json
```

### Advanced Features
```bash
# Use Google DNS resolver
geodns -d google.com -r google

# Set 50 concurrent threads
geodns -d domains.txt -t 50

# Verbose mode (show debug information)
geodns -d google.com -mx -v

# Combined usage
geodns -d domains.txt -recon -json -r alidns -t 20
```

## 🎨 Output Format

### Standard Output Format
```
domain [record_type] [value]
```

### JSON Output Format
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

## 🌍 Supported Global Regions

### Vercel Mode (Default)
- **Asia-Pacific**: hnd1(Tokyo), kix1(Osaka), sin1(Singapore), icn1(Seoul), bom1(Mumbai), syd1(Sydney), hkg1(Hong Kong)
- **Europe**: lhr1(London), fra1(Frankfurt), cdg1(Paris), dub1(Dublin), arn1(Stockholm)
- **Africa**: cpt1(Cape Town)

### Cloudflare Mode
- **Americas**: ams, den, dfw, ewr, iad, jfk, lax, ord, sea, sfo, yul, yyz, mex
- **Europe**: fra, gru, lhr, mad, man, otp, par, zag, zur
- **Asia-Pacific**: hkg, nrt, sgp, sin, tpe

## 🔧 Configuration Options

### DNS Resolvers
- `alidns` - Alibaba DNS (223.5.5.5)
- `google` - Google DNS (8.8.8.8)  
- `cloudflare` - Cloudflare DNS (1.1.1.1)

## 📁 Input File Formats

### Domain List File (domains.txt)
```
google.com
github.com
example.com
microsoft.com
```

## 🐛 Debug Mode

Use the `-v` parameter to enable verbose mode, displaying:
- Raw DNS response length
- Raw response hexadecimal data
- DNS response parsing status
- Record type processing information

```bash
geodns -d google.com -mx -v
```

## ⚠️ Important Notes

1. **Input Parameter Restrictions**: Cannot use `-l` and `-d` parameters simultaneously
2. **Default Behavior**: Queries A records by default when no query type is specified
3. **Concurrency Control**: Adjust thread count based on network environment (default: 10)
4. **DNS Resolvers**: Different resolvers may return slightly different results
5. **Network Environment**: Some regions may be inaccessible due to network restrictions

## 🚀 Performance Optimization

- **Connection Pool Reuse**: Reuses HTTP connections to reduce connection overhead
- **Concurrency Control**: Adjustable concurrent threads to balance performance and stability
- **Smart Timeout**: Sets reasonable timeout to avoid long waits
- **Memory Optimization**: Efficient result processing and memory management
- **Deduplication Algorithm**: Automatically removes duplicate records to reduce output redundancy

## 🔍 Troubleshooting

### Common Issues
1. **No Output Results**: Check domain format and network connection
2. **Partial Region No Response**: May be due to network restrictions, try other regions
3. **Resolver Issues**: Try switching to different DNS resolvers
4. **High Concurrency**: Reduce thread count to avoid being rate-limited

### Debugging Tips
```bash
# Enable verbose mode to view issues
geodns -d google.com -v

# Use different resolver
geodns -d google.com -r google

# Reduce concurrency
geodns -d domains.txt -t 5

# Silent mode debugging
geodns -d google.com -silent -v
```

## 🤝 Contributing

Welcome to submit Issues and Pull Requests to improve this project!

## 📞 Contact

- GitHub: https://github.com/JaveleyQAQ
- Project URL: https://github.com/JaveleyQAQ/geodns
- Data Source: https://dns.surf/