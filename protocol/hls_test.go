package protocol

import (
	"math"
	"testing"

	"github.com/lubezhang/hls-parse/types"

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

func TestHlsVodParse1(t *testing.T) {
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
	assetObj.Equal(len(hls.Extinf), 11)
}

func TestHlsVodParse2(t *testing.T) {
	assetObj := assert.New(t)

	var strHls = `#EXTM3U
	#EXT-X-VERSION:3
	#EXT-X-MEDIA-SEQUENCE:0
	#EXT-X-ALLOW-CACHE:YES
	#EXT-X-TARGETDURATION:11
	#EXTINF:10.000000,
	#EXT-X-PRIVINF:FILESIZE=871380
	https://ott-prepush-valipl.cp12.wasu.tv/67756D6080932713CFC02204E/03000600006257B6066DD8855FD6DE906E9825-3CD5-4956-9C3C-F9D857D003BF-00001.ts?ccode=0535&duration=2711&expire=18000&psid=e6c0af1b026360d2f1713f690d4759d140346&ups_client_netip=2a9d811c&ups_ts=1649921804&ups_userid=&apscid=&mnid=&rid=2000000061C096AD451A6457EC5CFD86C6BC4B5502000000&operate_type=1&umt=1&type=mp4hdv3&utid=AL3eGmfwC1YCASqdgRwHbxwR&vid=XNTg1NTg0MDYwOA%3D%3D&s=8287885e63b040e18d13&sp=&t=c51976b10323978&cug=2&bc=2&si=5&eo=0&ykfs=871380&vkey=B4aeb9df794df40a439d82820708a3826
	#EXTINF:10.000000,
	#EXT-X-PRIVINF:FILESIZE=993768
	https://ott-prepush-valipl.cp12.wasu.tv/67756D6080932713CFC02204E/03000600006257B6066DD8855FD6DE906E9825-3CD5-4956-9C3C-F9D857D003BF-00002.ts?ccode=0535&duration=2711&expire=18000&psid=e6c0af1b026360d2f1713f690d4759d140346&ups_client_netip=2a9d811c&ups_ts=1649921804&ups_userid=&apscid=&mnid=&rid=2000000061C096AD451A6457EC5CFD86C6BC4B5502000000&operate_type=1&umt=1&type=mp4hdv3&utid=AL3eGmfwC1YCASqdgRwHbxwR&vid=XNTg1NTg0MDYwOA%3D%3D&s=8287885e63b040e18d13&sp=&t=c51976b10323978&cug=2&bc=2&si=5&eo=0&ykfs=871380&vkey=Ba486f758e4decd3d2be1f0e2ff43da5c
	#EXTINF:10.000000,
	#EXT-X-PRIVINF:FILESIZE=1111268
	https://ott-prepush-valipl.cp12.wasu.tv/67756D6080932713CFC02204E/03000600006257B6066DD8855FD6DE906E9825-3CD5-4956-9C3C-F9D857D003BF-00003.ts?ccode=0535&duration=2711&expire=18000&psid=e6c0af1b026360d2f1713f690d4759d140346&ups_client_netip=2a9d811c&ups_ts=1649921804&ups_userid=&apscid=&mnid=&rid=2000000061C096AD451A6457EC5CFD86C6BC4B5502000000&operate_type=1&umt=1&type=mp4hdv3&utid=AL3eGmfwC1YCASqdgRwHbxwR&vid=XNTg1NTg0MDYwOA%3D%3D&s=8287885e63b040e18d13&sp=&t=c51976b10323978&cug=2&bc=2&si=5&eo=0&ykfs=871380&vkey=Bd42b2e6267ae05b55405e6c962d38c77
	#EXT-X-ENDLIST`

	hls, err := Parse(&strHls, "https://www.baidu.com")
	assetObj.Nil(err)
	assetObj.Equal(hls.ExtM3u, "#EXTM3U")
	assetObj.Equal(hls.PlayListType, types.PlayListTypeVod)
	assetObj.Equal(len(hls.Extinf), 3)
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
