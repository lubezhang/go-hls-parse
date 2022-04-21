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

type HLS struct {
	ExtM3u       string         // 协议根元素，在协议文件的第一行
	PlayListType PlayListType   // 协议类型
	ExtStreamInf []HlsStreamInf // 包含多个playlist列表，可以选择视频清晰度
	Extinf       []HlsExtInf    // 视频流地址
	Extkey       []HlsExtKey    // 加密密钥队列
	Endlist      string         // 文件结束标示
}

// 是否为主文件
func (hls *HLS) IsMaster() bool {
	return hls.PlayListType == PlayListTypeMaster
}

// 视频回放
func (hls *HLS) IsVod() bool {
	return hls.PlayListType == PlayListTypeVod
}

type HlsStreamInf struct {
	BandWidth  int    // 带宽。表示对于每个媒体文件所有比特率的上限，单位是 比特/秒
	ProgramId  int    // 唯一地标识一个在 Playlist 文件范围内的特定的描述。一个 Playlist 文件中可能包含多个有相同 ID 的此 tag
	Codecs     string // 编码类型
	Resolution string // 分辨率
	Url        string // 视频流协议文件链接
}

type HlsExtInf struct {
	Index        int
	Duration     float64 // 每个切片的实际时长。单位：秒
	Title        string  // 片的描述
	Url          string  // 每片的链接
	EncryptIndex int     // 当前链接在密钥队列的索引。值为-1 视频没有加密不需要密钥
}

type HlsExtKey struct {
	Index  int
	Method string // 文件加密方式
	Uri    string // 密钥链接
	Key    string // 密钥
	Iv     string
}

type ProtocolParams struct {
	Map   map[string]string // key value形式的参数
	Array []string          // 没有明确的参数key，以数组形式存放的参数
}
