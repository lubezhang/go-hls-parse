package protocol

import (
	"hls-parse/common"
	"hls-parse/types"
)

/**
 * Parse the given string as HLS
 * @return HLS
 */
func Parse(strHls *string) (result types.HLS, err error) {
	arrHls, err := common.ProtocolStrToArray(strHls)
	result = types.HLS{
		PlayListType: types.PlayListTypeMaster,
	}
	if err != nil {
		return result, err
	}
	for i := 0; i < len(arrHls); i++ {
		v := arrHls[i]
		switch common.ExtractTag(v) {
		case types.ProtocolTagExtm3U:
			result.ExtM3u = v
		case types.ProtocolTagEndlist:
			result.Endlist = v
		case types.ProtocolTagPlaylistType:
			result.PlayListType = parsePlayListType(v)
		case types.ProtocolTagStreamInf:
			result.ExtXStreamInf = append(result.ExtXStreamInf, parseStreamInf(v, arrHls[i+1]))
			i++
		}
	}
	err = nil
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
			result = types.PlayListTypeMaster
		}
	}
	return
}

func parseStreamInf(protocolLine string, value string) (result types.HlsStreamInf) {
	protocolParams, err := common.DestructureParams(protocolLine)
	param := protocolParams.Map
	if err == nil {
		streamInf := types.HlsStreamInf{}
		streamInf.BandWidth, _ = common.StringToInt(param["BANDWIDTH"])
		streamInf.ProgramId, _ = common.StringToInt(param["PROGRAM-ID"])
		streamInf.Codecs = param["CODECS"]
		streamInf.Resolution = param["RESOLUTION"]
		streamInf.Url = param["URI"]
		result = streamInf
	}
	return
}
