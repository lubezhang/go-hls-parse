package protocol

import (
	"io/ioutil"
	"math"
	"testing"

	"github.com/lubezhang/hls-parse/common"
	"github.com/lubezhang/hls-parse/types"

	"github.com/stretchr/testify/assert"
)

func _getMasterString() string {
	content, _ := ioutil.ReadFile("../testData/master.m3u8")
	return string(content)
}

func _getVodString() string {
	content, _ := ioutil.ReadFile("../testData/vod1.m3u8")
	return string(content)
}
func _getVod2String() string {
	content, _ := ioutil.ReadFile("../testData/vod2.m3u8")
	return string(content)
}

func TestHlsBaseParse(t *testing.T) {
	assetObj := assert.New(t)

	var strHls = _getMasterString()

	hlsBase, err := ParseString(&strHls, "https://www.baidu.com")
	assetObj.Nil(err)
	if err == nil {
		assetObj.Equal(hlsBase.ExtM3u, "#EXTM3U")
		assetObj.Equal(hlsBase.PlayListType, types.PlayListTypeMaster)
		assetObj.Equal(hlsBase.IsMaster(), true)
		assetObj.Equal(hlsBase.IsVod(), false)
	}
}

func TestHlsMasterParse(t *testing.T) {
	assetObj := assert.New(t)

	var strHls = _getMasterString()

	hlsBase, err := ParseString(&strHls, "https://www.baidu.com")
	assetObj.Nil(err)
	if err == nil {
		if hlsBase.IsMaster() {
			hlsMaster, _ := hlsBase.GetMaster()
			assetObj.Equal(len(hlsMaster.StreamInfs), 4)
		}
	}

	var strHls1 = _getVodString()
	hlsBase1, err1 := ParseString(&strHls1, "https://www.baidu.com")
	assetObj.Nil(err1)
	if err1 == nil {
		_, err2 := hlsBase1.GetMaster()
		assetObj.NotNil(err2)
		assetObj.EqualError(err2, "不是主协议文件")
	}
}

func TestHlsVodParse(t *testing.T) {
	assetObj := assert.New(t)
	var strHls = _getVodString()
	hlsBase, err := ParseString(&strHls, "https://www.baidu.com")
	assetObj.Nil(err)

	if err == nil {
		_, err2 := hlsBase.GetMaster()
		assetObj.NotNil(err2)
		assetObj.EqualError(err2, "不是主协议文件")
		if hlsBase.IsVod() {
			hlsVod, _ := hlsBase.GetVod()
			assetObj.Equal(len(hlsVod.ExtInfs), 11)
			assetObj.Equal(len(hlsVod.Extkeys), 2)
			assetObj.Equal(hlsVod.Endlist, "#EXT-X-ENDLIST")
		}
	}
}

func TestParsePlayListType(t *testing.T) {
	assetObj := assert.New(t)

	assetObj.Equal(parsePlayListType("#EXT-X-PLAYLIST-TYPE:VOD"), types.PlayListTypeVod)
	assetObj.Equal(parsePlayListType("#EXT-X-PLAYLIST-TYPE:live"), types.PlayListTypeLive)
	assetObj.Equal(parsePlayListType("#EXT-X-PLAYLIST-T2YPE:live"), types.PlayListTypeNone)
}

func TestGetPlaylistType(t *testing.T) {
	assetObj := assert.New(t)

	// 主文件
	masterContent := _getMasterString()
	arrHls, _ := common.ProtocolStrToArray(&masterContent)
	playListType, _ := getPlaylistType(arrHls)
	assetObj.Equal(playListType, types.PlayListTypeMaster)

	// 协议中有EXT-X-PLAYLIST-TYPE标签
	vodContent := _getVodString()
	arrHls1, _ := common.ProtocolStrToArray(&vodContent)
	playListType1, _ := getPlaylistType(arrHls1)
	assetObj.Equal(playListType1, types.PlayListTypeVod)

	// 协议中没有EXT-X-PLAYLIST-TYPE标签，但是存在结束标签#EXT-X-ENDLIST
	vodContent2 := _getVod2String()
	arrHls2, _ := common.ProtocolStrToArray(&vodContent2)
	playListType2, _ := getPlaylistType(arrHls2)
	assetObj.Equal(playListType2, types.PlayListTypeVod)
}

func TestParseStreamInf(t *testing.T) {
	assetObj := assert.New(t)

	si := parseStreamInf("#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000", "1000kbps.m3u8")
	assetObj.Equal(si.BandWidth, 1064000)
	assetObj.Equal(si.ProgramId, 1)
	assetObj.Equal(si.Url, "1000kbps.m3u8")
}

func TestParseExtInf(t *testing.T) {
	const MIN = 0.000001
	assetObj := assert.New(t)

	ei := parseExtInf("#EXTINF:3.127, title", "https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/70zdcWHN.ts")
	assetObj.Equal(math.Dim(ei.Duration, 3.127) < MIN, true)
	assetObj.Equal(ei.Title, "title")
	assetObj.Equal(ei.Url, "https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/70zdcWHN.ts")
}

func TestParseExtKey(t *testing.T) {
	assetObj := assert.New(t)

	tag1 := "#EXT-X-KEY:METHOD=AES-128,URI=\"https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/key.key\""
	ek := parseExtKey(tag1)

	assetObj.Equal(ek.Method, "AES-128")
	assetObj.Equal(ek.Uri, "https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/key.key")
}
