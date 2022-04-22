package protocol

import (
	"errors"

	"github.com/lubezhang/hls-parse/types"
)

type HlsBase struct {
	ExtM3u       string             // 协议根元素，在协议文件的第一行
	PlayListType types.PlayListType // 协议类型
	orgHls       []string           // 原始协议
	baseUrl      string             // 基础url
}

// 是否为主文件
func (hlsBase *HlsBase) IsMaster() bool {
	return hlsBase.PlayListType == types.PlayListTypeMaster
}

// 视频回放
func (hlsBase *HlsBase) IsVod() bool {
	return hlsBase.PlayListType == types.PlayListTypeVod
}

func (hlsBase *HlsBase) GetMaster() (result HlsMaster, err error) {
	if !hlsBase.IsMaster() {
		return result, errors.New("不是主协议文件")
	}
	result, _ = parseHlsMaster(hlsBase)
	return
}

func (hlsBase *HlsBase) GetVod() (result HlsVod, err error) {
	if !hlsBase.IsVod() {
		return result, errors.New("不是视频回放协议")
	}
	result, _ = parseHlsVod(hlsBase)
	return
}
