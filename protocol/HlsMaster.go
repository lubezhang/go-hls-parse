package protocol

import "github.com/lubezhang/hls-parse/types"

type HlsMaster struct {
	HlsBase
	StreamInfs []types.TagStreamInf // 多个视频流
}

type HlsVod struct {
	HlsBase
	ExtInfs  []types.TagExtInf // 视频分片集合
	Extkeys  []types.TagExtKey // 加密密钥集合
	Duration int               // 视频总时长。单位：秒
	Endlist  string            // 文件结束标示
}
