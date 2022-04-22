package types

type PlayListType int
type ProtocolTagType int

const (
	PlayListTypeNone   PlayListType = iota // 未知，未被支持的类型
	PlayListTypeMaster                     // 主文件
	PlayListTypeVod                        // 视频回放
	PlayListTypeLive                       // 直播
)

const (
	ProtocolTagExtm3U       ProtocolTagType = iota // 协议根元素
	ProtocolTagPlaylistType                        // 协议文件类型
	ProtocolTagStreamInf                           // 主文件，包含多个playlist列表，可以选择视频清晰度
	ProtocolTagExtinf                              // 视频流地址
	ProtocolTagEndlist                             // 视频流结束标示
	ProtocolTagKey                                 // 视频流加密密钥
	ProtocolTagValue                               // 协议对应的数据
	ProtocolTagNil                                 // 没有匹配到协议标签
	ProtocolTagExtPrivinf
)
