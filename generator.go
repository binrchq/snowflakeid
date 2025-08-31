package snowflakeid

import (
	"encoding/base32"
	"errors"
	"log"
	"strconv"
	"sync"
	"time"
)

type Generator struct {
	mu         sync.Mutex
	lastTime   int64
	sequence   uint8
	Prefix     uint8
	BusinessID uint8
	SystemID   uint8
	Logger     *log.Logger
}

func NewGenerator(BusinessID uint8, SystemID uint8) *Generator {
	return &Generator{
		BusinessID: BusinessID,
		SystemID:   SystemID,
	}
}

// 生成ID (返回十进制和Base32)
func (g *Generator) NextID() (int64, string, error) {
	prefix := g.Prefix // 使用Generator中的前缀
	if prefix == 0 {
		prefix = DefaultPrefixBit // 使用默认前缀
	}
	return g.NextIDWithPrefix(prefix)

}
func (g *Generator) NextIDWithPrefix(prefix uint8) (int64, string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().UnixMilli() - Epoch

	// 时钟回拨检查
	if now < g.lastTime {
		return 0, "", errors.New("clock moved backwards")
	}

	// 同一毫秒内递增序列号
	if now == g.lastTime {
		g.sequence = (g.sequence + 1) & MaxSequence
		if g.sequence == 0 {
			// 等待下一毫秒
			for now <= g.lastTime {
				now = time.Now().UnixMilli() - Epoch
			}
		}
	} else {
		g.sequence = 0
	}

	g.logf("now: %d, lastTime: %d, sequence: %d\n", now, g.lastTime, g.sequence)

	g.lastTime = now

	// 使用 uint64 组合 ID，确保每个部分不丢失信息
	id := (0 << SignShift) |
		(int64(prefix) << PrefixShift) |
		(int64(1) << VersionShift) | // 版本号固定为1
		(int64(g.BusinessID) << BusinessShift) |
		(int64(g.SystemID) << SystemShift) |
		(int64(now) << TimestampShift) |
		int64(g.sequence)

	g.logf("prefix: %d, version: %d, timestamp: %d, business: %d, system: %d, sequence: %d\n", prefix, 1, now, g.BusinessID, g.SystemID, g.sequence)
	g.logf("id: %d\n", id)
	//二进制的
	binaryID := strconv.FormatInt(int64(id), 2)
	g.logf("binaryID: %s, len %d\n", binaryID, len(binaryID))

	// Base32编码
	base32ID := ParseID2Base32(int64(id))

	return int64(id), base32ID, nil
}

func ParseID2Base32(id int64) string {
	//第一个bit换到最后

	rotatedID := (id << 1) | (id >> 63)

	// 转换为大端字节序
	buf := [8]byte{
		byte(rotatedID >> 56),
		byte(rotatedID >> 48),
		byte(rotatedID >> 40),
		byte(rotatedID >> 32),
		byte(rotatedID >> 24),
		byte(rotatedID >> 16),
		byte(rotatedID >> 8),
		byte(rotatedID),
	}

	// Base32编码（固定13字符，无填充）
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(buf[:])[:13]
}

func ParseBase322ID(base32ID string) (int64, error) {
	// 解码Base32
	bytes, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(base32ID)
	if err != nil {
		return 0, err
	}

	// 组合字节为uint64（注意大端序）
	var rotated uint64
	for _, b := range bytes {
		rotated = (rotated << 8) | uint64(b)
	}

	// 将最低位还原到最高位
	id := (rotated >> 1) | (rotated << 63)

	// 检查符号位是否为0（确保正数）
	if id < 0 {
		return 0, errors.New("invalid ID: sign bit is 1")
	}

	return int64(id), nil
}

// 解析ID
func ParseID(id int64) SnowflakeID {
	// 确保无符号右移处理时间戳
	timestamp := uint64(id>>TimestampShift) & MaxTimestamp

	return SnowflakeID{
		SignBit:     uint8(id >> SignShift & 0x1), // 明确取最后1位
		Prefix:      uint8(id>>PrefixShift) & MaxPrefix,
		Version:     uint8(id>>VersionShift) & MaxVersion,
		Business:    uint8(id>>BusinessShift) & MaxBusiness,
		SystemID:    uint8(id>>SystemShift) & MaxSystem,
		Timestamp:   int64(timestamp),        // 转换回有符号
		Sequence:    uint8(id & MaxSequence), // 直接取低7位
		CreatedTime: time.UnixMilli(int64(timestamp) + Epoch),
	}
}

// Base32解码
func ParseBase32(base32ID string) (SnowflakeID, error) {
	bytes, err := base32Encoder.DecodeString(base32ID)
	if err != nil {
		return SnowflakeID{}, err
	}

	var id int64
	for _, b := range bytes {
		id = (id << 8) | int64(b)
	}

	sf := ParseID(id)
	sf.Base32 = base32ID
	return sf, nil
}

func (g *Generator) logf(format string, args ...any) {
	if g.Logger != nil {
		g.Logger.Printf(format, args...)
	}
}
