package snowflakeid

import "time"

// 解构后的ID信息
type SnowflakeID struct {
	SignBit     uint8     `json:"signBit"`     // 固定为0（1位）
	Prefix      uint8     `json:"prefix"`      // 业务线+机器组（5位）
	Version     uint8     `json:"version"`     // 版本号（2位）
	Timestamp   int64     `json:"timestamp"`   // 毫秒时间戳（42位）
	Business    uint8     `json:"business"`    // 子业务线（3位）
	SystemID    uint8     `json:"systemId"`    // 系统节点（4位）
	Sequence    uint8     `json:"sequence"`    // 序列号（7位）
	Base32      string    `json:"base32"`      // Base32编码（13字符）
	CreatedTime time.Time `json:"createdTime"` // 生成时间
}
