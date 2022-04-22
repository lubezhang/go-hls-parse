package protocol

import (
	"errors"

	"github.com/lubezhang/hls-parse/common"
	"github.com/lubezhang/hls-parse/types"
)

func ParseString(strHls *string, baseUrl string) (result HlsBase, err error) {
	if (strHls == nil) || (len(*strHls) == 0) {
		return result, errors.New("协议为空")
	}
	arrHls, err := common.ProtocolStrToArray(strHls)
	result = HlsBase{
		ExtM3u: arrHls[0],
	}
	if err != nil {
		return result, err
	}
	result.PlayListType, _ = getPlaylistType(arrHls)

	if result.PlayListType == types.PlayListTypeNone {
		return result, errors.New("未知的协议")
	}
	result.orgHls = arrHls

	err = nil
	return
}

// 解析主文件
func parseHlsMaster(hlsBase *HlsBase) (result HlsMaster, err error) {
	result = HlsMaster{}
	result.HlsBase = *hlsBase
	orgHls := hlsBase.orgHls
	for i := 0; i < len(orgHls); i++ {
		line := orgHls[i]
		switch common.ExtractTag(line) {
		case types.ProtocolTagStreamInf:
			si := parseStreamInf(line, orgHls[i+1])
			si.Url, _ = common.JoinUrl(si.Url, hlsBase.baseUrl)
			result.StreamInfs = append(result.StreamInfs, si)
			i++
		}
	}
	return
}

func parseHlsVod(hlsBase *HlsBase) (result HlsVod, err error) {
	result = HlsVod{}
	result.HlsBase = *hlsBase
	orgHls := hlsBase.orgHls
	for i := 0; i < len(orgHls); i++ {
		line := orgHls[i]
		switch common.ExtractTag(line) {
		case types.ProtocolTagKey:
			ek := parseExtKey(line)
			ek.Uri, _ = common.JoinUrl(ek.Uri, hlsBase.baseUrl)
			result.Extkeys = append(result.Extkeys, ek)
		case types.ProtocolTagExtinf:
			// 兼容#EXT-X-PRIVINF标签
			step := 1
			if common.ExtractTag(orgHls[i+1]) == types.ProtocolTagExtPrivinf {
				step = 2
			}

			ei := parseExtInf(line, orgHls[i+step])
			ei.Url, _ = common.JoinUrl(ei.Url, hlsBase.baseUrl)
			ei.EncryptIndex = len(result.Extkeys) - 1
			result.ExtInfs = append(result.ExtInfs, ei)
			i = i + step
		case types.ProtocolTagEndlist:
			result.Endlist = line
		}
	}
	return
}
