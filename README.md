# SnowflakeID

A high-performance distributed ID generator based on Twitter's Snowflake algorithm, designed for Go language.

## Features

- 🚀 **High Performance**: Efficient ID generation based on 64-bit integers
- 🔒 **Distributed Safe**: Supports multi-node deployment, avoiding ID conflicts
- 📅 **Time Ordered**: IDs increase in chronological order, facilitating sorting and indexing
- 🏷️ **Business Classification**: Built-in business line classification prefixes, supporting multiple business scenarios
- 🔄 **Base32 Encoding**: Supports Base32 encoding output, URL-safe for transmission
- ⚡ **Clock Drift Protection**: Built-in clock drift detection and handling mechanisms
- 🧵 **Concurrency Safe**: Uses mutex locks to ensure safety in multi-threaded environments

## Bit Allocation

```
┌───────────────────────────────────────────────────────────────────────────────┐
│                               64-bit ID                                      │
├─────────────────┬─────────────┬─────────────┬─────────────┬─────────────────┤
│     1 bit      │    5 bits   │   2 bits   │   3 bits   │     4 bits      │
│   (Sign Bit)   │  (Prefix)   │ (Version)  │(Business)  │   (System ID)   │
├─────────────────┴─────────────┴─────────────┴─────────────┴─────────────────┤
│                               42 bits                                       │
│                            (Timestamp)                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                               7 bits                                        │
│                             (Sequence)                                      │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Detailed Description

- **Sign Bit (1 bit)**: Fixed at 0, reserved bit
- **Prefix (5 bits)**: Business line + machine group identifier, supports 32 classifications
- **Version (2 bits)**: Version number, supports 4 versions
- **Business (3 bits)**: Sub-business line, supports 8 sub-businesses
- **System ID (4 bits)**: System node identifier, supports 16 nodes
- **Timestamp (42 bits)**: Millisecond-level timestamp, supports 139 years
- **Sequence (7 bits)**: Sequence number, supports 128 IDs per millisecond

## Installation

```bash
go get binrc.com/pkg/snowflakeid
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "binrc.com/pkg/snowflakeid"
)

func main() {
    // Create generator instance
    // Param 1: BusinessID (0-7)
    // Param 2: SystemID (0-15)
    generator := snowflakeid.NewGenerator(1, 2)
    
    // Generate ID
    id, base32ID, err := generator.NextID()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Decimal ID: %d\n", id)
    fmt.Printf("Base32 ID: %s\n", base32ID)
}
```

### Using Custom Prefixes

```go
// Use business-related prefix
id, base32ID, err := generator.NextIDWithPrefix(snowflakeid.BusinessBit)

// Use customer-related prefix
id, base32ID, err := generator.NextIDWithPrefix(snowflakeid.CustomerBit)

// Use device-related prefix
id, base32ID, err := generator.NextIDWithPrefix(snowflakeid.DeviceBit)
```

### Parsing ID Information

```go
// Parse ID into structured information
snowflakeInfo := snowflakeid.ParseID(id)
fmt.Printf("Business ID: %d\n", snowflakeInfo.Business)
fmt.Printf("System ID: %d\n", snowflakeInfo.SystemID)
fmt.Printf("Timestamp: %s\n", snowflakeInfo.CreatedTime)
fmt.Printf("Base32: %s\n", snowflakeInfo.Base32)
```

## Business Prefix Classification

| Prefix | Binary | Decimal | Business Type | Description |
|--------|--------|---------|---------------|-------------|
| A      | 00000  | 0       | Default       | General business |
| B      | 00001  | 1       | Business      | Core business logic |
| C      | 00010  | 2       | Customer      | Customer management related |
| D      | 00011  | 3       | Device        | Device management related |
| E      | 00100  | 4       | Event         | Event recording related |
| F      | 00101  | 5       | File          | File management related |
| G      | 00110  | 6       | Gateway       | Gateway service related |
| H      | 00111  | 7       | Host          | Host management related |
| I      | 01000  | 8       | Instance      | Instance management related |
| J      | 01001  | 9       | Job           | Task scheduling related |
| K      | 01010  | 10      | Kubernetes    | Container orchestration related |
| L      | 01011  | 11      | Log           | Log management related |
| M      | 01100  | 12      | Module        | Module management related |
| N      | 01101  | 13      | Network       | Network management related |

## Configuration Options

### Epoch Time

Default epoch time is `2025-01-01 UTC`, which can be adjusted by modifying the `Epoch` constant.

### Logging Configuration

```go
import "log"

// Set custom logger
generator := snowflakeid.NewGenerator(1, 2)
generator.Logger = log.New(os.Stdout, "[SnowflakeID] ", log.LstdFlags)
```

## Performance Characteristics

- **Generation Speed**: Can generate 128 unique IDs per millisecond
- **Clock Precision**: Millisecond-level timestamp, supports high-concurrency scenarios
- **Memory Usage**: Extremely low memory footprint, suitable for high-density deployment
- **Concurrency Support**: Fully thread-safe, supports high-concurrency environments

## Important Notes

1. **Clock Synchronization**: Ensure all nodes use NTP service for time synchronization
2. **Node ID Uniqueness**: SystemID must be unique within the same business line
3. **Clock Drift**: System automatically detects and handles clock drift issues
4. **Sequence Overflow**: When sequence number reaches 127, automatically waits for next millisecond

## License

MIT License

## Contributing

Issues and Pull Requests are welcome!

## Changelog

### v1.0.0
- Initial version release
- Support for 64-bit distributed ID generation
- Built-in business classification prefixes
- Base32 encoding support
- Clock drift protection
- Concurrency safety guarantee 