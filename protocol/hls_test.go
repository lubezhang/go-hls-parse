package protocol

import (
	"hls-parse/types"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHlsMasterParse(t *testing.T) {
	assetObj := assert.New(t)

	var strHls = `#EXTM3U

	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000
	1000kbps.m3u8

	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=564000
	/500kbps.m3u8
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=282000
	/path1/250kbps.m3u8
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=2128000
	path2/2000kbps.m3u8`

	hls, err := Parse(&strHls, "https://www.baidu.com")
	assetObj.Nil(err)
	assetObj.Equal(hls.ExtM3u, "#EXTM3U")
	assetObj.Equal(hls.PlayListType, types.PlayListTypeMaster)
}

func TestHlsVodParse(t *testing.T) {
	assetObj := assert.New(t)

	var strHls = `#EXTM3U
    #EXT-X-VERSION:3
    #EXT-X-TARGETDURATION:5
    #EXT-X-PLAYLIST-TYPE:VOD
    #EXT-X-MEDIA-SEQUENCE:0
    #EXTINF:4.128,
    /1000kb/hls/YMgVK9tU.ts
    #EXTINF:3.127,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/3e9Ux5sa.ts
    #EXT-X-KEY:METHOD=AES-128,URI=\"https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/key.key\"
    #EXTINF:3.461,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/ZXyddo0d.ts
    #EXTINF:2.043,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/FsOLD1kG.ts
    #EXTINF:3.127,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/J1Xo6bvk.ts
    #EXTINF:4.253,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/70zdcWHN.ts
    #EXTINF:3.336,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/hZO2SoIF.ts
    #EXTINF:0.917,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/NtSoA2hU.ts
    #EXT-X-KEY:METHOD=AES-128,URI=\"https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/key.key\"
    #EXTINF:3.127,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/E3jKvOa0.ts
    #EXTINF:3.044,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/NeK9QXha.ts
    #EXTINF:3.002,
    https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/q51WSnXk.ts
    #EXT-X-ENDLIST`

	hls, err := Parse(&strHls, "https://www.baidu.com")
	assetObj.Nil(err)
	assetObj.Equal(hls.ExtM3u, "#EXTM3U")
	assetObj.Equal(hls.PlayListType, types.PlayListTypeVod)
}

func TestParsePlayListType(t *testing.T) {
	assetObj := assert.New(t)

	assetObj.Equal(parsePlayListType("#EXT-X-PLAYLIST-TYPE:VOD"), types.PlayListTypeVod)
	assetObj.Equal(parsePlayListType("#EXT-X-PLAYLIST-TYPE:live"), types.PlayListTypeLive)
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
