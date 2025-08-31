package snowflakeid

import "encoding/base32"

var (
	// Base32编码器 (无填充、大写)
	base32Encoder = base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding)
)

const (
	// 位分配
	SignBits      = 1  // 符号位固定为0（保留）
	PrefixBits    = 5  // 前缀类型（A-Z,2-6）
	VersionBits   = 2  // 版本号（0-3）
	BusinessBits  = 3  // 业务线（0-7）
	SystemBits    = 4  // 系统标识（0-15）
	TimestampBits = 42 // 时间戳（可支持139年，从Epoch起）
	SequenceBits  = 7  // 序列号（0-127）

	// 偏移量计算（从高位到低位）
	SignShift      = 63                          // 符号位偏移（第63位）
	PrefixShift    = SignShift - PrefixBits      // 前缀偏移（58-62位）
	VersionShift   = PrefixShift - VersionBits   // 版本偏移（56-57位）
	BusinessShift  = VersionShift - BusinessBits // 业务偏移（53-55位）
	SystemShift    = BusinessShift - SystemBits  // 系统偏移（49-52位）
	TimestampShift = SystemShift - TimestampBits // 时间戳偏移（7-48位）
	SequenceShift  = 0                           // 序列号偏移（0-6位）

	// 最大值计算
	MaxPrefix    = 1<<PrefixBits - 1    // 31（0b11111）
	MaxVersion   = 1<<VersionBits - 1   // 3（0b11）
	MaxBusiness  = 1<<BusinessBits - 1  // 7（0b111）
	MaxSystem    = 1<<SystemBits - 1    // 15（0b1111）
	MaxTimestamp = 1<<TimestampBits - 1 // 4398046511103（42位最大值）
	MaxSequence  = 1<<SequenceBits - 1  // 127（0b1111111）

	// // 时间起点（2020-01-01 UTC）
	// Epoch = 1577836800000

	// 时间起点（2025-01-01 UTC）
	Epoch = 1740758400000
)

const (
	// 默认前缀
	DefaultPrefixBit = 0b00000 // A 默认前缀
	DefaultPrefixTen = 0       // A 默认前缀

	// 业务相关
	BusinessBit = 0b00001 // B 业务相关
	BusinessTen = 1       // B 业务相关

	// 客户相关
	CustomerBit = 0b00010 // C 客户相关
	CustomerTen = 2       // C 客户相关

	// 设备相关
	DeviceBit = 0b00011 // D 设备相关
	DeviceTen = 3       // D 设备相关

	// 事件相关
	EventBit = 0b00100 // E 事件相关
	EventTen = 4       // E 事件相关

	// 文件相关
	FileBit = 0b00101 // F 文件相关
	FileTen = 5       // F 文件相关

	// 网关相关
	GatewayBit = 0b00110 // G 网关相关
	GatewayTen = 6       // G 网关相关

	// 主机相关
	HostBit = 0b00111 // H 主机相关
	HostTen = 7       // H 主机相关

	// 实例相关
	InstanceBit = 0b01000 // I 实例相关
	InstanceTen = 8       // I 实例相关

	// 任务相关
	JobBit = 0b01001 // J 任务相关
	JobTen = 9       // J 任务相关

	// Kubernetes相关
	KubernetesBit = 0b01010 // K Kubernetes相关
	KubernetesTen = 10      // K Kubernetes相关

	// 日志相关
	LogBit = 0b01011 // L 日志相关
	LogTen = 11      // L 日志相关

	// 模块相关
	ModuleBit = 0b01100 // M 模块相关
	ModuleTen = 12      // M 模块相关

	// 网络相关
	NetworkBit = 0b01101 // N 网络相关
	NetworkTen = 13      // N 网络相关

	// 组织相关
	OrganizationBit = 0b01110 // O 组织相关
	OrganizationTen = 14      // O 组织相关

	// 项目相关
	ProjectBit = 0b01111 // P 项目相关
	ProjectTen = 15      // P 项目相关

	// 队列相关
	QueueBit = 0b10000 // Q 队列相关
	QueueTen = 16      // Q 队列相关

	// 资源相关
	ResourceBit = 0b10001 // R 资源相关
	ResourceTen = 17      // R 资源相关

	// 服务相关
	ServiceBit = 0b10010 // S 服务相关
	ServiceTen = 18      // S 服务相关

	// 任务相关
	TaskBit = 0b10011 // T 任务相关
	TaskTen = 19      // T 任务相关

	// 用户相关
	UserBit = 0b10100 // U 用户相关
	UserTen = 20      // U 用户相关

	// 版本相关
	VersionBit = 0b10101 // V 版本相关
	VersionTen = 21      // V 版本相关

	// 工作流相关
	WorkflowBit = 0b10110 // W 工作流相关
	WorkflowTen = 22      // W 工作流相关

	// 实验相关
	ExperimentBit = 0b10111 // X 实验相关
	ExperimentTen = 23      // X 实验相关

	// 数据相关
	YieldBit = 0b11000 // Y 数据相关
	YieldTen = 24      // Y 数据相关

	// 区域相关
	ZoneBit = 0b11001 // Z 区域相关
	ZoneTen = 25      // Z 区域相关

	// 双因子认证相关
	TwoFactorBit = 0b11010 // 2 双因子认证相关
	TwoFactorTen = 26      // 2 双因子认证相关

	// 第三方相关
	ThirdPartyBit = 0b11011 // 3 第三方相关
	ThirdPartyTen = 27      // 3 第三方相关

	// 备份相关
	BackupBit = 0b11100 // 4 备份相关
	BackupTen = 28      // 4 备份相关

	// 测试相关
	TestBit = 0b11101 // 5 测试相关
	TestTen = 29      // 5 测试相关

	// 系统相关
	SystemBit = 0b11110 // 6 系统相关
	SystemTen = 30      // 6 系统相关

	// 保留未分配
	ReservedBit = 0b11111 // 7 保留未分配
	ReservedTen = 31      // 7 保留未分配
)
