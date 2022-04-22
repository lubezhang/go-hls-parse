package protocol

import (
	"github.com/lubezhang/hls-parse/common"
	"github.com/lubezhang/hls-parse/types"
)

// 获取协议文件类型
func getPlaylistType(arrHls []string) (result types.PlayListType, err error) {
	result = types.PlayListTypeNone

	// 解析协议中的类型配置
	for _, lineStr := range arrHls {
		if common.ExtractTag(lineStr) == types.ProtocolTagPlaylistType {
			result = parsePlayListType(lineStr)
			break
		}
	}

	if result == types.PlayListTypeNone {
		// 如果有结束标签，则认为是VOD格式
		if common.ExtractTag(arrHls[len(arrHls)-1]) == types.ProtocolTagEndlist {
			result = types.PlayListTypeVod
		}
	}

	if result == types.PlayListTypeNone {
		// 如果有视频流数据，则任务是主文件
		for _, lineStr := range arrHls {
			if common.ExtractTag(lineStr) == types.ProtocolTagStreamInf {
				result = types.PlayListTypeMaster
				break
			}
		}
	}

	return
}

// 解析协议文件类型
func parsePlayListType(protocolLine string) (result types.PlayListType) {
	params, err := common.DestructureParams(protocolLine)
	if err == nil {
		switch params.Array[0] {
		case "VOD":
			result = types.PlayListTypeVod
		case "live":
			result = types.PlayListTypeLive
		default:
			result = types.PlayListTypeNone
		}
	}
	return
}

func parseStreamInf(protocolLine string, value string) (result types.TagStreamInf) {
	protocolParams, err := common.DestructureParams(protocolLine)
	param := protocolParams.Map
	streamInf := types.TagStreamInf{}
	if err == nil {
		streamInf.BandWidth, _ = common.StringToInt(param["BANDWIDTH"])
		streamInf.ProgramId, _ = common.StringToInt(param["PROGRAM-ID"])
		streamInf.Codecs = param["CODECS"]
		streamInf.Resolution = param["RESOLUTION"]
		streamInf.Url = value
	}
	result = streamInf
	return
}

func parseExtInf(protocolLine string, value string) (result types.TagExtInf) {
	protocolParams, err := common.DestructureParams(protocolLine)
	param := protocolParams.Array
	extInf := types.TagExtInf{}
	if err == nil {
		extInf.Duration, _ = common.StringToFloat64(param[0])
		extInf.Title = param[1]
		extInf.Url = value
	}
	result = extInf
	return
}

func parseExtKey(protocolLine string) (result types.TagExtKey) {
	protocolParams, err := common.DestructureParams(protocolLine)
	param := protocolParams.Map
	extKey := types.TagExtKey{}
	if err == nil {
		extKey.Method = param["METHOD"]
		extKey.Uri = param["URI"]
		extKey.Iv = param["IV"]
	}
	result = extKey
	return
}
