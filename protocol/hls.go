package protocol

import (
	"github.com/lubezhang/hls-parse/common"
	"github.com/lubezhang/hls-parse/types"
)

/**
 * Parse the given string as HLS
 * @return HLS
 */
func Parse(strHls *string, baseUrl string) (result types.HLS, err error) {
	arrHls, err := common.ProtocolStrToArray(strHls)
	result = types.HLS{
		PlayListType: types.PlayListTypeNone,
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
			si := parseStreamInf(v, arrHls[i+1])
			si.Url, _ = common.JoinUrl(si.Url, baseUrl)
			result.ExtStreamInf = append(result.ExtStreamInf, si)
			i++
		case types.ProtocolTagExtinf:
			// 兼容#EXT-X-PRIVINF标签
			step := 1
			if common.ExtractTag(arrHls[i+1]) == types.ProtocolTagExtPrivinf {
				step = 2
			}

			ei := parseExtInf(v, arrHls[i+step])
			ei.Url, _ = common.JoinUrl(ei.Url, baseUrl)
			ei.EncryptIndex = len(result.Extkey) - 1
			result.Extinf = append(result.Extinf, ei)
			i = i + step
		case types.ProtocolTagKey:
			ek := parseExtKey(v)
			ek.Uri, _ = common.JoinUrl(ek.Uri, baseUrl)
			result.Extkey = append(result.Extkey, ek)
		}
	}

	// 如果文件有结束标签，则认为是VOD格式
	if common.ExtractTag(arrHls[len(arrHls)-1]) == types.ProtocolTagEndlist {
		result.PlayListType = types.PlayListTypeVod
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
	streamInf := types.HlsStreamInf{}
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

func parseExtInf(protocolLine string, value string) (result types.HlsExtInf) {
	protocolParams, err := common.DestructureParams(protocolLine)
	param := protocolParams.Array
	extInf := types.HlsExtInf{}
	if err == nil {
		extInf.Duration, _ = common.StringToFloat64(param[0])
		extInf.Title = param[1]
		extInf.Url = value
	}
	result = extInf
	return
}

func parseExtKey(protocolLine string) (result types.HlsExtKey) {
	protocolParams, err := common.DestructureParams(protocolLine)
	param := protocolParams.Map
	extKey := types.HlsExtKey{}
	if err == nil {
		extKey.Method = param["METHOD"]
		extKey.Uri = param["URI"]
		extKey.Iv = param["IV"]
	}
	result = extKey
	return
}
