package types

// 视频流信息
type TagStreamInf struct {
	BandWidth  int    // 带宽。表示对于每个媒体文件所有比特率的上限，单位是 比特/秒
	ProgramId  int    // 唯一地标识一个在 Playlist 文件范围内的特定的描述。一个 Playlist 文件中可能包含多个有相同 ID 的此 tag
	Codecs     string // 编码类型
	Resolution string // 分辨率
	Url        string // 视频流协议文件链接
}

// 视频分片信息
type TagExtInf struct {
	Index        int     //
	Duration     float64 // 每个切片的实际时长。单位：秒
	Title        string  // 片的描述
	Url          string  // 每片的链接
	EncryptIndex int     // 当前链接在密钥队列的索引。值为-1 视频没有加密不需要密钥
}

// 视频分片加密密钥
type TagExtKey struct {
	Index  int    //
	Method string // 文件加密方式
	Uri    string // 密钥链接
	Key    string // 密钥
	Iv     string
}

// 协议标签携带的参数集合
type ProtocolParams struct {
	Map   map[string]string // key / value形式的参数
	Array []string          // 没有明确的参数key，以数组形式存放的参数，根据参数所在的索引位置确定含义
}
